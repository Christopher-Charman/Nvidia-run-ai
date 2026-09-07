module github.com/dsx-ai-factory/run-ai/control-plane/bff-service/pkg/client

go 1.23.0

toolchain go1.24.6

require (
	github.com/run-ai/runai-common-packages/runaictx v1.1.3
	github.com/run-ai/runai-common-packages/tracing v1.0.20
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.62.0 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	go.opentelemetry.io/otel/metric v1.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.37.0 // indirect
)

replace github.com/dsx-ai-factory/run-ai/control-plane/bff-service/pkg/client => ./pkg/client
