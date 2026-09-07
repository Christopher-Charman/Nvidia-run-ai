package checker

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/run-ai/runai-common-packages/cache/v2"
	"github.com/run-ai/runai-common-packages/cache/v2/options"
	"github.com/run-ai/runai-common-packages/logging"
	"github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client/cluster-capabilities/clients/cluster"

	goVersion "github.com/hashicorp/go-version"
)

// ClusterCapabilitiesChecker this class is responsible to check if a given cluster is supporting capability X.
// Ideally, we prefer the cluster to report this,
// but since framework for this is not yet available we keep an internal map here which states
// for each cluster capability - its minimal supported version.
type ClusterCapabilitiesChecker struct {
	clusterId   string
	clusterInfo *clusterInfo
}

const (
	// MinVersionGitSyncOneParameter old cluster support both branch and revision,
	// but new cluster no longer support this and in this case the request for old clusters will be rejected.
	MinVersionGitSyncOneParameter = iota

	// MinVersionInferenceMetricsThroughputLatency cluster support for inference metrics throughput and latency
	MinVersionInferenceMetricsThroughputLatency

	// MinVersionOfGpuDeviceRequest - older clusters require whole gpus to be sent as gpu, not as gpuDevices
	MinVersionOfGpuDeviceRequest

	// MinVersionOfConfigMapAsVolume is the min cluster version that support ConfigMap as volume
	MinVersionOfConfigMapAsVolume

	// MinVersionOfSecretAsVolume is the min cluster version that support Secret as volume
	MinVersionOfSecretAsVolume

	// MinVersionOfInferenceWorkloadNotification is the min cluster version required for notification about workload creation
	// from asset service to workload service for inference workloads
	MinVersionOfInferenceWorkloadNotification

	// MinVersionOfTrainingWorkloadNotification is the min cluster version required for notification about workload creation
	// from asset service to workload service for training workloads (including also distributed training)
	MinVersionOfTrainingWorkloadNotification

	// MinVersionOfDatavolume is the min cluster version for datavolumes
	MinVersionOfDatavolume

	// MinVersionOfAuthenticatedGroup is the min cluster version where authenticated group is supported
	MinVersionOfAuthenticatedGroup

	// MinVersionOfProjectQuotaStatuses is the min cluster version where project quota statuses exist
	MinVersionOfProjectQuotaStatuses

	// MinVersionOfClusterResourceStatus - older clusters do not report to k8s-objects-tracker about cluster resource status and issues
	MinVersionOfClusterReportingResourceStatus

	//MinVersionOfPVCSyncStatus - older clusters do not report to asset-service about PVCs sync status
	MinVersionOfPVCSyncStatus

	// MinVersionOfServingPort - the min cluster version that supports serving port as service type
	MinVersionOfServingPort

	// MinVersionOfServingPortAuthorizationFields - the min cluster version that supports serving port authorization fields
	MinVersionOfServingPortAuthorizationFields

	// MinVersionOfWorkloadV2 - to support workload v2, the agent needs to support fully SAAS
	MinVersionOfWorkloadV2

	// MinVersionOfNfsHostPathForInference - the min cluster version that supports NFS and hostPath volumes for inference workloads
	MinVersionOfNfsHostPathForInference

	// MinVersionOfNfsSectionInWorkloadPayload - the min cluster version that supports NFS section in workload payload
	MinVersionOfNfsSectionInWorkloadPayload

	// MinVersionOfGpuResourceOptimization - the min cluster version supporting memory limit and other optimization fields
	MinVersionOfGpuResourceOptimization

	// MinVersionOfClusterConfigMapWideScope is the cluster minimal version, which supports Config Map cluster wide scope
	MinVersionOfClusterConfigMapWideScope

	MinVersionOfNewWorkloadOwnershipProtection

	// MinVersionOfDeprecatedSpecName is the cluster minimal version, which workload name in spec is deprecated.
	MinVersionOfDeprecatedSpecName

	// MinVersionOfWorkloadVolumesForInference - the min cluster version that supports workload volumes for inference workloads
	MinVersionOfWorkloadVolumesForInference

	// MinVersionOfAdvancedAutoscalingSettingsForInference - the min cluster version that supports advanced autoscaling fields for inference workloads
	MinVersionOfAdvancedAutoscalingSettingsForInference

	// MinVersionOfTerminationGracePeriodSeconds - the mini cluster version supporting this field
	MinVersionOfTerminationGracePeriodSeconds

	// MinVersionOfMultiFractionalGpuDevices - the min cluster version that supports multi-fractional gpu devices
	MinVersionOfMultiFractionalGpuDevices

	// MinVersionOfDistributedCleanPodPolicy - the min cluster version that supports clean pod policy for distributed workloads
	MinVersionOfDistributedCleanPodPolicy

	// MinVersionOfConfigMapEnvValue - the min cluster version supporting the CONFIGMAP:xxx,yyy value format
	MinVersionOfConfigMapEnvValue

	// MinVersionOfPodFieldRefEnvValue - the min cluster version supporting the PODFIELDREF:xxx,yyy value format
	MinVersionOfPodFieldRefEnvValue

	// MinVersionOfAssetsAgentSync - the min version of cluster supporting agent sync for assets
	MinVersionOfAssetsAgentSync

	// MinVersionUpdateInferenceSpec - the min cluster version that supports update inference spec api
	MinVersionUpdateInferenceSpec

	// MinVersionOfOnlyProjectScopeExistingPvcSupported - the min cluster version that supports only project-scoped existing pvcs
	MinVersionOfOnlyProjectScopeExistingPvcSupported

	// MinVersionOfConfigMapSubPathSupported - the min cluster version that supports sub paths of config maps
	MinVersionOfConfigMapSubPathSupported

	// MinVersionOfDistributedJaxJobsSupported - the min cluster version that supports JAX distributed workload
	MinVersionOfDistributedJaxJobsSupported

	// MinVersionTrainingPriorityClassPatchFix - the min cluster version that supports training priority class patch fix
	MinVersionTrainingPriorityClassPatchFix

	// MinVersionRevisions - the min cluster version that supports revisions sync
	MinVersionRevisions

	// MinVersionOfClusterAndDepartmentScopeDataVolumes - the min cluster version that supports cluster/department scope data volumes
	MinVersionOfClusterAndDepartmentScopeDataVolumes

	// MinVersionAccessRulesStatuses - the min cluster version that supports access rules statuses
	MinVersionAccessRulesStatuses

	MinVersionOfPVCDataSharing

	// MinVersionOfLauncherCreationPolicy - the min cluster version that supports LauncherCreationPolicy
	MinVersionOfLauncherCreationPolicy

	// MinVersionOfNodeAffinity - the min cluster version that supports Node Affinity
	MinVersionOfNodeAffinity

	// MinVersionOfInferenceServingConfiguration - the min cluster version that supports inference serving configuration
	MinVersionOfInferenceServingConfiguration

	// MinVersionOfPvcAddedAttributes - PVC added attributes support
	MinVersionOfPvcAddedAttributes

	// MinVersionOfUserCredentials - the min cluster version that supports user credentials
	MinVersionOfUserCredentials

	// MinVersionOfImagePullSecrets - the min cluster version that supports image pull secrets
	MinVersionOfImagePullSecrets

	// MinVersionOfFlexiblePriorityClasses - Flexible priority classes from control plane
	MinVersionOfFlexiblePriorityClasses

	// NOTE -> Make sure to add the min version to the map below
)

