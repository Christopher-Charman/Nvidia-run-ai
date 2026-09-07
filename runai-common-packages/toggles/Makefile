include ../../build/makefile/index.mk

SERVICE_NAME=toggles
TAG_FORMAT=runai-common-packages/$(SERVICE_NAME)

DOCKER_WORK_DIR=/local/runai-common-packages/${SERVICE_NAME}
DOCKER_COMMAND=${DOCKER_COMMAND_WORK_DIR}

MOCKERY_VERSION = v2.53.3
BIN_DIR := $(shell pwd)/bin
MOCKERY := ${BIN_DIR}/mockery

all: purify test
.PHONY: all

test: test-go integration-tests
.PHONY: test

lint: 
	# no op
.PHONY: lint

validate:
	# no op
.PHONY: validate

build:
	# no op
.PHONY: build

# Integration tests
build-integration-tests:
	docker image rm toggles-test --force
	env GOOS=linux GOARCH=amd64 GOCACHE=/tmp/.cache GOLANGCI_LINT_CACHE=/tmp/.cache CGO_ENABLED=0 GO111MODULE=on \
		go test -c -v -tags integration ./integration-tests/tests/ -o ./integration-tests/bin/integ-tests
	docker compose -f ./integration-tests/docker-compose.yaml build
.PHONY: build-integration-tests

integration-tests: build-integration-tests
	docker compose -f ./integration-tests/docker-compose.yaml --profile unleash-setup --profile unleash-test up --remove-orphans --force-recreate --abort-on-container-exit
.PHONY: integration-tests

integration-tests-debug:
	docker compose -f ./integration-tests/docker-compose.yaml --profile unleash-setup up --remove-orphans --force-recreate --abort-on-container-exit
.PHONY: integration-tests-debug

purify:
	go fmt ./...
	go vet ./...
	go mod tidy
.PHONY: purify

validate_mockery:
	@mkdir -p ${BIN_DIR}
	@if [ -e ${MOCKERY} ] && [ "$$(${MOCKERY} --version | awk '{print $$NF}')" != "${MOCKERY_VERSION}" ]; then \
		echo "existing mockery binary is not of required version ${MOCKERY_VERSION}. Reinstalling it"; \
		rm -f ${MOCKERY}; \
	fi
	GOBIN=${BIN_DIR} go install github.com/vektra/mockery/v2/...@${MOCKERY_VERSION}

mockery: validate_mockery
	${MOCKERY}
.PHONY: mockery
