# To use the make file please run 'make help' to see the available commands.
INFRA_PATH=../control-plane/infra

include ../control-plane/infra/makefile/index.mk

# Change for every new service
SERVICE_NAME=shared-packages
# Constants 


##@ Basic Commands 

test: ## Run unit-tests.
	@echo "Running tests..."
	 ${DOCKER_GO_COMMAND} \
	 /bin/sh -c \
	 'for name in *; do \
	 if [ ! -d "$${name}" ]; then continue; fi; \
	 echo "Testing $${name}"; \
	 cd $${name} && \
	 go test -coverprofile ./cover.out -v ./... ; \
	 cd ../; \
	 done'
	@echo "Tests completed."
.PHONY: test

lint: ## Run linter.
	# no op
.PHONY: lint

validate: ## Run validate.
	# no-op
.PHONY: validate

purify: ## Purify the code.
	${DOCKER_GO_COMMAND} \
	 /bin/sh -c \
	 'for name in *; do \
	 if [ ! -d "$${name}" ]; then continue; fi; \
	 echo "Purify $${name}"; \
	 cd $${name} && \
	 go fmt ./... && \
	 go vet ./... && \
	 go mod tidy && \
	 cd ../ ; \
	 done'
.PHONY: purify

build: ## Build the service.
	# no op
.PHONY: build

##@ Docker

docker-build: build docker-build-generic ## Build the docker image.
.PHONY: docker-build

docker-push: docker-push-generic ## Push the docker image to repository.
.PHONY: docker-push

docker-all: docker-build docker-push ## Build and push the docker image.
.PHONY: docker-all