import { uniqBy } from "lodash-es";
import { computed } from "vue";
import { unknownInstanceResource } from "@/types";
import type { InstanceResource } from "@/types/proto/v1/instance_service";
import { useDatabaseV1Store } from "./database";
import { useInstanceV1List } from "./instance";

// Instance resource list is a list of all instance resources in the database list.
// Current user should have access to all instance resources in the list.
export const useInstanceResourceList = () => {
  const { instanceList } = useInstanceV1List();
  const { databaseList } = useDatabaseV1Store();
  return computed(() => {
    return uniqBy(
      [
        ...databaseList.map((db) => db.instanceResource),
        // Merge possible instance resources from the instance store.
        ...instanceList.value,
      ],
      (i) => i.name
    ) as InstanceResource[];
  });
};

// Instance resource list is a list of all instance resources in the database list.
// Current user should have access to all instance resources in the list.
export const useInstanceResourceByName = (
  instanceName: string // Format: instances/{instance}
) => {
  return (
    useInstanceResourceList().value.find((i) => i.name === instanceName) ||
    unknownInstanceResource()
  );
};
