<template>
  <div class="min-w-[14rem] max-w-[18rem] gap-y-1">
    <InfoItem :title="$t('common.name')">
      {{ name }}
    </InfoItem>
    <InfoItem v-if="engine" :title="$t('database.engine')">
      {{ engine }}
    </InfoItem>
    <InfoItem :title="$t('database.row-count-estimate')">
      {{ table.rowCount }}
    </InfoItem>
    <InfoItem :title="$t('database.data-size')">
      {{ bytesToString(table.dataSize.toNumber()) }}
    </InfoItem>
    <InfoItem v-if="indexSize" :title="$t('database.index-size')">
      {{ indexSize }}
    </InfoItem>
    <InfoItem v-if="collation" :title="$t('db.collation')">
      {{ collation }}
    </InfoItem>
    <InfoItem v-if="comment" :title="$t('database.comment')">
      {{ comment }}
    </InfoItem>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import type { ComposedDatabase } from "@/types";
import { Engine } from "@/types/proto/v1/common";
import type {
  DatabaseMetadata,
  SchemaMetadata,
  TableMetadata,
} from "@/types/proto/v1/database_service";
import { bytesToString, hasSchemaProperty } from "@/utils";
import InfoItem from "./InfoItem.vue";

const props = defineProps<{
  db: ComposedDatabase;
  database: DatabaseMetadata;
  schema: SchemaMetadata;
  table: TableMetadata;
}>();

const instanceEngine = computed(() => props.db.instanceResource.engine);

const name = computed(() => {
  const { schema, table } = props;
  if (hasSchemaProperty(instanceEngine.value)) {
    return `${schema.name}.${table.name}`;
  }
  return table.name;
});

const engine = computed(() => {
  if ([Engine.POSTGRES, Engine.SNOWFLAKE].includes(instanceEngine.value)) {
    return "";
  }
  return props.table.engine;
});

const indexSize = computed(() => {
  if ([Engine.CLICKHOUSE, Engine.SNOWFLAKE].includes(instanceEngine.value)) {
    return "";
  }
  return bytesToString(props.table.indexSize.toNumber());
});

const collation = computed(() => {
  if (
    [Engine.CLICKHOUSE, Engine.SNOWFLAKE, Engine.POSTGRES].includes(
      instanceEngine.value
    )
  ) {
    return "";
  }
  return props.table.collation;
});

const comment = computed(() => {
  return props.table.userComment;
});
</script>
