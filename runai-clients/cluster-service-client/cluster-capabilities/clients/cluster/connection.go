package cluster

import (
	"net/url"

	"github.com/run-ai/runai-common-packages/logging"
	clusterService "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client/generated"
)

var serverConfig *clusterService.Configuration

var factory FactoryInterface = Factory{}

func InitConnection(clusterServiceUrlStr string, logLevel string) {
	clusterServiceUrl, err := url.Parse(clusterServiceUrlStr)
	if err != nil {
		logging.Fatalf("failed to connect to cluster-service on '%s': %v", clusterServiceUrl, err)
	}

	serverConfig = clusterService.NewConfiguration()
	serverConfig.Scheme = clusterServiceUrl.Scheme
	serverConfig.Host = clusterServiceUrl.Host
	if logLevel == "debug" {
		serverConfig.Debug = true
	}
}

func SetMock(clientFactory FactoryInterface) {
	serverConfig = &clusterService.Configuration{}
	factory = clientFactory
}

func Get() FactoryInterface {
	return factory
}
