package v1

import (
	"context"
	"log/slog"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/common/log"
	"github.com/bytebase/bytebase/backend/component/dbfactory"
	"github.com/bytebase/bytebase/backend/component/iam"
	"github.com/bytebase/bytebase/backend/component/secret"
	"github.com/bytebase/bytebase/backend/component/state"
	enterprise "github.com/bytebase/bytebase/backend/enterprise/api"
	api "github.com/bytebase/bytebase/backend/legacyapi"
	metricapi "github.com/bytebase/bytebase/backend/metric"
	"github.com/bytebase/bytebase/backend/plugin/db"
	"github.com/bytebase/bytebase/backend/plugin/metric"
	pgparser "github.com/bytebase/bytebase/backend/plugin/parser/pg"
	"github.com/bytebase/bytebase/backend/runner/metricreport"
	"github.com/bytebase/bytebase/backend/runner/schemasync"
	"github.com/bytebase/bytebase/backend/store"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
	v1pb "github.com/bytebase/bytebase/proto/generated-go/v1"
)

// InstanceService implements the instance service.
type InstanceService struct {
	v1pb.UnimplementedInstanceServiceServer
	store          *store.Store
	licenseService enterprise.LicenseService
	metricReporter *metricreport.Reporter
	secret         string
	stateCfg       *state.State
	dbFactory      *dbfactory.DBFactory
	schemaSyncer   *schemasync.Syncer
	iamManager     *iam.Manager
}

// NewInstanceService creates a new InstanceService.
func NewInstanceService(store *store.Store, licenseService enterprise.LicenseService, metricReporter *metricreport.Reporter, secret string, stateCfg *state.State, dbFactory *dbfactory.DBFactory, schemaSyncer *schemasync.Syncer, iamManager *iam.Manager) *InstanceService {
	return &InstanceService{
		store:          store,
		licenseService: licenseService,
		metricReporter: metricReporter,
		secret:         secret,
		stateCfg:       stateCfg,
		dbFactory:      dbFactory,
		schemaSyncer:   schemaSyncer,
		iamManager:     iamManager,
	}
}

// GetInstance gets an instance.
func (s *InstanceService) GetInstance(ctx context.Context, request *v1pb.GetInstanceRequest) (*v1pb.Instance, error) {
	instance, err := getInstanceMessage(ctx, s.store, request.Name)
	if err != nil {
		return nil, err
	}
	return convertInstanceMessage(instance)
}

// ListInstances lists all instances.
func (s *InstanceService) ListInstances(ctx context.Context, request *v1pb.ListInstancesRequest) (*v1pb.ListInstancesResponse, error) {
	find := &store.FindInstanceMessage{
		ShowDeleted: request.ShowDeleted,
	}
	instances, err := s.store.ListInstancesV2(ctx, find)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &v1pb.ListInstancesResponse{}
	for _, instance := range instances {
		ins, err := convertInstanceMessage(instance)
		if err != nil {
			return nil, err
		}
		response.Instances = append(response.Instances, ins)
	}
	return response, nil
}

// ListInstanceDatabase list all databases in the instance.
func (s *InstanceService) ListInstanceDatabase(ctx context.Context, request *v1pb.ListInstanceDatabaseRequest) (*v1pb.ListInstanceDatabaseResponse, error) {
	var instanceMessage *store.InstanceMessage

	if request.Instance != nil {
		instanceID, err := common.GetInstanceID(request.Name)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if instanceMessage, err = s.convertInstanceToInstanceMessage(instanceID, request.Instance); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	} else {
		instance, err := getInstanceMessage(ctx, s.store, request.Name)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		instanceMessage = instance
	}

	instanceMeta, err := s.schemaSyncer.GetInstanceMeta(ctx, instanceMessage)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &v1pb.ListInstanceDatabaseResponse{}
	for _, database := range instanceMeta.Databases {
		response.Databases = append(response.Databases, database.Name)
	}
	return response, nil
}

// CreateInstance creates an instance.
func (s *InstanceService) CreateInstance(ctx context.Context, request *v1pb.CreateInstanceRequest) (*v1pb.Instance, error) {
	if request.Instance == nil {
		return nil, status.Errorf(codes.InvalidArgument, "instance must be set")
	}
	if !isValidResourceID(request.InstanceId) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid instance ID %v", request.InstanceId)
	}

	if err := s.instanceCountGuard(ctx); err != nil {
		return nil, err
	}

	instanceMessage, err := s.convertInstanceToInstanceMessage(request.InstanceId, request.Instance)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Test connection.
	if request.ValidateOnly {
		for _, ds := range instanceMessage.Metadata.GetDataSources() {
			err := func() error {
				driver, err := s.dbFactory.GetDataSourceDriver(ctx, instanceMessage, ds, "", false /* datashare */, ds.GetType() == storepb.DataSourceType_READ_ONLY, db.ConnectionContext{})
				if err != nil {
					return status.Errorf(codes.Internal, "failed to get database driver with error: %v", err.Error())
				}
				defer driver.Close(ctx)
				if err := driver.Ping(ctx); err != nil {
					return status.Errorf(codes.InvalidArgument, "invalid datasource %s, error %s", ds.GetType(), err)
				}
				return nil
			}()
			if err != nil {
				return nil, err
			}
		}

		return convertInstanceMessage(instanceMessage)
	}

	instanceCountLimit := s.licenseService.GetInstanceLicenseCount(ctx)
	if instanceMessage.Metadata.GetActivation() {
		count, err := s.store.GetActivatedInstanceCount(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if count >= instanceCountLimit {
			return nil, status.Errorf(codes.ResourceExhausted, instanceExceededError, instanceCountLimit)
		}
	}

	if err := s.checkInstanceDataSources(instanceMessage, instanceMessage.Metadata.GetDataSources()); err != nil {
		return nil, err
	}

	instance, err := s.store.CreateInstanceV2(ctx, instanceMessage)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	driver, err := s.dbFactory.GetAdminDatabaseDriver(ctx, instance, nil /* database */, db.ConnectionContext{})
	if err == nil {
		defer driver.Close(ctx)
		updatedInstance, _, _, err := s.schemaSyncer.SyncInstance(ctx, instance)
		if err != nil {
			slog.Warn("Failed to sync instance",
				slog.String("instance", instance.ResourceID),
				log.BBError(err))
		} else {
			instance = updatedInstance
		}
		// Sync all databases in the instance asynchronously.
		s.schemaSyncer.SyncAllDatabases(ctx, instance)
	}

	s.metricReporter.Report(ctx, &metric.Metric{
		Name:  metricapi.InstanceCreateMetricName,
		Value: 1,
		Labels: map[string]any{
			"engine": instance.Metadata.GetEngine(),
		},
	})

	return convertInstanceMessage(instance)
}

