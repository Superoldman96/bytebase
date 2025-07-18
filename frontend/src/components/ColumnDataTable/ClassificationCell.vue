<template>
  <div class="flex items-center gap-x-1">
    <ClassificationLevelBadge
      :classification="classification"
      :classification-config="classificationConfig"
    />
    <template v-if="!readonly && !disabled && !setClassificationFromComment">
      <NPopconfirm v-if="classification" @positive-click="removeClassification">
        <template #trigger>
          <MiniActionButton>
            <XIcon class="w-3 h-3" />
          </MiniActionButton>
        </template>
        <template #default>
          <div>
            {{ $t("settings.sensitive-data.remove-classification-tips") }}
          </div>
        </template>
      </NPopconfirm>
      <MiniActionButton v-if="classificationConfig" @click.prevent="openDrawer">
        <PencilIcon class="w-3 h-3" />
      </MiniActionButton>
    </template>
  </div>

  <SelectClassificationDrawer
    v-if="classificationConfig"
    :show="showClassificationDrawer"
    :classification-config="classificationConfig"
    @dismiss="showClassificationDrawer = false"
    @apply="$emit('apply', $event)"
  />
</template>

<script lang="ts" setup>
import { PencilIcon, XIcon } from "lucide-vue-next";
import { NPopconfirm } from "naive-ui";
import { ref, computed } from "vue";
import ClassificationLevelBadge from "@/components/SchemaTemplate/ClassificationLevelBadge.vue";
import { MiniActionButton } from "@/components/v2";
import type { Engine } from "@/types/proto-es/v1/common_pb";
import type { DataClassificationSetting_DataClassificationConfig as DataClassificationConfig } from "@/types/proto-es/v1/setting_service_pb";
import SelectClassificationDrawer from "../SchemaTemplate/SelectClassificationDrawer.vue";
import { supportSetClassificationFromComment } from "./utils";

const props = defineProps<{
  classification?: string | undefined;
  readonly?: boolean;
  disabled?: boolean;
  classificationConfig?: DataClassificationConfig;
  engine: Engine;
}>();

const emit = defineEmits<{
  (event: "apply", id: string): void;
}>();

const showClassificationDrawer = ref(false);

const openDrawer = (e: MouseEvent) => {
  e.stopPropagation();
  showClassificationDrawer.value = true;
};

const removeClassification = (e: MouseEvent) => {
  e.stopPropagation();
  emit("apply", "");
};

const setClassificationFromComment = computed(() => {
  return supportSetClassificationFromComment(
    props.engine,
    props.classificationConfig?.classificationFromConfig ?? false
  );
});
</script>
