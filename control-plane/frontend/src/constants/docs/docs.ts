/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

import { getBootConfig } from "@/services/infra/boot-config.service/boot-config.service";

const RUNAI_DOCS_ORIGIN = "https://run-ai-docs.nvidia.com";
const LATEST_URL = "https://docs.run.ai/latest";
// when you change this link make sure you update the login-success.html as well
const bootConfig = getBootConfig();

export function getDocsUrl(path: string): string {
  const docsBaseUrl = getBootConfig().docs.baseUrl.replace(/\/+$/, "");
  const normalizedPath = path.replace(/^\/+/, "");
  return normalizedPath ? `${docsBaseUrl}/${normalizedPath}` : docsBaseUrl;
}

export const DOCS = {
  ROOT: bootConfig.docs.baseUrl,

  INSTALLATION: {
    CLUSTER: getDocsUrl("getting-started/overview#nvidia-run-ai-cluster"),
    PREREQUISITES: getDocsUrl("getting-started/installation/install-using-helm/system-requirements"),
    NETWORK_REQUIREMENTS: getDocsUrl("getting-started/installation/install-using-helm/network-requirements"),
    HELM_INSTALL: getDocsUrl("getting-started/installation/install-using-helm/helm-install"),
    CUSTOM_INSTALL: getDocsUrl("getting-started/installation/custom-install"),
    UNINSTALL: getDocsUrl("getting-started/installation/install-using-helm/uninstall"),
    FQDN_REQUIREMENTS: getDocsUrl(
      "getting-started/installation/install-using-helm/system-requirements#fully-qualified-domain-name-fqdn",
    ),
    DISTRIBUTED_TRAINING: getDocsUrl(
      "getting-started/installation/install-using-helm/system-requirements#distributed-training",
    ),
    EMAIL_NOTIFICATIONS: getDocsUrl("settings/general-settings/notifications#email-notifications"),
    USER_INVITATIONS: getDocsUrl("infrastructure-setup/authentication/users#creating-a-local-user"),
  },

  AUTHENTICATION: {
    OVERVIEW: getDocsUrl("infrastructure-setup/authentication/overview"),
    ACCESS_RULES: getDocsUrl("infrastructure-setup/authentication/accessrules"),
    SSO: getDocsUrl("infrastructure-setup/authentication/sso/"),
    SAML: getDocsUrl("infrastructure-setup/authentication/sso/saml"),
    OIDC: getDocsUrl("infrastructure-setup/authentication/sso/openidconnect"),
    OPENSHIFT: getDocsUrl("infrastructure-setup/authentication/sso/openshift"),
    ROLES: getDocsUrl("infrastructure-setup/authentication/roles#roles-in-nvidia-run-ai"),
    CUSTOM_ROLES_API: getDocsUrl("infrastructure-setup/authentication/roles#custom-roles-api-only"),
    RESEARCHER: {
      OVERVIEW: `${LATEST_URL}/admin/runai-setup/authentication/researcher-authentication`,
      CLI_ACCESS: `${LATEST_URL}/admin/runai-setup/authentication/researcher-authentication/#command-line-interface-access`,
      K8S_CONFIG: getDocsUrl("infrastructure-setup/authentication/cluster-authentication"),
    },
  },

  CLI: {
    INSTALL: getDocsUrl("reference/cli/install-cli"),
    SUBMIT: getDocsUrl("reference/cli/runai"),
  },

  PLATFORM_MANAGEMENT: {
    PROJECTS: getDocsUrl("platform-management/aiinitiatives/organization/projects"),
    DEPARTMENTS: getDocsUrl("platform-management/aiinitiatives/organization/departments"),
    POLICIES: getDocsUrl("platform-management/policies/policies-and-rules"),
    SCHEDULING_RULES: getDocsUrl("platform-management/policies/scheduling-rules"),
    SCHEDULING_CONCEPTS: getDocsUrl("platform-management/runai-scheduler/scheduling/concepts-and-principles"),
    DYNAMIC_FRACTION: getDocsUrl("platform-management/runai-scheduler/resource-optimization/dynamic-fractions"),
    ORGANIZATION: {
      PROJECTS: getDocsUrl("platform-management/aiinitiatives/organization/projects"),
      DEPARTMENTS: getDocsUrl("platform-management/aiinitiatives/organization/departments"),
    },
    NODE_GROUPS: `${LATEST_URL}/admin/researcher-setup/limit-to-node-group`,
    REPORTS: getDocsUrl("platform-management/monitor-performance/reports"),
  },

  WORKLOADS: {
    DATA_VOLUMES: getDocsUrl("workloads-in-nvidia-run-ai/assets/data-volumes"),
    EXTERNAL_ACCESS_TO_CONTAINERS: getDocsUrl(
      "infrastructure-setup/advanced-setup/container-access/external-access-to-containers",
    ),
    CREATING_A_HUGGING_FACE_INFERENCE_WORKLOAD: getDocsUrl(
      "workloads-in-nvidia-run-ai/using-inference/hugging-face-inference#creating-a-hugging-face-inference-workload",
    ),
  },

  SLACK_NOTIFICATIONS: {
    CONNECT_WORKSPACE: getDocsUrl("settings/general-settings/notifications#slack-notifications"),
  },

  TROUBLESHOOTING: {
    CLUSTERS: getDocsUrl("infrastructure-setup/procedures/clusters"),
    TERMINAL: getDocsUrl("workloads-in-nvidia-run-ai/workloads#troubleshooting"),
  },

  CONFIGURATION: {
    NON_ROOT_CONTAINERS: getDocsUrl(
      "latest/infrastructure-setup/advanced-setup/container-access/user-identity-in-containers",
    ),
  },

  TOPOLOGIES: {
    SCHEDULING: getDocsUrl("platform-management/aiinitiatives/resources/topology-aware-scheduling"),
    GB200: getDocsUrl("platform-management/aiinitiatives/resources/using-gb200"),
  },

  NODE_POOLS: {
    TIME_BASED_FAIR_SHARE: getDocsUrl(
      "platform-management/runai-scheduler/scheduling/how-the-scheduler-works#time-based-fairshare",
    ),
    PERSISTENT_DATA: getDocsUrl(
      "infrastructure-setup/advanced-setup/cluster-config#all-configurations-helm-and-runaiconfig",
    ),
    NUMA_AWARE_SCHEDULING: getDocsUrl("platform-management/aiinitiatives/resources/numa-aware-scheduling"),
  },

  WHATS_NEW: {
    URL: bootConfig.docs.whatsNew.url,
    LABEL: bootConfig.docs.whatsNew.label,
  },

  EXTERNAL: {
    VLLM: {
      SUPPORTED_MODELS: "https://docs.vllm.ai/en/latest/models/supported_models.html",
    },
    HUGGING_FACE: {
      BASE_URL: "https://huggingface.co",
      TGI_SUPPORTED_MODELS: "https://huggingface.co/docs/text-generation-inference/en/supported_models",
    },
    NIM: {
      CATALOG_BASE_URL: "https://catalog.ngc.nvidia.com/orgs/nim/teams",
      NGC_CATALOG: "https://docs.nvidia.com/ngc/gpu-cloud/ngc-catalog-user-guide/index.html#what-is-ngc-catalog",
      DEPLOY_INFERENCE: getDocsUrl("workloads-in-nvidia-run-ai/using-inference/nim-inference"),
    },
  },

  // when you change these links make sure you update the login-success.html as well
  SUPPORT: {
    URL: bootConfig.support.url ?? "",
    LABEL: bootConfig.support.label,
  },

  KUBERNETES: {
    TASKS: {
      EXTENDED_RESOURCES: "https://kubernetes.io/docs/tasks/configure-pod-container/extended-resource/",
    },
    REFERENCE: {
      API: {
        QUANTITY: "https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/",
      },
    },
  },
} as const;

