// Package store is the implementation for managing Bytebase's own metadata in a PostgreSQL database.
package store

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/pkg/errors"

	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/store/model"
)

// Store provides database access to all raw objects.
type Store struct {
	db          *sql.DB
	enableCache bool

	// Cache.
	Secret               string
	userIDCache          *lru.Cache[int, *UserMessage]
	userEmailCache       *lru.Cache[string, *UserMessage]
	instanceCache        *lru.Cache[string, *InstanceMessage]
	databaseCache        *lru.Cache[string, *DatabaseMessage]
	projectCache         *lru.Cache[string, *ProjectMessage]
	policyCache          *lru.Cache[string, *PolicyMessage]
	issueCache           *lru.Cache[int, *IssueMessage]
	issueByPipelineCache *lru.Cache[int, *IssueMessage]
	pipelineCache        *lru.Cache[int, *PipelineMessage]
	settingCache         *lru.Cache[storepb.SettingName, *SettingMessage]
	idpCache             *lru.Cache[string, *IdentityProviderMessage]
	risksCache           *lru.Cache[int, []*RiskMessage] // Use 0 as the key.
	databaseGroupCache   *lru.Cache[string, *DatabaseGroupMessage]
	rolesCache           *lru.Cache[string, *RoleMessage]
	groupCache           *lru.Cache[string, *GroupMessage]
	sheetCache           *lru.Cache[int, *SheetMessage]

	// Large objects.
	sheetStatementCache *lru.Cache[int, string]
	dbSchemaCache       *lru.Cache[string, *model.DatabaseSchema]
}

// New creates a new instance of Store.
func New(ctx context.Context, pgURL string, enableCache bool) (*Store, error) {
	userIDCache, err := lru.New[int, *UserMessage](32768)
	if err != nil {
		return nil, err
	}
	userEmailCache, err := lru.New[string, *UserMessage](32768)
	if err != nil {
		return nil, err
	}
	instanceCache, err := lru.New[string, *InstanceMessage](32768)
	if err != nil {
		return nil, err
	}
	databaseCache, err := lru.New[string, *DatabaseMessage](32768)
	if err != nil {
		return nil, err
	}
	projectCache, err := lru.New[string, *ProjectMessage](32768)
	if err != nil {
		return nil, err
	}
	policyCache, err := lru.New[string, *PolicyMessage](128)
	if err != nil {
		return nil, err
	}
	issueCache, err := lru.New[int, *IssueMessage](256)
	if err != nil {
		return nil, err
	}
	issueByPipelineCache, err := lru.New[int, *IssueMessage](256)
	if err != nil {
		return nil, err
	}
	pipelineCache, err := lru.New[int, *PipelineMessage](256)
	if err != nil {
		return nil, err
	}
	settingCache, err := lru.New[storepb.SettingName, *SettingMessage](64)
	if err != nil {
		return nil, err
	}
	idpCache, err := lru.New[string, *IdentityProviderMessage](4)
	if err != nil {
		return nil, err
	}
	risksCache, err := lru.New[int, []*RiskMessage](4)
	if err != nil {
		return nil, err
	}
	databaseGroupCache, err := lru.New[string, *DatabaseGroupMessage](1024)
	if err != nil {
		return nil, err
	}
	rolesCache, err := lru.New[string, *RoleMessage](64)
	if err != nil {
		return nil, err
	}
	sheetCache, err := lru.New[int, *SheetMessage](64)
	if err != nil {
		return nil, err
	}
	sheetStatementCache, err := lru.New[int, string](10)
	if err != nil {
		return nil, err
	}
	dbSchemaCache, err := lru.New[string, *model.DatabaseSchema](128)
	if err != nil {
		return nil, err
	}
	groupCache, err := lru.New[string, *GroupMessage](1024)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("pgx", pgURL)
	if err != nil {
		return nil, err
	}

	// Set the max open connections so that we won't exceed the connection limit of metaDB.
	// The limit is the max connections minus connections reserved for superuser.
	var maxConns, reservedConns int
	if err := db.QueryRowContext(ctx, `SHOW max_connections`).Scan(&maxConns); err != nil {
		return nil, errors.Wrap(err, "failed to get max_connections from metaDB")
	}
	if err := db.QueryRowContext(ctx, `SHOW superuser_reserved_connections`).Scan(&reservedConns); err != nil {
		return nil, errors.Wrap(err, "failed to get superuser_reserved_connections from metaDB")
	}
	maxOpenConns := maxConns - reservedConns
	if maxOpenConns > 50 {
		// capped to 50
		maxOpenConns = 50
	}
	db.SetMaxOpenConns(maxOpenConns)

	return &Store{
		db:          db,
		enableCache: enableCache,

		// Cache.
		userIDCache:          userIDCache,
		userEmailCache:       userEmailCache,
		instanceCache:        instanceCache,
		databaseCache:        databaseCache,
		projectCache:         projectCache,
		policyCache:          policyCache,
		issueCache:           issueCache,
		issueByPipelineCache: issueByPipelineCache,
		pipelineCache:        pipelineCache,
		settingCache:         settingCache,
		idpCache:             idpCache,
		risksCache:           risksCache,
		databaseGroupCache:   databaseGroupCache,
		rolesCache:           rolesCache,
		sheetCache:           sheetCache,
		sheetStatementCache:  sheetStatementCache,
		dbSchemaCache:        dbSchemaCache,
		groupCache:           groupCache,
	}, nil
}

// Close closes underlying db.
func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

func getInstanceCacheKey(instanceID string) string {
	return instanceID
}

func getPolicyCacheKey(resourceType storepb.Policy_Resource, resource string, policyType storepb.Policy_Type) string {
	return fmt.Sprintf("policies/%s/%s/%s", resourceType, resource, policyType)
}

func getDatabaseCacheKey(instanceID, databaseName string) string {
	return fmt.Sprintf("%s/%s", instanceID, databaseName)
}

func getDatabaseGroupCacheKey(projectID, resourceID string) string {
	return fmt.Sprintf("%s/%s", projectID, resourceID)
}

func getPlaceholders(start int, count int) string {
	var placeholders []string
	for i := start; i < start+count; i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}
	return fmt.Sprintf("(%s)", strings.Join(placeholders, ","))
}
