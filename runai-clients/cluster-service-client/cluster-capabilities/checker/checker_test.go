package checker

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/run-ai/runai-application-packages/convutils"
	"github.com/run-ai/runai-application-packages/testutils"
	"github.com/run-ai/runai-common-packages/cache/v2"
	cache_testing "github.com/run-ai/runai-common-packages/cache/v2/testing"
	clusterClient "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-capabilities/clients/cluster"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client/generated"
	workloadsService "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client/generated"
)

var defaultCheckerClusterConfig = ClustersCacheConfig{ExpirationSec: 600, CleanupSec: 10}

var _ = Describe("Cluster capabilities checker", func() {
	clusterId := "cluster-id"
	differentClusterId := "different-cluster-id"
	wrongVersionClusterId := "wrong-version-cluster-id"
	var clusterClientMock clusterClient.SimpleMockClient

	BeforeEach(func() {
		// Before each test
		clusterClientMock = clusterClient.NewSimpleMockClient().
			WithCluster(clusterId, buildDisplayedCluster(testutils.NewRandomValues(), "2.13.0", nil)).
			WithCluster(differentClusterId, buildDisplayedCluster(testutils.NewRandomValues(), "2.17.0", nil)).
			WithCluster(wrongVersionClusterId, buildDisplayedCluster(testutils.NewRandomValues(), "wrong-version", nil))
		clusterClient.SetMock(clusterClientMock)
		cache_testing.EnableCache()
		InitClustersCache(defaultCheckerClusterConfig)
	})

	Context("cluster capability checker functionality", func() {

		Context("Should not be capable", func() {
			It("should not be capable due to incompatible version", func() {
				checker := NewClusterCapabilitiesChecker(clusterId)
				isCapable, _ := checker.ValidateClusterCapability(context.Background(), MinVersionGitSyncOneParameter)
				Expect(isCapable).To(BeFalse())
			})

			It("should not be capable due to wrong version", func() {
				checker := NewClusterCapabilitiesChecker(wrongVersionClusterId)
				isCapable, _ := checker.ValidateClusterCapability(context.Background(), MinVersionGitSyncOneParameter)
				Expect(isCapable).To(BeFalse())
			})
		})

		It("should be capable for compatible version", func() {
			checker := NewClusterCapabilitiesChecker(differentClusterId)
			isCapable, _ := checker.ValidateClusterCapability(context.Background(), MinVersionGitSyncOneParameter)
			Expect(isCapable).To(BeTrue())
		})

		It("should not call the cluster client if the cluster is in the cache", func() {
			err := cache.ClearStore[clusterInfo]()
			Expect(err).To(BeNil())

			checker := NewClusterCapabilitiesChecker(clusterId)
			_, _ = checker.ValidateClusterCapability(context.Background(), MinVersionGitSyncOneParameter)
			_, _ = checker.ValidateClusterCapability(context.Background(), MinVersionGitSyncOneParameter)
			Expect(*clusterClientMock.GetClusterByIdCallCount).To(Equal(1))
		})
	})

	Context("cluster capability checker ", func() {
		It("workload ownership protection", func() {
			clusterClient.SetMock(clusterClient.NewSimpleMockClient().
				WithCluster("no value", buildDisplayedCluster(testutils.NewRandomValues(), "2.13.0", nil)).
				WithCluster("protection off", buildDisplayedCluster(testutils.NewRandomValues(), "2.17.0", convutils.Ref(false))).
				WithCluster("protection on", buildDisplayedCluster(testutils.NewRandomValues(), "2.19.0", convutils.Ref(true))))
			checker := NewClusterCapabilitiesChecker("no value")
			protection, err := checker.ClusterConfiguration(context.Background())
			Expect(err).To(BeNil())
			Expect(protection.WorkloadOwnershipProtection).To(BeFalse())

			checker = NewClusterCapabilitiesChecker("protection off")
			protection, err = checker.ClusterConfiguration(context.Background())
			Expect(err).To(BeNil())
			Expect(protection.WorkloadOwnershipProtection).To(BeFalse())

			checker = NewClusterCapabilitiesChecker("protection on")
			protection, err = checker.ClusterConfiguration(context.Background())
			Expect(err).To(BeNil())
			Expect(protection.WorkloadOwnershipProtection).To(BeTrue())
		})

		It("distributed frameworks status", func() {
			clusterClient.SetMock(clusterClient.NewSimpleMockClient().
				WithCluster("id", buildDisplayedCluster(testutils.NewRandomValues(), "2.20.0", nil)))

			checker := NewClusterCapabilitiesChecker("id")
			configuration, err := checker.ClusterConfiguration(context.Background())
			Expect(err).To(BeNil())

			frameworksStatus := configuration.DistributedFrameworksStatus
			Expect(len(DistributedFrameworksDependencies)).To(Equal(5))
			Expect(len(frameworksStatus)).To(Equal(len(DistributedFrameworksDependencies)))

			mpiStatus, ok := frameworksStatus[workloadsService.DISTRIBUTEDFRAMEWORK_MPI]
			Expect(ok).To(BeTrue())
			Expect(mpiStatus.Available).To(BeTrue())
			Expect(mpiStatus.Suspendable).To(BeTrue())

			pytorchStatus, ok := frameworksStatus[workloadsService.DISTRIBUTEDFRAMEWORK_PY_TORCH]
			Expect(ok).To(BeTrue())
			Expect(pytorchStatus.Available).To(BeTrue())
			Expect(pytorchStatus.Suspendable).To(BeFalse())

			tensorflowStatus, ok := frameworksStatus[workloadsService.DISTRIBUTEDFRAMEWORK_TF]
			Expect(ok).To(BeTrue())
			Expect(tensorflowStatus.Available).To(BeFalse())
			Expect(tensorflowStatus.Suspendable).To(BeFalse())

			// not in config, should be false by default
			xgboostStatus, ok := frameworksStatus[workloadsService.DISTRIBUTEDFRAMEWORK_XG_BOOST]
			Expect(ok).To(BeTrue())
			Expect(xgboostStatus.Available).To(BeFalse())
			Expect(xgboostStatus.Suspendable).To(BeFalse())

			// not in config, should be false by default
			jaxStatus, ok := frameworksStatus[workloadsService.DISTRIBUTEDFRAMEWORK_JAX]
			Expect(ok).To(BeTrue())
			Expect(jaxStatus.Available).To(BeFalse())
			Expect(jaxStatus.Suspendable).To(BeFalse())
		})
	})

	Context("knative availability and scale-to-zero capability", func() {
		It("should handle missing knative sub components", func() {
			clusterClient.SetMock(clusterClient.NewSimpleMockClient().
				WithCluster("no-sub-components", buildDisplayedClusterWithKnative(testutils.NewRandomValues(), "2.20.0", true, nil)))

			checker := NewClusterCapabilitiesChecker("no-sub-components")
			configuration, err := checker.ClusterConfiguration(context.Background())
			Expect(err).To(BeNil())
			Expect(configuration.KnativeAvailable).To(BeTrue())
			Expect(configuration.KnativeScaleToZeroAvailable).To(BeTrue())
		})

		It("should handle missing scale-to-zero component", func() {
			clusterClient.SetMock(clusterClient.NewSimpleMockClient().
				WithCluster("no-scale-zero", buildDisplayedClusterWithKnative(testutils.NewRandomValues(), "2.20.0", true, &map[string]openapi.ClusterDependencyStatus{})))

			checker := NewClusterCapabilitiesChecker("no-scale-zero")
			configuration, err := checker.ClusterConfiguration(context.Background())
			Expect(err).To(BeNil())
			Expect(configuration.KnativeAvailable).To(BeTrue())
			Expect(configuration.KnativeScaleToZeroAvailable).To(BeTrue())
		})

		It("should handle scale-to-zero enabled", func() {
			clusterClient.SetMock(clusterClient.NewSimpleMockClient().
				WithCluster("scale-zero-enabled", buildDisplayedClusterWithKnative(testutils.NewRandomValues(), "2.20.0", true, &map[string]openapi.ClusterDependencyStatus{
					ScaleToZeroDependency: {
						Available: true,
					},
				})))

			checker := NewClusterCapabilitiesChecker("scale-zero-enabled")
			configuration, err := checker.ClusterConfiguration(context.Background())
			Expect(err).To(BeNil())
			Expect(configuration.KnativeAvailable).To(BeTrue())
			Expect(configuration.KnativeScaleToZeroAvailable).To(BeTrue())
		})

		It("should handle scale-to-zero disabled", func() {
			clusterClient.SetMock(clusterClient.NewSimpleMockClient().
				WithCluster("scale-zero-disabled", buildDisplayedClusterWithKnative(testutils.NewRandomValues(), "2.20.0", true, &map[string]openapi.ClusterDependencyStatus{
					ScaleToZeroDependency: {
						Available: false,
					},
				})))

			checker := NewClusterCapabilitiesChecker("scale-zero-disabled")
			configuration, err := checker.ClusterConfiguration(context.Background())
			Expect(err).To(BeNil())
			Expect(configuration.KnativeAvailable).To(BeTrue())
			Expect(configuration.KnativeScaleToZeroAvailable).To(BeFalse())
		})
	})
})