export function resolveDocsLinksInHtml(html: string): string {
  return html.replace(/(\bhref\s*=\s*)(["'])([^"']+)\2/gi, (match, prefix, quote, href: string) => {
    const parsedUrl = parseRunaiDocsUrl(href);
    if (!parsedUrl) return match;

    const pathParts = parsedUrl.pathname.split("/").filter(Boolean);
    pathParts.shift();
    if (isDocsVersionSegment(pathParts[0])) pathParts.shift();

    const resolvedUrl = `${getDocsUrl(pathParts.join("/"))}${parsedUrl.search}${parsedUrl.hash}`;
    return `${prefix}${quote}${resolvedUrl}${quote}`;
  });
}

function parseRunaiDocsUrl(url: string): URL | null {
  try {
    const parsedUrl = new URL(url);
    return parsedUrl.origin === RUNAI_DOCS_ORIGIN ? parsedUrl : null;
  } catch {
    return null;
  }
}

function isDocsVersionSegment(value: string | undefined): boolean {
  if (!value) return false;
  const normalized = value.startsWith("v") ? value.slice(1) : value;
  const parts = normalized.split(".");
  return (
    parts.length === 2 && parts.every((part) => part.length > 0 && [...part].every((char) => char >= "0" && char <= "9"))
  );
}

export type DocsType = typeof DOCS;

export default DOCS;