func (s *InstanceService) checkInstanceDataSources(instance *store.InstanceMessage, dataSources []*storepb.DataSource) error {
	dsIDMap := map[string]bool{}
	for _, ds := range dataSources {
		if err := s.checkDataSource(instance, ds); err != nil {
			return err
		}
		if dsIDMap[ds.GetId()] {
			return status.Errorf(codes.InvalidArgument, `duplicate data source id "%s"`, ds.GetId())
		}
		dsIDMap[ds.GetId()] = true
	}

	return nil
}

var instanceExceededError = "activation instance count has reached the limit (%v)"

func (s *InstanceService) checkDataSource(instance *store.InstanceMessage, dataSource *storepb.DataSource) error {
	if dataSource.GetId() == "" {
		return status.Errorf(codes.InvalidArgument, "data source id is required")
	}
	password, err := common.Unobfuscate(dataSource.GetObfuscatedPassword(), s.secret)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := s.licenseService.IsFeatureEnabledForInstance(api.FeatureExternalSecretManager, instance); err != nil {
		missingFeatureError := status.Error(codes.PermissionDenied, err.Error())
		if dataSource.GetExternalSecret() != nil {
			return missingFeatureError
		}
		if ok, _ := secret.GetExternalSecretURL(password); !ok {
			return nil
		}
		return missingFeatureError
	}

	return nil
}

// UpdateInstance updates an instance.
func (s *InstanceService) UpdateInstance(ctx context.Context, request *v1pb.UpdateInstanceRequest) (*v1pb.Instance, error) {
	if request.Instance == nil {
		return nil, status.Errorf(codes.InvalidArgument, "instance must be set")
	}
	if request.UpdateMask == nil {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be set")
	}

	instance, err := getInstanceMessage(ctx, s.store, request.Instance.Name)
	if err != nil {
		return nil, err
	}
	if instance.Deleted {
		return nil, status.Errorf(codes.NotFound, "instance %q has been deleted", request.Instance.Name)
	}

	metadata, ok := proto.Clone(instance.Metadata).(*storepb.Instance)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to convert instance metadata type")
	}
	patch := &store.UpdateInstanceMessage{
		ResourceID: instance.ResourceID,
		Metadata:   metadata,
	}
	updateActivation := false
	for _, path := range request.UpdateMask.Paths {
		switch path {
		case "title":
			patch.Metadata.Title = request.Instance.Title
		case "environment":
			environmentID, err := common.GetEnvironmentID(request.Instance.Environment)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			environment, err := s.store.GetEnvironmentV2(ctx, &store.FindEnvironmentMessage{
				ResourceID:  &environmentID,
				ShowDeleted: true,
			})
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			if environment == nil {
				return nil, status.Errorf(codes.NotFound, "environment %q not found", environmentID)
			}
			if environment.Deleted {
				return nil, status.Errorf(codes.FailedPrecondition, "environment %q is deleted", environmentID)
			}
			patch.EnvironmentID = &environment.ResourceID
		case "external_link":
			patch.Metadata.ExternalLink = request.Instance.ExternalLink
		case "data_sources":
			dataSources, err := s.convertV1DataSources(request.Instance.DataSources)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			if err := s.checkInstanceDataSources(instance, dataSources); err != nil {
				return nil, err
			}
			patch.Metadata.DataSources = dataSources
		case "activation":
			if !instance.Metadata.GetActivation() && request.Instance.Activation {
				updateActivation = true
			}
			patch.Metadata.Activation = request.Instance.Activation
		case "sync_interval":
			if err := s.licenseService.IsFeatureEnabledForInstance(api.FeatureCustomInstanceSynchronization, instance); err != nil {
				return nil, status.Error(codes.PermissionDenied, err.Error())
			}
			patch.Metadata.SyncInterval = request.Instance.SyncInterval
		case "maximum_connections":
			if err := s.licenseService.IsFeatureEnabledForInstance(api.FeatureCustomInstanceSynchronization, instance); err != nil {
				return nil, status.Error(codes.PermissionDenied, err.Error())
			}
			patch.Metadata.MaximumConnections = request.Instance.MaximumConnections
		case "sync_databases":
			if err := s.licenseService.IsFeatureEnabledForInstance(api.FeatureCustomInstanceSynchronization, instance); err != nil {
				return nil, status.Error(codes.PermissionDenied, err.Error())
			}
			patch.Metadata.SyncDatabases = request.Instance.SyncDatabases
		default:
			return nil, status.Errorf(codes.InvalidArgument, `unsupported update_mask "%s"`, path)
		}
	}

	instanceCountLimit := s.licenseService.GetInstanceLicenseCount(ctx)
	if updateActivation {
		count, err := s.store.GetActivatedInstanceCount(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if count >= instanceCountLimit {
			return nil, status.Errorf(codes.ResourceExhausted, instanceExceededError, instanceCountLimit)
		}
	}

	ins, err := s.store.UpdateInstanceV2(ctx, patch)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return convertInstanceMessage(ins)
}

func (s *InstanceService) syncSlowQueriesForInstance(ctx context.Context, instanceName string) (*emptypb.Empty, error) {
	instance, err := getInstanceMessage(ctx, s.store, instanceName)
	if err != nil {
		return nil, err
	}
	if instance.Deleted {
		return nil, status.Errorf(codes.NotFound, "instance %q has been deleted", instanceName)
	}

	slowQueryPolicy, err := s.store.GetSlowQueryPolicy(ctx, instance.ResourceID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if slowQueryPolicy == nil || !slowQueryPolicy.Active {
		return nil, status.Errorf(codes.FailedPrecondition, "slow query policy is not active for instance %q", instanceName)
	}

	if err := s.syncSlowQueriesImpl(ctx, (*store.ProjectMessage)(nil), instance); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *InstanceService) syncSlowQueriesImpl(ctx context.Context, project *store.ProjectMessage, instance *store.InstanceMessage) error {
	switch instance.Metadata.GetEngine() {
	case storepb.Engine_MYSQL:
		driver, err := s.dbFactory.GetAdminDatabaseDriver(ctx, instance, nil /* database */, db.ConnectionContext{})
		if err != nil {
			return err
		}
		defer driver.Close(ctx)
		if err := driver.CheckSlowQueryLogEnabled(ctx); err != nil {
			slog.Warn("slow query log is not enabled", slog.String("instance", instance.ResourceID), log.BBError(err))
			return nil
		}

		// Sync slow queries for instance.
		message := &state.InstanceSlowQuerySyncMessage{
			InstanceID: instance.ResourceID,
		}
		if project != nil {
			message.ProjectID = project.ResourceID
		}
		s.stateCfg.InstanceSlowQuerySyncChan <- message
	case storepb.Engine_POSTGRES:
		findDatabase := &store.FindDatabaseMessage{
			InstanceID: &instance.ResourceID,
		}
		databases, err := s.store.ListDatabases(ctx, findDatabase)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to list databases: %s", err.Error())
		}

		enabled := false
		for _, database := range databases {
			if database.Deleted {
				continue
			}
			if pgparser.IsSystemDatabase(database.DatabaseName) {
				continue
			}
			if err := func() error {
				driver, err := s.dbFactory.GetAdminDatabaseDriver(ctx, instance, database, db.ConnectionContext{})
				if err != nil {
					return err
				}
				defer driver.Close(ctx)
				return driver.CheckSlowQueryLogEnabled(ctx)
			}(); err != nil {
				slog.Warn("slow query log is not enabled", slog.String("database", database.DatabaseName), log.BBError(err))
				continue
			}

			enabled = true
			break
		}

		if !enabled {
			return nil
		}

		// Sync slow queries for instance.
		message := &state.InstanceSlowQuerySyncMessage{
			InstanceID: instance.ResourceID,
		}
		if project != nil {
			message.ProjectID = project.ResourceID
		}
		s.stateCfg.InstanceSlowQuerySyncChan <- message
	default:
		return status.Errorf(codes.InvalidArgument, "unsupported engine %q", instance.Metadata.GetEngine())
	}
	return nil
}

