# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

include ../../build/makefile/control-plane.mk
include ../../build/makefile/index.mk

# Change for every new service
SERVICE_NAME=metrics-service

OPENAPI_GENERATOR_VERSION_SERVER=v7.12.0
OPENAPI_GENERATOR_VERSION_CLIENT_GO=v7.12.0

VERSIONED_CLIENT=true

GITHUB_REPO_NAME=run-ai
SHARED_PACKAGES_PATH=pkg
OPENAPI_CLIENT_NAME_GO=client
SHARED_PACKAGES_PATH_ABS=control-plane/metrics-service/pkg

test: test-ginkgo integration-tests ## Run unit and integration tests
.PHONY: test

lint: lint-go ## Run linter.
.PHONY: lint

validate: openapi-validate ## Run validate.
.PHONY: validate

purify: purify-go ## Purify the code.
.PHONY: purify

build: openapi-build build-go generate-mocks-client ## Build the service.
.PHONY: build

##@ OpenAPI

openapi-lint: openapi-lint-generic ## Lint the swagger bundle.
.PHONY: openapi-lint

build-integration-tests:
	# Two independent Go builds; both images embed their output, so the compose
	# build below has to follow them.
	$(MAKE) -j2 _build-promtool-agent build-integration-tests-generic
	# Every service sits behind a profile, so an unfiltered `build` selects
	# nothing and the images get built serially by `up` instead.
	AGENT_PORT=8090 docker compose --profile "*" -f ./integration-tests/docker-compose.yaml build

#@ Integration tests
integration-tests: _compose-pull build-for-test
	AGENT_PORT=8090 docker compose --profile "*" -f ./integration-tests/docker-compose.yaml down --volumes --remove-orphans
	AGENT_PORT=8090 docker compose --profile "*" -f ./integration-tests/docker-compose.yaml up --remove-orphans --force-recreate --abort-on-container-exit

# used for launching docker compose without the integration tests docker, in order to run them from IDE
integration-tests-debug:
	AGENT_PORT=8090 docker compose --profile run-setup -f ./integration-tests/docker-compose.yaml up --remove-orphans --force-recreate --abort-on-container-exit

_build-promtool-agent:
	go fmt ./integration-tests/prometheus-with-agent/...
	${DOCKER_GO_BASE_COMMAND} -e GOOS=linux -e GOARCH=amd64 ${DOCKERHUB_PROXY}library/golang:${GO_VERSION} go build -o ./integration-tests/prometheus-with-agent/bin/promtool-agent -buildvcs=false ./integration-tests/prometheus-with-agent/main/...
	chmod +x ./integration-tests/prometheus-with-agent/bin/promtool-agent

# Generates the same code as `build` but compiles a single binary: compose
# builds the image for the host platform, so only that architecture is ever
# copied in, and the second binary `build` produces is dead weight here.
ifeq ($(ARCH),arm64)
BUILD_GO_FOR_TEST=build-go-arm
else
BUILD_GO_FOR_TEST=build-go-amd
endif

build-for-test:
	$(MAKE) -j2 $(BUILD_GO_FOR_TEST) generate-mocks-client build-integration-tests
.PHONY: build-for-test

# The ginkgo recipe lives on these targets. Depending from `test-ginkgo` would
# leave the recipe free to run beside `build-for-test` under `make -j`.
test-ginkgo-local test-ginkgo-privileged: build-for-test

# Purely an optimization — `up` pulls anything missing anyway, so a failure here
# (an old compose without --ignore-buildable, a flaky registry) must not fail the build.
_compose-pull:
	-AGENT_PORT=8090 docker compose --profile "*" -f ./integration-tests/docker-compose.yaml pull --ignore-buildable
.PHONY: _compose-pull

local-deploy: export OS=linux
local-deploy: export ARCH=amd64
local-deploy: openapi-generate-code build-go docker-build-generic docker-push-generic
    ## Build and push the docker image.
.PHONY: local-deploy
