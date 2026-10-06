<!--
  SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
  SPDX-License-Identifier: Apache-2.0
-->

<template>
  <section class="assignments-modal">
    <runai-base-dialog :model-value="true" @close="$emit('close')" size="width-xl" no-padding>
      <template #header> {{ modalOptions.header }} </template>

      <template #body>
        <runai-list-modal-search
          :model-value="tableFilter.searchTerm"
          :debounce="300"
          placeholder="Search roles or scopes"
          @update:model-value="handleSearchChange"
        />
        <q-separator />
        <runai-table
          :columns="modalOptions.columns"
          :rows="displayAccessRules"
          :loading="displayLoading"
          :filter-by="tableFilter"
          @update-filters="updateFilter"
          modal-view
          disable-selection
          :is-server-side-pagination="isServerSidePagination"
          :rows-per-page-options="[1, 5, 7, 10, 15, 20, 25]"
        >
          <template #no-data>
            <runai-table-no-data
              :filter-by="tableFilter"
              entity-name="access rule"
              :show-create-btn="false"
              :show-filter-icon-and-button="false"
              class="access-rule-no-data"
            />
          </template>
        </runai-table>
      </template>

      <template #footer>
        <q-btn label="Close" color="primary" @click="$emit('close')" />
      </template>
    </runai-base-dialog>
  </section>
</template>

<script lang="ts">
import { defineComponent, type PropType } from "vue";

// cmps
import { RunaiBaseDialog } from "@/components/common/runai-base-dialog";
import { RunaiListModalSearch } from "@/components/common";
import { RunaiTable } from "@/components/common/runai-table";
import { RunaiTableNoData } from "@/components/common/runai-table/runai-table-no-data";

// models
import type { IFilterBy, IPaginationFilter } from "@/models/filter.model";
import type { IAccessRuleTableModalOptions } from "@/models/access-rule.model";
import { allAccessRulesColumnsMap } from "@/table-models/access-rule.table-model";
import {
  AccessRulesSortFilterFields,
  SubjectTypeExtended,
  type AccessRule,
} from "@/swagger-models/authorization-client";

// services
import { filterService } from "@/services/filter.service/filter.service";
import { accessRuleService } from "@/services/control-plane/rbac/access-rule.service/access-rule.service";

// utils
import { filterUtil } from "@/utils/filter.util/filter.util";
import { ErrorAlert } from "@/utils/error-alert.util";

// all-authenticated-users rules apply to every authenticated identity
const ALL_AUTH_RELEVANT_SUBJECT_TYPES: SubjectTypeExtended[] = [
  SubjectTypeExtended.User,
  SubjectTypeExtended.ServiceAccount,
];

