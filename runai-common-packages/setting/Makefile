include ../../build/makefile/index.mk

SERVICE_NAME=setting
TAG_FORMAT=runai-common-packages/$(SERVICE_NAME)

DOCKER_WORK_DIR=/local/runai-common-packages/${SERVICE_NAME}
DOCKER_COMMAND=${DOCKER_COMMAND_WORK_DIR}

MOCKERY_VERSION = v2.53.3
BIN_DIR := $(shell pwd)/bin
MOCKERY := ${BIN_DIR}/mockery

test: test-go
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