var clusterVersionToCapabilityMap = map[int]string{
	MinVersionGitSyncOneParameter:                       "2.15.0",
	MinVersionOfGpuDeviceRequest:                        "2.16.0",
	MinVersionInferenceMetricsThroughputLatency:         "2.18.0",
	MinVersionOfConfigMapAsVolume:                       "2.17.0",
	MinVersionOfInferenceWorkloadNotification:           "2.16.0",
	MinVersionOfTrainingWorkloadNotification:            "2.13.0",
	MinVersionOfSecretAsVolume:                          "2.18.0",
	MinVersionOfDatavolume:                              "2.18.0",
	MinVersionOfAuthenticatedGroup:                      "2.18.0",
	MinVersionOfProjectQuotaStatuses:                    "2.15.0",
	MinVersionOfClusterReportingResourceStatus:          "2.17.0",
	MinVersionOfPVCSyncStatus:                           "2.17.0",
	MinVersionOfServingPort:                             "2.17.0",
	MinVersionOfServingPortAuthorizationFields:          "2.19.0",
	MinVersionOfWorkloadV2:                              "2.18.0",
	MinVersionOfNfsHostPathForInference:                 "2.19.0",
	MinVersionOfNfsSectionInWorkloadPayload:             "2.16.0",
	MinVersionOfGpuResourceOptimization:                 "2.17.0",
	MinVersionOfClusterConfigMapWideScope:               "2.19.0",
	MinVersionOfNewWorkloadOwnershipProtection:          "2.20.0",
	MinVersionOfDeprecatedSpecName:                      "2.20.0",
	MinVersionOfWorkloadVolumesForInference:             "2.20.0",
	MinVersionOfAdvancedAutoscalingSettingsForInference: "2.20.0",
	MinVersionOfTerminationGracePeriodSeconds:           "2.20.0",
	MinVersionOfMultiFractionalGpuDevices:               "2.20.0",
	MinVersionOfDistributedCleanPodPolicy:               "2.20.0",
	MinVersionOfConfigMapEnvValue:                       "2.21.0",
	MinVersionOfPodFieldRefEnvValue:                     "2.19.0",
	MinVersionOfAssetsAgentSync:                         "2.21.0",
	MinVersionUpdateInferenceSpec:                       "2.19.0",
	MinVersionOfOnlyProjectScopeExistingPvcSupported:    "2.21.0",
	MinVersionOfConfigMapSubPathSupported:               "2.21.0",
	MinVersionOfDistributedJaxJobsSupported:             "2.21.0",
	MinVersionTrainingPriorityClassPatchFix:             "2.18.0",
	MinVersionRevisions:                                 "2.21.0",
	MinVersionOfClusterAndDepartmentScopeDataVolumes:    "2.22.0",
	MinVersionAccessRulesStatuses:                       "2.22.0",
	MinVersionOfPVCDataSharing:                          "2.22.0",
	MinVersionOfLauncherCreationPolicy:                  "2.22.0",
	MinVersionOfNodeAffinity:                            "2.22.0",
	MinVersionOfInferenceServingConfiguration:           "2.22.0",
	MinVersionOfPvcAddedAttributes:                      "2.22.0",
	MinVersionOfUserCredentials:                         "2.23.0",
	MinVersionOfImagePullSecrets:                        "2.23.0",
	MinVersionOfFlexiblePriorityClasses:                 "2.22.0",
}

