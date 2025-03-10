import type { InjectionKey, Ref } from "vue";
import { computed, inject, provide, ref } from "vue";
import { useAppFeature, useDatabaseV1ByName } from "@/store";
import {
  databaseNamePrefix,
  instanceNamePrefix,
} from "@/store/modules/v1/common";
import type { ComposedDatabase, Permission } from "@/types";
import { DEFAULT_PROJECT_NAME } from "@/types";
import {
  hasProjectPermissionV2,
  instanceV1HasAlterSchema,
  isArchivedDatabaseV1,
  hasPermissionToCreateChangeDatabaseIssue,
} from "@/utils";

export type DatabaseDetailContext = {
  database: Ref<ComposedDatabase>;
  pagedRevisionTableSessionKey: Ref<string>;
  allowGetDatabase: Ref<boolean>;
  allowUpdateDatabase: Ref<boolean>;
  allowSyncDatabase: Ref<boolean>;
  allowTransferDatabase: Ref<boolean>;
  allowGetSchema: Ref<boolean>;
  allowChangeData: Ref<boolean>;
  allowAlterSchema: Ref<boolean>;
  allowListSecrets: Ref<boolean>;
  allowUpdateSecrets: Ref<boolean>;
  allowDeleteSecrets: Ref<boolean>;
  allowListChangelogs: Ref<boolean>;
};

export const KEY = Symbol(
  "bb.database.detail"
) as InjectionKey<DatabaseDetailContext>;

export const useDatabaseDetailContext = () => {
  return inject(KEY)!;
};

export const provideDatabaseDetailContext = (
  instanceId: Ref<string>,
  databaseName: Ref<string>
) => {
  const databaseOperations = useAppFeature("bb.feature.databases.operations");

  const { database } = useDatabaseV1ByName(
    computed(
      () =>
        `${instanceNamePrefix}${instanceId.value}/${databaseNamePrefix}${databaseName.value}`
    )
  );

  const pagedRevisionTableSessionKey = ref(
    `bb.paged-revision-table.${Date.now()}`
  );

  const checkPermission = (permission: Permission): boolean => {
    return hasProjectPermissionV2(database.value.projectEntity, permission);
  };

  const allowGetDatabase = computed(() => checkPermission("bb.databases.get"));
  const allowUpdateDatabase = computed(
    () =>
      !isArchivedDatabaseV1(database.value) &&
      checkPermission("bb.databases.update")
  );
  const allowSyncDatabase = computed(() => {
    return (
      databaseOperations.value.has("SYNC-SCHEMA") &&
      checkPermission("bb.databases.sync")
    );
  });
  const allowTransferDatabase = computed(() => {
    if (!databaseOperations.value.has("TRANSFER-OUT")) return false;

    if (database.value.project === DEFAULT_PROJECT_NAME) {
      return true;
    }
    return allowUpdateDatabase.value;
  });

  const allowGetSchema = computed(() =>
    checkPermission("bb.databases.getSchema")
  );

  const allowChangeData = computed(() => {
    return (
      databaseOperations.value.has("CHANGE-DATA") &&
      database.value.project !== DEFAULT_PROJECT_NAME &&
      hasPermissionToCreateChangeDatabaseIssue(database.value)
    );
  });
  const allowAlterSchema = computed(() => {
    return (
      databaseOperations.value.has("EDIT-SCHEMA") &&
      database.value.project !== DEFAULT_PROJECT_NAME &&
      hasPermissionToCreateChangeDatabaseIssue(database.value) &&
      instanceV1HasAlterSchema(database.value.instanceResource)
    );
  });

  const allowListSecrets = computed(() =>
    checkPermission("bb.databaseSecrets.list")
  );
  const allowUpdateSecrets = computed(() =>
    checkPermission("bb.databaseSecrets.update")
  );
  const allowDeleteSecrets = computed(() =>
    checkPermission("bb.databaseSecrets.delete")
  );

  const allowListChangelogs = computed(() =>
    checkPermission("bb.changelogs.list")
  );

  const context: DatabaseDetailContext = {
    database,
    pagedRevisionTableSessionKey,
    allowGetDatabase,
    allowUpdateDatabase,
    allowSyncDatabase,
    allowTransferDatabase,
    allowGetSchema,
    allowChangeData,
    allowAlterSchema,
    allowListSecrets,
    allowUpdateSecrets,
    allowDeleteSecrets,
    allowListChangelogs,
  };

  provide(KEY, context);

  return context;
};