func (s *InstanceService) syncSlowQueriesForProject(ctx context.Context, projectName string) (*emptypb.Empty, error) {
	project, err := s.getProjectMessage(ctx, projectName)
	if err != nil {
		return nil, err
	}
	if project.Deleted {
		return nil, status.Errorf(codes.NotFound, "project %q has been deleted", projectName)
	}
	databases, err := s.store.ListDatabases(ctx, &store.FindDatabaseMessage{ProjectID: &project.ResourceID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list databases: %s", err.Error())
	}

	instanceMap := make(map[string]bool)
	var errs error
	for _, database := range databases {
		instance, err := s.store.GetInstanceV2(ctx, &store.FindInstanceMessage{ResourceID: &database.InstanceID})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get instance %q: %s", database.InstanceID, err.Error())
		}

		switch instance.Metadata.GetEngine() {
		case storepb.Engine_MYSQL, storepb.Engine_POSTGRES:
			if instance.Deleted {
				continue
			}

			slowQueryPolicy, err := s.store.GetSlowQueryPolicy(ctx, instance.ResourceID)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			if slowQueryPolicy == nil || !slowQueryPolicy.Active {
				continue
			}

			if _, ok := instanceMap[instance.ResourceID]; ok {
				continue
			}

			if err := s.syncSlowQueriesImpl(ctx, project, instance); err != nil {
				errs = multierr.Append(errs, errors.Wrapf(err, "failed to sync slow queries for instance %q", instance.ResourceID))
			}
		default:
			continue
		}
	}

	if errs != nil {
		return nil, status.Errorf(codes.Internal, "failed to sync slow queries for following instances: %s", errs.Error())
	}

	return &emptypb.Empty{}, nil
}

// SyncSlowQueries syncs slow queries for an instance.
func (s *InstanceService) SyncSlowQueries(ctx context.Context, request *v1pb.SyncSlowQueriesRequest) (*emptypb.Empty, error) {
	switch {
	case strings.HasPrefix(request.Parent, common.InstanceNamePrefix):
		return s.syncSlowQueriesForInstance(ctx, request.Parent)
	case strings.HasPrefix(request.Parent, common.ProjectNamePrefix):
		return s.syncSlowQueriesForProject(ctx, request.Parent)
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q", request.Parent)
	}
}