// ClustersCacheConfig is a struct that holds the configuration for the clusters cache.
// Holds the expiration time for the cache and the cleanup frequency.
type ClustersCacheConfig struct {
	ExpirationSec int
	CleanupSec    int
}

func MinVersionForSupporting(capability int) string {
	return clusterVersionToCapabilityMap[capability]
}

// InitClusterCapabilitiesChecker initializes the cluster capabilities checker with the cluster client connection.
// and the clusters cache configuration.
// note that if you initialize the checker with specific version, this initialization is not really needed.
func InitClusterCapabilitiesChecker(clusterUrl string, clustersCache ClustersCacheConfig, logLevel string) {
	cluster.InitConnection(clusterUrl, logLevel)
	InitClustersCache(clustersCache)
}

// InitClustersCache initializes the clusters cache with the given configuration.
func InitClustersCache(clustersCacheConfig ClustersCacheConfig) {
	cache.ClearStore[clusterInfo]()
	err := cache.RegisterStore[clusterInfo](options.Ttl(time.Duration(clustersCacheConfig.ExpirationSec) * time.Second))
	if err != nil {
		logging.Warnf("Failed to register cache store for type clusterInfo: %v", err)
	}
}

// NewClusterCapabilitiesChecker creates a new instance of ClusterCapabilitiesChecker
// with clusterId. cluster version will be loaded from the cluster service when needed.
// This is one option to initialize the checker, another option is to initialize it with
// cluster version (see below)
func NewClusterCapabilitiesChecker(clusterId string) *ClusterCapabilitiesChecker {
	return &ClusterCapabilitiesChecker{clusterId: clusterId}
}