export default defineComponent({
  components: {
    RunaiTableNoData,
    RunaiTable,
    RunaiBaseDialog,
    RunaiListModalSearch,
  },
  emits: ["close"],
  props: {
    modalOptions: {
      type: Object as PropType<IAccessRuleTableModalOptions>,
      required: true,
    },
  },
  data() {
    return {
      tableFilter: {} as IFilterBy,
      loading: false as boolean,
      accessRules: [] as AccessRule[],
    };
  },
  created() {
    this.loadFilters();
    if (this.includeAllAuth || this.includeGroups) {
      void this.loadSubjectAndRelatedRules();
    }
  },
  computed: {
    includeAllAuth(): boolean {
      // all-authenticated-users rules apply to every authenticated identity, so surface them
      // alongside a user/service-account subject's own rules (merged + paginated client-side).
      return (
        !this.modalOptions.accessRules &&
        !this.modalOptions.scopeId &&
        ALL_AUTH_RELEVANT_SUBJECT_TYPES.includes(this.modalOptions.subjectType as SubjectTypeExtended)
      );
    },
    includeGroups(): boolean {
      // a user subject inherits access rules from their (e.g. SSO) groups, so surface those
      // alongside the user's own direct rules (merged + paginated client-side).
      return (
        !this.modalOptions.accessRules &&
        !this.modalOptions.scopeId &&
        this.modalOptions.subjectType === SubjectTypeExtended.User &&
        !!this.modalOptions.subjectGroups?.length
      );
    },
    isServerSidePagination(): boolean {
      return !this.modalOptions.accessRules && !this.includeAllAuth && !this.includeGroups;
    },
    displayAccessRules(): AccessRule[] {
      return this.modalOptions.accessRules || this.accessRules;
    },
    displayLoading(): boolean {
      return this.loading || !!this.modalOptions.loading;
    },
  },
  methods: {
    loadFilters(): void {
      const defaultFilters: IFilterBy = filterService.getDefaultFilters(
        allAccessRulesColumnsMap.role.name,
        this.modalOptions.columns,
        true,
      );

      if (this.isServerSidePagination) {
        const filtersWithPagination = filterService.getEmptyFilterByModel({
          ...defaultFilters,
          rowsPerPage: 10,
        });
        this.updateFilter(filtersWithPagination);
      } else {
        this.tableFilter = filterService.getEmptyFilterByModel(defaultFilters);
      }
    },

    handleSearchChange(searchTerm: string): void {
      this.updateFilter({ ...this.tableFilter, searchTerm: searchTerm || "", page: 1 });
    },

    async updateFilter(filter: IFilterBy): Promise<void> {
      this.tableFilter = filter;

      if (this.isServerSidePagination) {
        try {
          this.loading = true;
          await Promise.all([this.loadAccessRulesCount(), this.loadAccessRules()]);
        } catch (error: unknown) {
          this.handleLoadAccessRulesError(error);
        } finally {
          this.loading = false;
        }
      }
    },

    async loadAccessRules(): Promise<void> {
      if (!this.isServerSidePagination) return;

      try {
        this.loading = true;
        this.accessRules = await this.getAccessRules();
      } catch (e) {
        this.handleLoadAccessRulesError(e);
      } finally {
        this.loading = false;
      }
    },

    mapFiltersToParams(): IPaginationFilter {
      return filterService.mapColumnsFilterToFilterParams(
        this.tableFilter,
        Object.values(AccessRulesSortFilterFields),
        true,
      );
    },

    getFilters(): IPaginationFilter {
      const paginationFilters = this.mapFiltersToParams();
      let filterBy: string[] = [];

      //Scope filters
      if (this.modalOptions.scopeType && this.modalOptions.scopeId) {
        filterBy = [
          filterUtil.getEqualsFilterString(AccessRulesSortFilterFields.ScopeType, this.modalOptions.scopeType),
          filterUtil.getEqualsFilterString(AccessRulesSortFilterFields.ScopeId, this.modalOptions.scopeId),
        ];
      }

      // Subject filters
      else if (this.modalOptions.subjectType && this.modalOptions.subjectId) {
        filterBy = [
          filterUtil.getEqualsFilterString(AccessRulesSortFilterFields.SubjectType, this.modalOptions.subjectType),
          filterUtil.getEqualsFilterString(AccessRulesSortFilterFields.SubjectId, this.modalOptions.subjectId),
        ];
      }

      return {
        ...paginationFilters,
        filterBy: [...(paginationFilters.filterBy || []), ...filterBy],
      };
    },

    async loadSubjectAndRelatedRules(): Promise<void> {
      try {
        this.loading = true;
        const subjectFilter = [
          filterUtil.getEqualsFilterString(
            AccessRulesSortFilterFields.SubjectType,
            this.modalOptions.subjectType as SubjectTypeExtended,
          ),
          filterUtil.getEqualsFilterString(AccessRulesSortFilterFields.SubjectId, this.modalOptions.subjectId as string),
        ];
        const subjectRules = await accessRuleService.getAllAccessRules(subjectFilter);

        // All-auth and per-group rules are enrichment on top of the subject's own rules, so each
        // fetch is isolated: one failing group (or the all-auth fetch) never blanks out the rest.
        const enrichmentFetches: Promise<AccessRule[]>[] = [];

        if (this.includeAllAuth) {
          const allAuthFilter = [
            filterUtil.getEqualsFilterString(
              AccessRulesSortFilterFields.SubjectType,
              SubjectTypeExtended.AllAuthenticatedUsers,
            ),
          ];
          enrichmentFetches.push(this.fetchRulesIsolated(allAuthFilter));
        }

        if (this.includeGroups) {
          for (const groupName of this.modalOptions.subjectGroups as string[]) {
            const groupFilter = [
              filterUtil.getEqualsFilterString(AccessRulesSortFilterFields.SubjectType, SubjectTypeExtended.Group),
              filterUtil.getEqualsFilterString(AccessRulesSortFilterFields.SubjectId, groupName),
            ];
            enrichmentFetches.push(this.fetchRulesIsolated(groupFilter));
          }
        }

        const enrichmentRules = await Promise.all(enrichmentFetches);
        this.accessRules = [subjectRules, ...enrichmentRules].flat();
      } catch (error: unknown) {
        this.handleLoadAccessRulesError(error);
      } finally {
        this.loading = false;
      }
    },

    async fetchRulesIsolated(filterBy: string[]): Promise<AccessRule[]> {
      try {
        return await accessRuleService.getAllAccessRules(filterBy);
      } catch (error: unknown) {
        this.handleLoadAccessRulesError(error);
        return [];
      }
    },

    async getAccessRules(): Promise<AccessRule[]> {
      const filters = this.getFilters();
      return await accessRuleService.getAccessRulesV2(
        filters.filterBy,
        filters.sortBy as AccessRulesSortFilterFields | undefined,
        filters.sortOrder,
        filters.offset,
        filters.limit,
        filters.search,
      );
    },

    async loadAccessRulesCount(): Promise<void> {
      const filters = this.getFilters();
      this.tableFilter.rowsNumber = await accessRuleService.getAccessRulesCount(filters.filterBy, filters.search);
    },

    handleLoadAccessRulesError(error: unknown): void {
      const errorAlert = new ErrorAlert({
        generalMessage: ErrorAlert.failedLoadingMessage("access rules"),
      });
      this.$q.notify(errorAlert.getNotification(error));
    },
  },
});
</script>

<style scoped lang="scss">
:deep(.access-rule-no-data .empty-state-container) {
  min-height: 124px;
  align-items: center;
}
</style>
