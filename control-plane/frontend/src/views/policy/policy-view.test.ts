/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

import { beforeEach, describe, expect, it, vi } from "vitest";
import { flushPromises, shallowMount } from "@vue/test-utils";
import { createTestingPinia } from "@pinia/testing";
import PolicyView from "./policy-view.vue";
import { gvkPolicyService } from "@/services/control-plane/gvk-policy.service/gvk-policy.service";
import { ScopeType, type GetGvkPolicyResp } from "@/swagger-models/policy-service-client";

vi.mock("@/services/control-plane/gvk-policy.service/gvk-policy.service");

const workloadType = {
  group: "apps",
  version: "v1",
  kind: "Deployment",
};

function mountComponent() {
  return shallowMount(PolicyView, {
    global: {
      plugins: [createTestingPinia({ createSpy: vi.fn })],
      mocks: {
        $route: {
          params: { id: "550e8400-e29b-41d4-a716-446655440000" },
          query: {},
          meta: {},
        },
        $router: { push: vi.fn() },
      },
    },
  });
}

describe("policy-view.vue", () => {
  beforeEach(() => {
    vi.clearAllMocks();
    vi.mocked(gvkPolicyService.getEffectivePolicy).mockResolvedValue([]);
  });

  it.each([
    [ScopeType.Cluster, { clusterId: "cluster-uuid" }, { clusterId: "cluster-uuid" }],
    [
      ScopeType.Department,
      { clusterId: "cluster-uuid", departmentId: "department-id" },
      { departmentId: "department-id" },
    ],
    [
      ScopeType.Project,
      { clusterId: "cluster-uuid", departmentId: "department-id", projectId: "project-id" },
      { projectId: "project-id" },
    ],
  ])(
    "loads effective GVK rules with only the identifier for %s scope",
    async (scope, policyScopeIds, requestScopeIds) => {
      const policy: GetGvkPolicyResp = {
        id: "550e8400-e29b-41d4-a716-446655440000",
        workloadType,
        scope,
        createdBy: "test@run.ai",
        createdAt: "2026-09-03T00:00:00Z",
        updatedBy: "test@run.ai",
        updatedAt: "2026-09-03T00:00:00Z",
        ...policyScopeIds,
      };
      vi.mocked(gvkPolicyService.getGvkPolicy).mockResolvedValue(policy);

      mountComponent();
      await flushPromises();

      expect(gvkPolicyService.getEffectivePolicy).toHaveBeenCalledWith({
        workloadType,
        scope,
        ...requestScopeIds,
      });
    },
  );
});