// DeleteInstance deletes an instance.
func (s *InstanceService) DeleteInstance(ctx context.Context, request *v1pb.DeleteInstanceRequest) (*emptypb.Empty, error) {
	instance, err := getInstanceMessage(ctx, s.store, request.Name)
	if err != nil {
		return nil, err
	}
	if instance.Deleted {
		return nil, status.Errorf(codes.NotFound, "instance %q has been deleted", request.Name)
	}

	databases, err := s.store.ListDatabases(ctx, &store.FindDatabaseMessage{InstanceID: &instance.ResourceID})
	if err != nil {
		return nil, err
	}
	if request.Force {
		if len(databases) > 0 {
			defaultProjectID := api.DefaultProjectID
			if _, err := s.store.BatchUpdateDatabases(ctx, databases, &store.BatchUpdateDatabases{ProjectID: &defaultProjectID}); err != nil {
				return nil, err
			}
		}
	} else {
		var databaseNames []string
		for _, database := range databases {
			if database.ProjectID != api.DefaultProjectID {
				databaseNames = append(databaseNames, database.DatabaseName)
			}
		}
		if len(databaseNames) > 0 {
			return nil, status.Errorf(codes.FailedPrecondition, "all databases should be transferred to the unassigned project before deleting the instance")
		}
	}

	metadata, ok := proto.Clone(instance.Metadata).(*storepb.Instance)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to convert instance metadata type")
	}
	metadata.Activation = false
	if _, err := s.store.UpdateInstanceV2(ctx, &store.UpdateInstanceMessage{
		ResourceID: instance.ResourceID,
		Deleted:    &deletePatch,
		Metadata:   metadata,
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// UndeleteInstance undeletes an instance.
func (s *InstanceService) UndeleteInstance(ctx context.Context, request *v1pb.UndeleteInstanceRequest) (*v1pb.Instance, error) {
	instance, err := getInstanceMessage(ctx, s.store, request.Name)
	if err != nil {
		return nil, err
	}
	if !instance.Deleted {
		return nil, status.Errorf(codes.InvalidArgument, "instance %q is active", request.Name)
	}

	ins, err := s.store.UpdateInstanceV2(ctx, &store.UpdateInstanceMessage{
		ResourceID: instance.ResourceID,
		Deleted:    &undeletePatch,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return convertInstanceMessage(ins)
}

// SyncInstance syncs the instance.
func (s *InstanceService) SyncInstance(ctx context.Context, request *v1pb.SyncInstanceRequest) (*v1pb.SyncInstanceResponse, error) {
	instance, err := getInstanceMessage(ctx, s.store, request.Name)
	if err != nil {
		return nil, err
	}
	if instance.Deleted {
		return nil, status.Errorf(codes.NotFound, "instance %q has been deleted", request.Name)
	}

	updatedInstance, allDatabases, newDatabases, err := s.schemaSyncer.SyncInstance(ctx, instance)
	if err != nil {
		return nil, err
	}
	if request.EnableFullSync {
		// Sync all databases in the instance asynchronously.
		s.schemaSyncer.SyncAllDatabases(ctx, updatedInstance)
	} else {
		s.schemaSyncer.SyncDatabasesAsync(newDatabases)
	}

	response := &v1pb.SyncInstanceResponse{}
	for _, database := range allDatabases {
		response.Databases = append(response.Databases, database.Name)
	}
	return response, nil
}

// SyncInstance syncs the instance.
func (s *InstanceService) BatchSyncInstances(ctx context.Context, request *v1pb.BatchSyncInstancesRequest) (*v1pb.BatchSyncInstancesResponse, error) {
	for _, r := range request.Requests {
		instance, err := getInstanceMessage(ctx, s.store, r.Name)
		if err != nil {
			return nil, err
		}
		if instance.Deleted {
			return nil, status.Errorf(codes.NotFound, "instance %q has been deleted", r.Name)
		}

		updatedInstance, _, newDatabases, err := s.schemaSyncer.SyncInstance(ctx, instance)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to sync instance, %v", err)
		}
		if r.EnableFullSync {
			// Sync all databases in the instance asynchronously.
			s.schemaSyncer.SyncAllDatabases(ctx, updatedInstance)
		} else {
			s.schemaSyncer.SyncDatabasesAsync(newDatabases)
		}
	}

	return &v1pb.BatchSyncInstancesResponse{}, nil
}

// AddDataSource adds a data source to an instance.
func (s *InstanceService) AddDataSource(ctx context.Context, request *v1pb.AddDataSourceRequest) (*v1pb.Instance, error) {
	if request.DataSource == nil {
		return nil, status.Errorf(codes.InvalidArgument, "data sources is required")
	}
	// We only support add RO type datasouce to instance now, see more details in instance_service.proto.
	if request.DataSource.Type != v1pb.DataSourceType_READ_ONLY {
		return nil, status.Errorf(codes.InvalidArgument, "only support adding read-only data source")
	}

	dataSource, err := s.convertV1DataSource(request.DataSource)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert data source")
	}

	instance, err := getInstanceMessage(ctx, s.store, request.Name)
	if err != nil {
		return nil, err
	}
	if instance.Deleted {
		return nil, status.Errorf(codes.NotFound, "instance %q has been deleted", request.Name)
	}
	for _, ds := range instance.Metadata.GetDataSources() {
		if ds.GetId() == request.DataSource.Id {
			return nil, status.Errorf(codes.NotFound, "data source already exists with the same name")
		}
	}
	if err := s.checkDataSource(instance, dataSource); err != nil {
		return nil, err
	}

	// Test connection.
	if request.ValidateOnly {
		err := func() error {
			driver, err := s.dbFactory.GetDataSourceDriver(ctx, instance, dataSource, "", false /* datashare */, dataSource.GetType() == storepb.DataSourceType_READ_ONLY, db.ConnectionContext{})
			if err != nil {
				return status.Errorf(codes.Internal, "failed to get database driver with error: %v", err.Error())
			}
			defer driver.Close(ctx)
			if err := driver.Ping(ctx); err != nil {
				return status.Errorf(codes.InvalidArgument, "invalid datasource %s, error %s", dataSource.GetType(), err)
			}
			return nil
		}()
		if err != nil {
			return nil, err
		}
		return convertInstanceMessage(instance)
	}

	if dataSource.GetType() != storepb.DataSourceType_READ_ONLY {
		return nil, status.Error(codes.InvalidArgument, "only read-only data source can be added.")
	}
	if err := s.licenseService.IsFeatureEnabledForInstance(api.FeatureReadReplicaConnection, instance); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	metadata, ok := proto.Clone(instance.Metadata).(*storepb.Instance)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to convert instance metadata type")
	}
	metadata.DataSources = append(metadata.DataSources, dataSource)
	instance, err = s.store.UpdateInstanceV2(ctx, &store.UpdateInstanceMessage{ResourceID: instance.ResourceID, Metadata: metadata})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return convertInstanceMessage(instance)
}

// UpdateDataSource updates a data source of an instance.
func (s *InstanceService) UpdateDataSource(ctx context.Context, request *v1pb.UpdateDataSourceRequest) (*v1pb.Instance, error) {
	if request.DataSource == nil {
		return nil, status.Errorf(codes.InvalidArgument, "datasource is required")
	}
	if request.UpdateMask == nil {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be set")
	}

	instance, err := getInstanceMessage(ctx, s.store, request.Name)
	if err != nil {
		return nil, err
	}
	if instance.Deleted {
		return nil, status.Errorf(codes.NotFound, "instance %q has been deleted", request.Name)
	}
	metadata, ok := proto.Clone(instance.Metadata).(*storepb.Instance)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to convert instance metadata type")
	}
	var dataSource *storepb.DataSource
	for _, ds := range metadata.GetDataSources() {
		if ds.GetId() == request.DataSource.Id {
			dataSource = ds
			break
		}
	}
	if dataSource == nil {
		return nil, status.Errorf(codes.NotFound, `cannot found data source "%s"`, request.DataSource.Id)
	}

	if dataSource.GetType() == storepb.DataSourceType_READ_ONLY {
		if err := s.licenseService.IsFeatureEnabledForInstance(api.FeatureReadReplicaConnection, instance); err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
	}

	hasSSH := false
	for _, path := range request.UpdateMask.Paths {
		switch path {
		case "username":
			dataSource.Username = request.DataSource.Username
		case "password":
			dataSource.ObfuscatedPassword = common.Obfuscate(request.DataSource.Password, s.secret)
		case "ssl_ca":
			dataSource.ObfuscatedSslCa = common.Obfuscate(request.DataSource.SslCa, s.secret)
		case "ssl_cert":
			dataSource.ObfuscatedSslCert = common.Obfuscate(request.DataSource.SslCert, s.secret)
		case "ssl_key":
			dataSource.ObfuscatedSslKey = common.Obfuscate(request.DataSource.SslKey, s.secret)
		case "host":
			dataSource.Host = request.DataSource.Host
		case "port":
			dataSource.Port = request.DataSource.Port
		case "database":
			dataSource.Database = request.DataSource.Database
		case "srv":
			dataSource.Srv = request.DataSource.Srv
		case "authentication_database":
			dataSource.AuthenticationDatabase = request.DataSource.AuthenticationDatabase
		case "sid":
			dataSource.Sid = request.DataSource.Sid
		case "service_name":
			dataSource.ServiceName = request.DataSource.ServiceName
		case "ssh_host":
			dataSource.SshHost = request.DataSource.SshHost
			hasSSH = true
		case "ssh_port":
			dataSource.SshPort = request.DataSource.SshPort
			hasSSH = true
		case "ssh_user":
			dataSource.SshUser = request.DataSource.SshUser
			hasSSH = true
		case "ssh_password":
			dataSource.ObfuscatedSshPassword = common.Obfuscate(request.DataSource.SshPassword, s.secret)
			hasSSH = true
		case "ssh_private_key":
			dataSource.ObfuscatedSshPrivateKey = common.Obfuscate(request.DataSource.SshPrivateKey, s.secret)
			hasSSH = true
		case "authentication_private_key":
			dataSource.ObfuscatedAuthenticationPrivateKey = common.Obfuscate(request.DataSource.AuthenticationPrivateKey, s.secret)
		case "external_secret":
			externalSecret, err := convertV1DataSourceExternalSecret(request.DataSource.ExternalSecret)
			if err != nil {
				return nil, err
			}
			dataSource.ExternalSecret = externalSecret
		case "sasl_config":
			dataSource.SaslConfig = convertV1DataSourceSaslConfig(request.DataSource.SaslConfig)
		case "authentication_type":
			dataSource.AuthenticationType = convertV1AuthenticationType(request.DataSource.AuthenticationType)
		case "additional_addresses":
			dataSource.AdditionalAddresses = convertAdditionalAddresses(request.DataSource.AdditionalAddresses)
		case "replica_set":
			dataSource.ReplicaSet = request.DataSource.ReplicaSet
		case "direct_connection":
			dataSource.DirectConnection = request.DataSource.DirectConnection
		case "region":
			dataSource.Region = request.DataSource.Region
		case "warehouse_id":
			dataSource.WarehouseId = request.DataSource.WarehouseId
		case "use_ssl":
			dataSource.UseSsl = request.DataSource.UseSsl
		case "redis_type":
			dataSource.RedisType = convertV1RedisType(request.DataSource.RedisType)
		case "master_name":
			dataSource.MasterName = request.DataSource.MasterName
		case "master_username":
			dataSource.MasterUsername = request.DataSource.MasterUsername
		case "master_password":
			dataSource.ObfuscatedMasterPassword = common.Obfuscate(request.DataSource.MasterPassword, s.secret)
		case "iam_extension":
			if v := request.DataSource.IamExtension; v != nil {
				switch v := v.(type) {
				case *v1pb.DataSource_ClientSecretCredential_:
					v1ClientSecretCredential := v.ClientSecretCredential
					v1ClientSecretCredential.ClientSecret = common.Obfuscate(v1ClientSecretCredential.ClientSecret, s.secret)
					dataSource.IamExtension = &storepb.DataSource_ClientSecretCredential_{
						ClientSecretCredential: s.convertV1ClientSecretCredential(v.ClientSecretCredential),
					}
				default:
				}
			}
		// TODO(zp): Remove the hack while frontend use new oneof artifact.
		case "client_secret_credential":
			if request.DataSource.GetClientSecretCredential() == nil {
				dataSource.IamExtension = nil
			} else {
				v1ClientSecretCredential := request.DataSource.GetClientSecretCredential()
				v1ClientSecretCredential.ClientSecret = common.Obfuscate(v1ClientSecretCredential.ClientSecret, s.secret)
				dataSource.IamExtension = &storepb.DataSource_ClientSecretCredential_{
					ClientSecretCredential: s.convertV1ClientSecretCredential(request.DataSource.GetClientSecretCredential()),
				}
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, `unsupported update_mask "%s"`, path)
		}
	}

	if err := s.checkDataSource(instance, dataSource); err != nil {
		return nil, err
	}
	if hasSSH {
		if err := s.licenseService.IsFeatureEnabledForInstance(api.FeatureInstanceSSHConnection, instance); err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
	}

	// Test connection.
	if request.ValidateOnly {
		err := func() error {
			driver, err := s.dbFactory.GetDataSourceDriver(ctx, instance, dataSource, "", false /* datashare */, dataSource.GetType() == storepb.DataSourceType_READ_ONLY, db.ConnectionContext{})
			if err != nil {
				return status.Errorf(codes.Internal, "failed to get database driver with error: %v", err.Error())
			}
			defer driver.Close(ctx)
			if err := driver.Ping(ctx); err != nil {
				return status.Errorf(codes.InvalidArgument, "invalid datasource %s, error %s", dataSource.GetType(), err)
			}
			return nil
		}()
		if err != nil {
			return nil, err
		}
		return convertInstanceMessage(instance)
	}

	instance, err = s.store.UpdateInstanceV2(ctx, &store.UpdateInstanceMessage{ResourceID: instance.ResourceID, Metadata: metadata})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return convertInstanceMessage(instance)
}

