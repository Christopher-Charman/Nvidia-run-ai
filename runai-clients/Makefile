# To use the make file please run 'make help' to see the available commands.
INFRA_PATH=../control-plane/infra

include ../control-plane/infra/makefile/index.mk

# Change for every new service
SERVICE_NAME=runai-clients
DOCKER_COMMAND=docker run --rm -w /local/${SERVICE_NAME} -v "${PWD}/..:/local" -u $(shell id -u):$(shell id -g)
# Constants 


##@ Basic Commands 

test: find-cross-imports
	# @echo "Running tests..."
	# @${DOCKER_GO_COMMAND} \
	# /bin/sh -c '\
	#   set -e; \
	#   for dir in */ ; do \
	#     [ -d "$${dir}" ] || continue; \
	#     echo "Testing $${dir}"; \
	#     cd "$${dir}" && \
	#     go test -coverprofile=cover.out -v ./... || exit 1; \
	#     cd - > /dev/null; \
	#   done'
	# @echo "Tests completed."
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

find-cross-imports:
	@echo "Searching for cross-imports in runai-clients..."
	@found=0; \
	for dir in */ ; do \
	  [ -d "$$dir" ] || continue; \
	  base=$$(basename $$dir); \
	  if grep -Pr --include='*.go' --include='*.mod' --include='*.sum' \
	    '"github.com/dsx-ai-factory/run-ai/runai-clients/(?!'"$$base"')(.*)"' ./$$dir ; \
	  then \
	    found=1; \
	  fi; \
	done; \
	if [ "$$found" -eq 1 ]; then \
	  echo "Cross-imports were found and are not allowed!"; \
	  exit 1; \
	else \
	  echo "No cross-imports found."; \
	fi
.PHONY: find-cross-imports
