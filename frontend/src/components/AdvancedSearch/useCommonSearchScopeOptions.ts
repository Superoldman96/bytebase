import type { VNode } from "vue";
import { computed, h, unref } from "vue";
import {
  InstanceV1Name,
  ProjectV1Name,
  EnvironmentV1Name,
  RichEngineName,
} from "@/components/v2";
import { t } from "@/plugins/i18n";
import {
  environmentNamePrefix,
  useEnvironmentV1List,
  useEnvironmentV1Store,
  useInstanceV1Store,
  useProjectV1Store,
} from "@/store";
import type { MaybeRef } from "@/types";
import { engineToJSON } from "@/types/proto/v1/common";
import type { SearchScopeId } from "@/utils";
import {
  environmentV1Name,
  extractEnvironmentResourceName,
  extractInstanceResourceName,
  extractProjectResourceName,
  supportedEngineV1List,
  getDefaultPagination,
} from "@/utils";
import type { ScopeOption, ValueOption } from "./types";

export const useCommonSearchScopeOptions = (
  supportOptionIdList: MaybeRef<SearchScopeId[]>
) => {
  const projectStore = useProjectV1Store();
  const instanceStore = useInstanceV1Store();
  const environmentStore = useEnvironmentV1Store();
  const environmentList = useEnvironmentV1List();

  // fullScopeOptions provides full search scopes and options.
  // we need this as the source of truth.
  const fullScopeOptions = computed((): ScopeOption[] => {
    const scopeCreators = {
      project: () => ({
        id: "project",
        title: t("issue.advanced-search.scope.project.title"),
        description: t("issue.advanced-search.scope.project.description"),
        search: ({
          keyword,
          nextPageToken,
        }: {
          keyword: string;
          nextPageToken?: string;
        }) => {
          return projectStore
            .fetchProjectList({
              pageToken: nextPageToken,
              pageSize: getDefaultPagination(),
              filter: {
                query: keyword,
              },
            })
            .then((resp) => ({
              nextPageToken: resp.nextPageToken,
              options: resp.projects.map<ValueOption>((project) => {
                const name = extractProjectResourceName(project.name);
                return {
                  value: name,
                  keywords: [
                    name,
                    project.title,
                    extractProjectResourceName(project.name),
                  ],
                  render: () => {
                    const children: VNode[] = [
                      h(ProjectV1Name, { project: project, link: false }),
                    ];
                    return h(
                      "div",
                      { class: "flex items-center gap-x-2" },
                      children
                    );
                  },
                };
              }),
            }));
        },
      }),
      instance: () => ({
        id: "instance",
        title: t("issue.advanced-search.scope.instance.title"),
        description: t("issue.advanced-search.scope.instance.description"),
        search: ({
          keyword,
          nextPageToken,
        }: {
          keyword: string;
          nextPageToken?: string;
        }) => {
          return instanceStore
            .fetchInstanceList({
              pageToken: nextPageToken,
              pageSize: getDefaultPagination(),
              filter: {
                query: keyword,
              },
            })
            .then((resp) => ({
              nextPageToken: resp.nextPageToken,
              options: resp.instances.map((ins) => {
                const name = extractInstanceResourceName(ins.name);
                return {
                  value: name,
                  keywords: [
                    name,
                    ins.title,
                    engineToJSON(ins.engine),
                    extractEnvironmentResourceName(ins.environment),
                  ],
                  render: () => {
                    return h("div", { class: "flex items-center gap-x-1" }, [
                      h(InstanceV1Name, {
                        instance: ins,
                        link: false,
                        tooltip: false,
                      }),
                      h(
                        "span",
                        {},
                        `(${environmentV1Name(environmentStore.getEnvironmentByName(ins.environment))})`
                      ),
                    ]);
                  },
                };
              }),
            }));
        },
      }),
      environment: () => ({
        id: "environment",
        title: t("issue.advanced-search.scope.environment.title"),
        description: t("issue.advanced-search.scope.environment.description"),
        options: environmentList.value.map((env) => {
          return {
            value: env.id,
            keywords: [`${environmentNamePrefix}${env.id}`, env.title],
            render: () =>
              h(EnvironmentV1Name, {
                environment: env,
                link: false,
              }),
          };
        }),
      }),
      "database-label": () => ({
        id: "database-label",
        title: t("issue.advanced-search.scope.database-label.title"),
        description: t(
          "issue.advanced-search.scope.database-label.description"
        ),
        allowMultiple: true,
      }),
      engine: () => ({
        id: "engine",
        title: t("issue.advanced-search.scope.engine.title"),
        description: t("issue.advanced-search.scope.engine.description"),
        options: supportedEngineV1List().map((engine) => {
          return {
            value: engine,
            keywords: [engineToJSON(engine).toLowerCase()],
            render: () => h(RichEngineName, { engine, tag: "p" }),
          };
        }),
        allowMultiple: true,
      }),
    } as Record<SearchScopeId, () => ScopeOption>;

    const scopes: ScopeOption[] = [];
    unref(supportOptionIdList).forEach((id) => {
      const create = scopeCreators[id];
      if (create) {
        scopes.push(create());
      }
    });
    return scopes;
  });

  return fullScopeOptions;
};