// RemoveDataSource removes a data source to an instance.
func (s *InstanceService) RemoveDataSource(ctx context.Context, request *v1pb.RemoveDataSourceRequest) (*v1pb.Instance, error) {
	if request.DataSource == nil {
		return nil, status.Errorf(codes.InvalidArgument, "data sources is required")
	}

	instance, err := getInstanceMessage(ctx, s.store, request.Name)
	if err != nil {
		return nil, err
	}
	if instance.Deleted {
		return nil, status.Errorf(codes.NotFound, "instance %q has been deleted", request.Name)
	}

	metadata, ok := proto.Clone(instance.Metadata).(*storepb.Instance)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to convert instance metadata type")
	}
	var updatedDataSources []*storepb.DataSource
	var dataSource *storepb.DataSource
	for _, ds := range instance.Metadata.GetDataSources() {
		if ds.GetId() == request.DataSource.Id {
			dataSource = ds
		} else {
			updatedDataSources = append(updatedDataSources, ds)
		}
	}
	if dataSource == nil {
		return nil, status.Errorf(codes.NotFound, "data source not found")
	}

	// We only support remove RO type datasource to instance now, see more details in instance_service.proto.
	if dataSource.GetType() != storepb.DataSourceType_READ_ONLY {
		return nil, status.Errorf(codes.InvalidArgument, "only support remove read-only data source")
	}

	metadata.DataSources = updatedDataSources
	instance, err = s.store.UpdateInstanceV2(ctx, &store.UpdateInstanceMessage{ResourceID: instance.ResourceID, Metadata: metadata})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	instance, err = s.store.GetInstanceV2(ctx, &store.FindInstanceMessage{
		ResourceID: &instance.ResourceID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return convertInstanceMessage(instance)
}

func (s *InstanceService) getProjectMessage(ctx context.Context, name string) (*store.ProjectMessage, error) {
	projectID, err := common.GetProjectID(name)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	project, err := s.store.GetProjectV2(ctx, &store.FindProjectMessage{
		ResourceID:  &projectID,
		ShowDeleted: true,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if project == nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name)
	}

	return project, nil
}

func getInstanceMessage(ctx context.Context, stores *store.Store, name string) (*store.InstanceMessage, error) {
	instanceID, err := common.GetInstanceID(name)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	find := &store.FindInstanceMessage{
		ResourceID: &instanceID,
	}
	instance, err := stores.GetInstanceV2(ctx, find)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if instance == nil {
		return nil, status.Errorf(codes.NotFound, "instance %q not found", name)
	}

	return instance, nil
}

// buildInstanceName builds the instance name with the given instance ID.
func buildInstanceName(instanceID string) string {
	var b strings.Builder
	b.Grow(len(common.InstanceNamePrefix) + len(instanceID))
	_, _ = b.WriteString(common.InstanceNamePrefix)
	_, _ = b.WriteString(instanceID)
	return b.String()
}

// buildEnvironmentName builds the environment name with the given environment ID.
func buildEnvironmentName(environmentID string) string {
	var b strings.Builder
	b.Grow(len("environments/") + len(environmentID))
	_, _ = b.WriteString("environments/")
	_, _ = b.WriteString(environmentID)
	return b.String()
}

func convertInstanceMessage(instance *store.InstanceMessage) (*v1pb.Instance, error) {
	engine := convertToEngine(instance.Metadata.GetEngine())
	dataSources, err := convertDataSources(instance.Metadata.GetDataSources())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert data source with error: %v", err.Error())
	}

	return &v1pb.Instance{
		Name:               buildInstanceName(instance.ResourceID),
		Title:              instance.Metadata.GetTitle(),
		Engine:             engine,
		EngineVersion:      instance.Metadata.GetVersion(),
		ExternalLink:       instance.Metadata.GetExternalLink(),
		DataSources:        dataSources,
		State:              convertDeletedToState(instance.Deleted),
		Environment:        buildEnvironmentName(instance.EnvironmentID),
		Activation:         instance.Metadata.GetActivation(),
		SyncInterval:       instance.Metadata.GetSyncInterval(),
		MaximumConnections: instance.Metadata.GetMaximumConnections(),
		SyncDatabases:      instance.Metadata.GetSyncDatabases(),
		Roles:              convertInstanceRoles(instance, instance.Metadata.GetRoles()),
	}, nil
}

