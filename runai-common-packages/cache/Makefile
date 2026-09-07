include ../../build/makefile/index.mk

SERVICE_NAME=cache
TAG_FORMAT=runai-common-packages/$(SERVICE_NAME)

DOCKER_WORK_DIR=/local/runai-common-packages/${SERVICE_NAME}
DOCKER_COMMAND=${DOCKER_COMMAND_WORK_DIR}

MOCKERY_VERSION = v2.53.3
BIN_DIR := $(shell pwd)/bin
MOCKERY := ${BIN_DIR}/mockery

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
	docker image rm cache-tests-redis --force
	docker image rm cache-tests-nats --force
	env GOOS=linux GOARCH=amd64 GOCACHE=/tmp/.cache GOLANGCI_LINT_CACHE=/tmp/.cache CGO_ENABLED=0 GO111MODULE=on \
		go test -c -v -tags integration ./integration-tests/tests/ -o ./integration-tests/bin/integ-tests 
	docker compose -f ./integration-tests/docker-compose.yaml build 
		
integration-tests: build-integration-tests
	docker compose -f ./integration-tests/docker-compose.yaml --profile inmemory-test up --remove-orphans --force-recreate --abort-on-container-exit
	docker compose -f ./integration-tests/docker-compose.yaml --profile redis-setup --profile redis-test up --remove-orphans --force-recreate --abort-on-container-exit
	docker compose -f ./integration-tests/docker-compose.yaml --profile nats-setup --profile nats-test up --remove-orphans --force-recreate --abort-on-container-exit
	docker compose -f ./integration-tests/docker-compose.yaml --profile keydb-setup --profile keydb-test up --remove-orphans --force-recreate --abort-on-container-exit
	docker compose -f ./integration-tests/docker-compose.yaml --profile valkey-setup --profile valkey-test up --remove-orphans --force-recreate --abort-on-container-exit

integration-tests-debug:
	docker compose -f ./integration-tests/docker-compose.yaml --profile redis-setup  --profile valkey-setup  --profile nats-setup --profile keydb-setup up --remove-orphans --force-recreate --abort-on-container-exit

validate_mockery:
	@mkdir -p ${BIN_DIR}
	@if [ -e ${MOCKERY} ] && [ "$$(${MOCKERY} --version | awk '{print $$NF}')" != "${MOCKERY_VERSION}" ]; then \
		echo "existing mockery binary is not of required version ${MOCKERY_VERSION}. Reinstalling it"; \
		rm -f ${MOCKERY}; \
	fi
	GOBIN=${BIN_DIR} go install github.com/vektra/mockery/v2/...@${MOCKERY_VERSION}

mockery: validate_mockery
	${MOCKERY}