func buildDisplayedCluster(randoms *testutils.RandomValues, version string, protection *bool) openapi.DisplayedCluster {
	result := openapi.DisplayedCluster{
		Uuid: randoms.ClusterId,
	}
	result.Version.Set(&version)

	displayedStatus := openapi.NewClusterDisplayedStatus()

	optionalDependencies := map[string]openapi.ClusterDependencyStatus{
		MPIDependency: {
			Available: true,
			Components: &map[string]openapi.ClusterDependencyStatus{
				DistributedSuspendComponent: {
					Available: true,
				},
			},
		},
		PyTorchDependency: {
			Available: true,
			Components: &map[string]openapi.ClusterDependencyStatus{
				DistributedSuspendComponent: {
					Available: false,
				},
			},
		},
		TensorflowDependency: {
			Available: false,
			Components: &map[string]openapi.ClusterDependencyStatus{
				DistributedSuspendComponent: {
					Available: false,
				},
			},
		},
	}
	dependenciesStatus := openapi.NewClusterDependenciesStatusWithDefaults()
	dependenciesStatus.SetOptional(optionalDependencies)
	displayedStatus.SetDependencies(*dependenciesStatus)

	if protection != nil {
		clusterConfig := openapi.NewClusterReportedConfig()
		clusterConfig.SetWorkloadOwnershipProtection(*protection)
		displayedStatus.SetConfig(*clusterConfig)
	}

	result.SetStatus(*displayedStatus)

	return result
}

func buildDisplayedClusterWithKnative(randoms *testutils.RandomValues, version string, knativeAvailable bool, knativeSubComponents *map[string]openapi.ClusterDependencyStatus) openapi.DisplayedCluster {
	result := openapi.DisplayedCluster{
		Uuid: randoms.ClusterId,
	}
	result.Version.Set(&version)

	displayedStatus := openapi.NewClusterDisplayedStatus()

	optionalDependencies := map[string]openapi.ClusterDependencyStatus{
		KnativeDependency: {
			Available: knativeAvailable,
			Components: &map[string]openapi.ClusterDependencyStatus{
				KnativeDependency: {
					Available:  knativeAvailable,
					Components: knativeSubComponents,
				},
			},
		},
	}
	dependenciesStatus := openapi.NewClusterDependenciesStatusWithDefaults()
	dependenciesStatus.SetOptional(optionalDependencies)
	displayedStatus.SetDependencies(*dependenciesStatus)

	result.SetStatus(*displayedStatus)

	return result
}