// buildRoleName builds the role name with the given instance ID and role name.
func buildRoleName(b *strings.Builder, instanceID, roleName string) string {
	b.Reset()
	_, _ = b.WriteString(common.InstanceNamePrefix)
	_, _ = b.WriteString(instanceID)
	_, _ = b.WriteString("/")
	_, _ = b.WriteString(common.RolePrefix)
	_, _ = b.WriteString(roleName)
	return b.String()
}

func convertInstanceRoles(instance *store.InstanceMessage, roles []*storepb.InstanceRole) []*v1pb.InstanceRole {
	var v1Roles []*v1pb.InstanceRole
	var b strings.Builder

	// preallocate memory for the builder
	b.Grow(len(common.InstanceNamePrefix) + len(instance.ResourceID) + 1 + len(common.RolePrefix) + 20)

	for _, role := range roles {
		v1Roles = append(v1Roles, &v1pb.InstanceRole{
			Name:      buildRoleName(&b, instance.ResourceID, role.Name),
			RoleName:  role.Name,
			Attribute: role.Attribute,
		})
	}
	return v1Roles
}

func (s *InstanceService) convertInstanceToInstanceMessage(instanceID string, instance *v1pb.Instance) (*store.InstanceMessage, error) {
	datasources, err := s.convertV1DataSources(instance.DataSources)
	if err != nil {
		return nil, err
	}
	environmentID, err := common.GetEnvironmentID(instance.Environment)
	if err != nil {
		return nil, err
	}

	return &store.InstanceMessage{
		ResourceID:    instanceID,
		EnvironmentID: environmentID,
		Metadata: &storepb.Instance{
			Title:              instance.GetTitle(),
			Engine:             convertEngine(instance.Engine),
			ExternalLink:       instance.GetExternalLink(),
			Activation:         instance.GetActivation(),
			DataSources:        datasources,
			SyncInterval:       instance.GetSyncInterval(),
			MaximumConnections: instance.GetMaximumConnections(),
			SyncDatabases:      instance.GetSyncDatabases(),
		},
	}, nil
}

func convertInstanceMessageToInstanceResource(instanceMessage *store.InstanceMessage) (*v1pb.InstanceResource, error) {
	instance, err := convertInstanceMessage(instanceMessage)
	if err != nil {
		return nil, err
	}
	return &v1pb.InstanceResource{
		Name:          instance.Name,
		Title:         instance.Title,
		Engine:        instance.Engine,
		EngineVersion: instance.EngineVersion,
		DataSources:   instance.DataSources,
		Activation:    instance.Activation,
		Environment:   instance.Environment,
	}, nil
}

func (s *InstanceService) convertV1DataSources(dataSources []*v1pb.DataSource) ([]*storepb.DataSource, error) {
	var values []*storepb.DataSource
	for _, ds := range dataSources {
		dataSource, err := s.convertV1DataSource(ds)
		if err != nil {
			return nil, err
		}
		values = append(values, dataSource)
	}

	return values, nil
}

func convertDataSourceExternalSecret(externalSecret *storepb.DataSourceExternalSecret) (*v1pb.DataSourceExternalSecret, error) {
	if externalSecret == nil {
		return nil, nil
	}
	secret := new(v1pb.DataSourceExternalSecret)
	if err := convertProtoToProto(externalSecret, secret); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert external secret with error: %v", err.Error())
	}

	resp := &v1pb.DataSourceExternalSecret{
		SecretType:      secret.SecretType,
		Url:             secret.Url,
		AuthType:        secret.AuthType,
		EngineName:      secret.EngineName,
		SecretName:      secret.SecretName,
		PasswordKeyName: secret.PasswordKeyName,
	}

	// clear sensitive data.
	switch resp.AuthType {
	case v1pb.DataSourceExternalSecret_VAULT_APP_ROLE:
		appRole := secret.GetAppRole()
		resp.AuthOption = &v1pb.DataSourceExternalSecret_AppRole{
			AppRole: &v1pb.DataSourceExternalSecret_AppRoleAuthOption{
				Type:      appRole.Type,
				MountPath: appRole.MountPath,
			},
		}
	case v1pb.DataSourceExternalSecret_TOKEN:
		resp.AuthOption = &v1pb.DataSourceExternalSecret_Token{
			Token: "",
		}
	}

	return resp, nil
}