func NewClusterCapabilitiesCheckerForVersion(clusterVersion string) (*ClusterCapabilitiesChecker, error) {
	clusterSemverVersion, err := goVersion.NewSemver(normalizeVersion(clusterVersion))
	if err != nil {
		return nil, fmt.Errorf("failed to parse cluster version: %s, err: %w", clusterVersion, err)
	}
	return &ClusterCapabilitiesChecker{clusterInfo: &clusterInfo{Version: clusterSemverVersion}}, nil
}

// ValidateClusterCapability checks if the cluster supports the given capability.
func (c *ClusterCapabilitiesChecker) ValidateClusterCapability(ctx context.Context, capabilityToCheck int) (bool, error) {

	clusterInfo, err := c.getClusterInfo(ctx)
	if err != nil {
		return false, err
	}

	capabilityMinVersion, ok := clusterVersionToCapabilityMap[capabilityToCheck]
	if !ok {
		return false, errors.New(fmt.Sprintf("capabilty %d is not supported by the checker", capabilityToCheck))
	}

	capabilityRequiredVersion, err := goVersion.NewVersion(capabilityMinVersion)
	if err != nil {
		return false, errors.New(fmt.Sprintf("failed to parse capability version: %s, err: %v", capabilityMinVersion, err))
	}

	if clusterInfo.Version.LessThan(capabilityRequiredVersion) {
		return false, nil
	}

	return true, nil
}

// ClusterConfiguration returns the cluster configuration
func (c *ClusterCapabilitiesChecker) ClusterConfiguration(ctx context.Context) (*ClusterConfiguration, error) {
	clusterInfo, err := c.getClusterInfo(ctx)
	if err != nil {
		return nil, err
	}
	// If ClusterCapabilitiesChecker was init without fetching the cluster configuration(NewClusterCapabilitiesCheckerForVersion),
	// force fetch cluster info
	if clusterInfo.Configuration == nil {
		clusterInfo, err = c.fetchClusterInfo(ctx)
	}
	return clusterInfo.Configuration, nil
}

func (c *ClusterCapabilitiesChecker) getClusterInfo(ctx context.Context) (*clusterInfo, error) {
	//   if cluster version is not known to the capabilities checker, it will need to get it from the
	//   cluster service
	if c.clusterInfo != nil {
		return c.clusterInfo, nil
	}
	return c.fetchClusterInfo(ctx)
}

func (c *ClusterCapabilitiesChecker) fetchClusterInfo(ctx context.Context) (*clusterInfo, error) {
	cachedInfo, clusterFound := getCacheClusterInfo(c.clusterId)
	if clusterFound {
		return cachedInfo, nil
	}

	clusterDetails, apiErr := cluster.Get().Client().GetClusterByUuid(ctx, c.clusterId)
	if apiErr != nil {
		return nil, errors.New(fmt.Sprintf("failed to check cluster capabilities, err: %v", apiErr))
	}
	clusterInfo, err := setCacheClusterInfo(c.clusterId, clusterDetails)
	if err != nil {
		return nil, err
	}

	return clusterInfo, nil
}

func normalizeVersion(version string) string {
	if strings.HasPrefix(version, "0.0.0") || strings.Contains(version, "master") {
		// That indicates that this is a Dev version, therefore we consider it as Max value
		return "9.9.9"
	}

	index := strings.Index(version, "-")
	if index != -1 {
		// Extract substring before the dash
		return version[:index]
	}

	return version
}
