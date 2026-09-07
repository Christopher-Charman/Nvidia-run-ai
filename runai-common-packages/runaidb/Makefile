include ../../build/makefile/index.mk

SERVICE_NAME=runaidb
TAG_FORMAT=runai-common-packages/$(SERVICE_NAME)

DOCKER_WORK_DIR=/local/runai-common-packages/${SERVICE_NAME}
DOCKER_COMMAND=${DOCKER_COMMAND_WORK_DIR}

test: git-config-command
	go test -v -coverprofile=coverage.out ./...
	make integration-tests
.PHONY: test

integration-tests:
	cd integration-tests && go test -v ./...
.PHONY: integration-tests

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