func convertDataSources(dataSources []*storepb.DataSource) ([]*v1pb.DataSource, error) {
	var v1DataSources []*v1pb.DataSource
	for _, ds := range dataSources {
		externalSecret, err := convertDataSourceExternalSecret(ds.GetExternalSecret())
		if err != nil {
			return nil, err
		}

		dataSourceType := v1pb.DataSourceType_DATA_SOURCE_UNSPECIFIED
		switch ds.GetType() {
		case storepb.DataSourceType_ADMIN:
			dataSourceType = v1pb.DataSourceType_ADMIN
		case storepb.DataSourceType_READ_ONLY:
			dataSourceType = v1pb.DataSourceType_READ_ONLY
		}

		authenticationType := v1pb.DataSource_AUTHENTICATION_UNSPECIFIED
		switch ds.GetAuthenticationType() {
		case storepb.DataSource_AUTHENTICATION_UNSPECIFIED, storepb.DataSource_PASSWORD:
			authenticationType = v1pb.DataSource_PASSWORD
		case storepb.DataSource_GOOGLE_CLOUD_SQL_IAM:
			authenticationType = v1pb.DataSource_GOOGLE_CLOUD_SQL_IAM
		case storepb.DataSource_AWS_RDS_IAM:
			authenticationType = v1pb.DataSource_AWS_RDS_IAM
		case storepb.DataSource_AZURE_IAM:
			authenticationType = v1pb.DataSource_AZURE_IAM
		}

		dataSource := &v1pb.DataSource{
			Id:       ds.GetId(),
			Type:     dataSourceType,
			Username: ds.GetUsername(),
			// We don't return the password and SSLs on reads.
			Host:                   ds.GetHost(),
			Port:                   ds.GetPort(),
			Database:               ds.GetDatabase(),
			Srv:                    ds.GetSrv(),
			AuthenticationDatabase: ds.GetAuthenticationDatabase(),
			Sid:                    ds.GetSid(),
			ServiceName:            ds.GetServiceName(),
			ExternalSecret:         externalSecret,
			AuthenticationType:     authenticationType,
			SaslConfig:             convertDataSourceSaslConfig(ds.GetSaslConfig()),
			AdditionalAddresses:    convertDataSourceAddresses(ds.GetAdditionalAddresses()),
			ReplicaSet:             ds.GetReplicaSet(),
			DirectConnection:       ds.GetDirectConnection(),
			Region:                 ds.GetRegion(),
			WarehouseId:            ds.GetWarehouseId(),
			UseSsl:                 ds.GetUseSsl(),
			RedisType:              convertRedisType(ds.GetRedisType()),
			MasterName:             ds.GetMasterName(),
			MasterUsername:         ds.GetMasterUsername(),
		}
		if clientSecretCredential := convertClientSecretCredential(ds.GetClientSecretCredential()); clientSecretCredential != nil {
			clientSecretCredential.ClientSecret = ""
			dataSource.IamExtension = &v1pb.DataSource_ClientSecretCredential_{
				ClientSecretCredential: clientSecretCredential,
			}
		}

		v1DataSources = append(v1DataSources, dataSource)
	}

	return v1DataSources, nil
}

func convertClientSecretCredential(clientSecretCredential *storepb.DataSource_ClientSecretCredential) *v1pb.DataSource_ClientSecretCredential {
	if clientSecretCredential == nil {
		return nil
	}
	return &v1pb.DataSource_ClientSecretCredential{
		TenantId: clientSecretCredential.TenantId,
		ClientId: clientSecretCredential.ClientId,
	}
}

func convertV1DataSourceExternalSecret(externalSecret *v1pb.DataSourceExternalSecret) (*storepb.DataSourceExternalSecret, error) {
	if externalSecret == nil {
		return nil, nil
	}
	secret := new(storepb.DataSourceExternalSecret)
	if err := convertProtoToProto(externalSecret, secret); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert external secret with error: %v", err.Error())
	}
	switch secret.SecretType {
	case storepb.DataSourceExternalSecret_VAULT_KV_V2:
		if secret.Url == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing Vault URL")
		}
		if secret.EngineName == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing Vault engine name")
		}
		if secret.SecretName == "" || secret.PasswordKeyName == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing secret name or key name")
		}
	case storepb.DataSourceExternalSecret_AWS_SECRETS_MANAGER:
		if secret.SecretName == "" || secret.PasswordKeyName == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing secret name or key name")
		}
	case storepb.DataSourceExternalSecret_GCP_SECRET_MANAGER:
		if secret.SecretName == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing GCP secret name")
		}
	}

	switch secret.AuthType {
	case storepb.DataSourceExternalSecret_TOKEN:
		if secret.GetToken() == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing token")
		}
	case storepb.DataSourceExternalSecret_VAULT_APP_ROLE:
		if secret.GetAppRole() == nil {
			return nil, status.Errorf(codes.InvalidArgument, "missing Vault approle")
		}
	}

	return secret, nil
}

func convertV1DataSourceSaslConfig(saslConfig *v1pb.SASLConfig) *storepb.SASLConfig {
	if saslConfig == nil {
		return nil
	}
	storeSaslConfig := &storepb.SASLConfig{}
	switch m := saslConfig.Mechanism.(type) {
	case *v1pb.SASLConfig_KrbConfig:
		storeSaslConfig.Mechanism = &storepb.SASLConfig_KrbConfig{
			KrbConfig: &storepb.KerberosConfig{
				Primary:              m.KrbConfig.Primary,
				Instance:             m.KrbConfig.Instance,
				Realm:                m.KrbConfig.Realm,
				Keytab:               m.KrbConfig.Keytab,
				KdcHost:              m.KrbConfig.KdcHost,
				KdcPort:              m.KrbConfig.KdcPort,
				KdcTransportProtocol: m.KrbConfig.KdcTransportProtocol,
			},
		}
	default:
		return nil
	}
	return storeSaslConfig
}

func convertDataSourceSaslConfig(saslConfig *storepb.SASLConfig) *v1pb.SASLConfig {
	if saslConfig == nil {
		return nil
	}
	storeSaslConfig := &v1pb.SASLConfig{}
	switch m := saslConfig.Mechanism.(type) {
	case *storepb.SASLConfig_KrbConfig:
		storeSaslConfig.Mechanism = &v1pb.SASLConfig_KrbConfig{
			KrbConfig: &v1pb.KerberosConfig{
				Primary:              m.KrbConfig.Primary,
				Instance:             m.KrbConfig.Instance,
				Realm:                m.KrbConfig.Realm,
				Keytab:               m.KrbConfig.Keytab,
				KdcHost:              m.KrbConfig.KdcHost,
				KdcPort:              m.KrbConfig.KdcPort,
				KdcTransportProtocol: m.KrbConfig.KdcTransportProtocol,
			},
		}
	default:
		return nil
	}
	return storeSaslConfig
}

func convertDataSourceAddresses(addresses []*storepb.DataSource_Address) []*v1pb.DataSource_Address {
	res := make([]*v1pb.DataSource_Address, 0, len(addresses))
	for _, address := range addresses {
		res = append(res, &v1pb.DataSource_Address{
			Host: address.Host,
			Port: address.Port,
		})
	}
	return res
}

func convertAdditionalAddresses(addresses []*v1pb.DataSource_Address) []*storepb.DataSource_Address {
	res := make([]*storepb.DataSource_Address, 0, len(addresses))
	for _, address := range addresses {
		res = append(res, &storepb.DataSource_Address{
			Host: address.Host,
			Port: address.Port,
		})
	}
	return res
}

func convertV1AuthenticationType(authType v1pb.DataSource_AuthenticationType) storepb.DataSource_AuthenticationType {
	authenticationType := storepb.DataSource_AUTHENTICATION_UNSPECIFIED
	switch authType {
	case v1pb.DataSource_AUTHENTICATION_UNSPECIFIED, v1pb.DataSource_PASSWORD:
		authenticationType = storepb.DataSource_PASSWORD
	case v1pb.DataSource_GOOGLE_CLOUD_SQL_IAM:
		authenticationType = storepb.DataSource_GOOGLE_CLOUD_SQL_IAM
	case v1pb.DataSource_AWS_RDS_IAM:
		authenticationType = storepb.DataSource_AWS_RDS_IAM
	case v1pb.DataSource_AZURE_IAM:
		authenticationType = storepb.DataSource_AZURE_IAM
	}
	return authenticationType
}

func convertV1RedisType(redisType v1pb.DataSource_RedisType) storepb.DataSource_RedisType {
	authenticationType := storepb.DataSource_REDIS_TYPE_UNSPECIFIED
	switch redisType {
	case v1pb.DataSource_STANDALONE:
		authenticationType = storepb.DataSource_STANDALONE
	case v1pb.DataSource_SENTINEL:
		authenticationType = storepb.DataSource_SENTINEL
	case v1pb.DataSource_CLUSTER:
		authenticationType = storepb.DataSource_CLUSTER
	}
	return authenticationType
}

func convertRedisType(redisType storepb.DataSource_RedisType) v1pb.DataSource_RedisType {
	authenticationType := v1pb.DataSource_STANDALONE
	switch redisType {
	case storepb.DataSource_STANDALONE:
		authenticationType = v1pb.DataSource_STANDALONE
	case storepb.DataSource_SENTINEL:
		authenticationType = v1pb.DataSource_SENTINEL
	case storepb.DataSource_CLUSTER:
		authenticationType = v1pb.DataSource_CLUSTER
	}
	return authenticationType
}

func (s *InstanceService) convertV1DataSource(dataSource *v1pb.DataSource) (*storepb.DataSource, error) {
	dsType, err := convertV1DataSourceType(dataSource.Type)
	if err != nil {
		return nil, err
	}
	externalSecret, err := convertV1DataSourceExternalSecret(dataSource.ExternalSecret)
	if err != nil {
		return nil, err
	}
	saslConfig := convertV1DataSourceSaslConfig(dataSource.SaslConfig)

	storeDataSource := &storepb.DataSource{
		Id:                                 dataSource.Id,
		Type:                               dsType,
		Username:                           dataSource.Username,
		ObfuscatedPassword:                 common.Obfuscate(dataSource.Password, s.secret),
		ObfuscatedSslCa:                    common.Obfuscate(dataSource.SslCa, s.secret),
		ObfuscatedSslCert:                  common.Obfuscate(dataSource.SslCert, s.secret),
		ObfuscatedSslKey:                   common.Obfuscate(dataSource.SslKey, s.secret),
		Host:                               dataSource.Host,
		Port:                               dataSource.Port,
		Database:                           dataSource.Database,
		Srv:                                dataSource.Srv,
		AuthenticationDatabase:             dataSource.AuthenticationDatabase,
		Sid:                                dataSource.Sid,
		ServiceName:                        dataSource.ServiceName,
		SshHost:                            dataSource.SshHost,
		SshPort:                            dataSource.SshPort,
		SshUser:                            dataSource.SshUser,
		ObfuscatedSshPassword:              common.Obfuscate(dataSource.SshPassword, s.secret),
		ObfuscatedSshPrivateKey:            common.Obfuscate(dataSource.SshPrivateKey, s.secret),
		ObfuscatedAuthenticationPrivateKey: common.Obfuscate(dataSource.AuthenticationPrivateKey, s.secret),
		ExternalSecret:                     externalSecret,
		SaslConfig:                         saslConfig,
		AuthenticationType:                 convertV1AuthenticationType(dataSource.AuthenticationType),
		AdditionalAddresses:                convertAdditionalAddresses(dataSource.AdditionalAddresses),
		ReplicaSet:                         dataSource.ReplicaSet,
		DirectConnection:                   dataSource.DirectConnection,
		Region:                             dataSource.Region,
		WarehouseId:                        dataSource.WarehouseId,
		UseSsl:                             dataSource.UseSsl,
		RedisType:                          convertV1RedisType(dataSource.RedisType),
		MasterName:                         dataSource.MasterName,
		MasterUsername:                     dataSource.MasterUsername,
		ObfuscatedMasterPassword:           common.Obfuscate(dataSource.MasterPassword, s.secret),
	}
	if v := dataSource.GetClientSecretCredential(); v != nil {
		v.ClientSecret = common.Obfuscate(v.ClientSecret, s.secret)
		storeDataSource.IamExtension = &storepb.DataSource_ClientSecretCredential_{ClientSecretCredential: s.convertV1ClientSecretCredential(v)}
	}

	return storeDataSource, nil
}

func (s *InstanceService) convertV1ClientSecretCredential(credential *v1pb.DataSource_ClientSecretCredential) *storepb.DataSource_ClientSecretCredential {
	if credential == nil {
		return nil
	}
	return &storepb.DataSource_ClientSecretCredential{
		TenantId:               credential.TenantId,
		ClientId:               credential.ClientId,
		ObfuscatedClientSecret: common.Obfuscate(credential.ClientSecret, s.secret),
	}
}

func convertV1DataSourceType(tp v1pb.DataSourceType) (storepb.DataSourceType, error) {
	switch tp {
	case v1pb.DataSourceType_READ_ONLY:
		return storepb.DataSourceType_READ_ONLY, nil
	case v1pb.DataSourceType_ADMIN:
		return storepb.DataSourceType_ADMIN, nil
	default:
		return storepb.DataSourceType_DATA_SOURCE_UNSPECIFIED, errors.Errorf("invalid data source type %v", tp)
	}
}

func (s *InstanceService) instanceCountGuard(ctx context.Context) error {
	instanceLimit := s.licenseService.GetPlanLimitValue(ctx, enterprise.PlanLimitMaximumInstance)

	count, err := s.store.CountInstance(ctx, &store.CountInstanceMessage{})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if count >= instanceLimit {
		return status.Errorf(codes.ResourceExhausted, "reached the maximum instance count %d", instanceLimit)
	}

	return nil
}
