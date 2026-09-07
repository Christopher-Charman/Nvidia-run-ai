# To use the make file please run 'make help' to see the available commands.
INFRA_PATH=../control-plane/infra

include ../control-plane/infra/makefile/index.mk

GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
BLUE=\033[0;34m
NC=\033[0m
# Change for every new service
SERVICE_NAME=runai-clients
DOCKER_COMMAND=docker run --rm -w /local/${SERVICE_NAME} -v "${PWD}/..:/local" -u $(shell id -u):$(shell id -g)
# Constants 
MAX_PATCH:=100

##@ Basic Commands 

test: find-cross-imports git-config-command gocache
	@echo "Running tests..."
	${DOCKER_GO_COMMAND_LINUX} \
	/bin/sh -c '\
	  set -e; \
	  for dir in */ ; do \
	    [ -d "$${dir}" ] || continue; \
	    echo "Testing $${dir}"; \
	    cd "$${dir}" && \
	    go test -coverprofile=cover.out -v ./... || exit 1; \
	    cd - > /dev/null; \
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

bump-version-dev:
	@set -e; \
	git fetch --tags --force origin > /dev/null 2>&1; \
	latest_tag=$$(git tag -l 'runai-clients/v*' | sort -V | tail -1); \
	echo "Latest tag: $$latest_tag"; \
	\
	if echo "$$latest_tag" | grep -Eq '^runai-clients/v[0-9]+\.[0-9]+\.[0-9]+-dev\.[0-9]+$$'; then \
		base=$$(echo $$latest_tag | sed -E 's/(.*-dev\.)([0-9]+)$$/\1/'); \
		dev_ver=$$(echo $$latest_tag | sed -E 's/.*-dev\.([0-9]+)$$/\1/'); \
		new_dev_ver=$$((dev_ver + 1)); \
		new_tag="$$base$$new_dev_ver"; \
	else \
		ver=$$(echo $$latest_tag | sed -E 's/^runai-clients\/v([0-9]+)\.([0-9]+)\.([0-9]+).*$$/\1.\2.\3/'); \
		major=$$(echo $$ver | cut -d. -f1); \
		minor=$$(echo $$ver | cut -d. -f2); \
		patch=$$(echo $$ver | cut -d. -f3); \
		new_patch=$$((patch + 1)); \
		if [ "$$new_patch" -ge $(MAX_PATCH) ]; then \
			new_patch=0; \
			minor=$$((minor + 1)); \
		fi; \
		new_tag="runai-clients/v$${major}.$${minor}.$${new_patch}-dev.0"; \
	fi; \
	\
	echo "$(GREEN)Published tag: $$new_tag$(NC)";\

bump-version: _check-can-tag _bump_version

_bump_version:
	@set -e; \
	git fetch --tags --force origin > /dev/null 2>&1; \
	latest_dev_tag=$$(git tag -l 'runai-clients/v*-dev*' | sort -V | tail -1); \
	latest_prod_tag=$$(git tag -l 'runai-clients/v*' | sort -V | grep -v dev | tail -1); \
	echo "Latest dev tag: $$latest_dev_tag"; \
	echo "Latest prod tag: $$latest_prod_tag"; \
	prod_version=$$(echo $$latest_prod_tag | sed -E 's/^runai-clients\/v([0-9]+\.[0-9]+\.[0-9]+).*$$/\1/'); \
	dev_version=$$(echo $$latest_dev_tag | sed -E 's/^runai-clients\/v([0-9]+\.[0-9]+\.[0-9]+)-dev\.[0-9]+$$/\1/'); \
	\
	prod_major=$$(echo $$prod_version | cut -d. -f1); \
	prod_minor=$$(echo $$prod_version | cut -d. -f2); \
	prod_patch=$$(echo $$prod_version | cut -d. -f3); \
	\
	dev_major=$$(echo $$dev_version | cut -d. -f1); \
	dev_minor=$$(echo $$dev_version | cut -d. -f2); \
	dev_patch=$$(echo $$dev_version | cut -d. -f3); \
	\
	# ------ If dev version is next release ------ \
	\
	if [ "$$dev_major" -gt "$$prod_major" ] || \
	   ([ "$$dev_major" -eq "$$prod_major" ] && [ "$$dev_minor" -gt "$$prod_minor" ]) || \
	   ([ "$$dev_major" -eq "$$prod_major" ] && [ "$$dev_minor" -eq "$$prod_minor" ] && [ "$$dev_patch" -gt "$$prod_patch" ]); then \
		cand_major=$$dev_major; \
		cand_minor=$$dev_minor; \
		cand_patch=$$dev_patch; \
	\
	# ------ If dev version is not next release ------ \
	\
	else \
		cand_major=$$prod_major; \
		cand_minor=$$prod_minor; \
		cand_patch=$$((prod_patch + 1)); \
	fi; \
	# Apply the 100 patch rule \
	if [ "$$cand_patch" -ge $(MAX_PATCH) ]; then \
		cand_patch=0; \
		cand_minor=$$((cand_minor + 1)); \
	fi; \
	new_tag="runai-clients/v$${cand_major}.$${cand_minor}.$${cand_patch}"; \
	# git tag $$new_tag; \
	echo "$(GREEN)Published tag: $$new_tag$(NC)";



_validate-master-sync:
	@current_branch=$$(git rev-parse --abbrev-ref HEAD); \
	if [ "$$current_branch" != "master" ]; then \
		echo "$(RED)Error: You must be on branch 'master' (currently on '$$current_branch')$(NC)"; \
		exit 42; \
	fi; \
	git fetch origin master > /dev/null 2>&1; \
	if git merge-base --is-ancestor origin/master master; then \
		echo "$(GREEN)On branch 'master' and up to date with origin/master.$(NC)"; \
	else \
		echo "$(RED)Error: Local 'master' is not up to date with 'origin/master'.$(NC)"; \
		exit 42; \
	fi

_check-can-tag: _check-deps
	@set -e;
	@$(MAKE) _validate-master-sync >/dev/null 2>&1; \
	EXIT_CODE=$$?; \
	if [ "$$EXIT_CODE" -eq 0 ]; then \
		printf "$(GREEN)Branch is up to date with origin/master. You can tag a production release.$(NC)\n"; \
		exit 0; \
	fi; \
	printf "$(YELLOW)Branch is NOT up to date with origin/master. Opening a PR instead.$(NC)\n"; \
	\
	# ------ No on master PR changes to release version, Get JIRA ticket ------ \
	\
	CURRENT_BRANCH=$$(git rev-parse --abbrev-ref HEAD); \
	JIRA_TICKET=$$(echo $$CURRENT_BRANCH | grep -oE 'RUN-[0-9]{5}' || true); \
	if [ -z "$$JIRA_TICKET" ]; then \
		read -p "Enter JIRA ticket (format RUN-12345): " JIRA_TICKET_INPUT; \
		while ! echo "$$JIRA_TICKET_INPUT" | grep -qE '^RUN-[0-9]{1,7}$$'; do \
			printf "$(RED)Invalid format. Please enter in format RUN-12345:$(NC)\n"; \
			read -p "Enter JIRA ticket (format RUN-12345): " JIRA_TICKET_INPUT; \
		done; \
		JIRA_TICKET=$$JIRA_TICKET_INPUT; \
	fi; \
	\
	# ------ Checkout to new branch and apply runai-clients changes ------ \
	\
	NEW_BRANCH="$${JIRA_TICKET}-clients"; \
	printf "$(YELLOW)Target branch: $$NEW_BRANCH$(NC)\n"; \
	if git show-ref --verify --quiet refs/heads/$$NEW_BRANCH; then \
		printf "$(YELLOW)Branch $$NEW_BRANCH already exists.$(NC)\n"; \
	else \
		printf "$(YELLOW)Creating branch $$NEW_BRANCH from master...$(NC)\n"; \
		git branch $$NEW_BRANCH master >/dev/null 2>&1; \
	fi; \
	tmpdir=$$(mktemp -d); \
	git worktree add --quiet $$tmpdir $$NEW_BRANCH >/dev/null 2>&1; \
	cd $$tmpdir; \
	printf "$(YELLOW)Applying runai-clients changes from $$CURRENT_BRANCH...$(NC)\n"; \
	git checkout $$CURRENT_BRANCH -- runai-clients >/dev/null 2>&1; \
	git add runai-clients >/dev/null 2>&1; \
	if git diff --cached --quiet; then \
		printf "$(GREEN)No runai-clients changes to commit.$(NC)\n"; \
	else \
		git commit -m "Add runai-clients changes from $$CURRENT_BRANCH for $$JIRA_TICKET" >/dev/null 2>&1; \
		printf "$(GREEN)Committed runai-clients changes to $$NEW_BRANCH.$(NC)\n"; \
	fi; \
	cd - >/dev/null 2>&1; \
	git worktree remove $$tmpdir --force >/dev/null 2>&1; \
	\
	# ------ Push new branch or update existing one and create PR ------ \
	\
	printf "$(YELLOW)Pushing branch $$NEW_BRANCH to origin...$(NC)\n"; \
	git push origin $$NEW_BRANCH >/dev/null 2>&1; \
	printf "$(GREEN)Branch $$NEW_BRANCH pushed to remote.$(NC)\n"; \
	REPO=$$(git config --get remote.origin.url | sed -E 's/.*github.com[:\/](.*)\.git/\1/'); \
	printf "$(YELLOW)Checking for existing pull request for branch $$NEW_BRANCH...$(NC)\n"; \
	PR_URL=$$(gh pr list --repo $$REPO --head $$NEW_BRANCH --json url --jq '.[0].url' || true); \
	if [ -z "$$PR_URL" ]; then \
		printf "$(YELLOW)No existing PR found. Creating a new pull request...$(NC)\n"; \
		PR_URL=$$(gh pr create --repo $$REPO --head $$NEW_BRANCH --base master --title "$$JIRA_TICKET: Update runai-clients" --body "Automated PR for $$NEW_BRANCH" | grep -Eo 'https://github.com/[^ ]+'); \
		PR_EXIT_CODE=$$?; \
		if [ "$$PR_EXIT_CODE" -ne 0 ] || [ -z "$$PR_URL" ]; then \
			printf "$(RED)Error: Failed to create pull request. Please check the output above.$(NC)\n"; \
			exit 1; \
		fi; \
	else \
		printf "$(YELLOW)Existing PR found. Updating branch...$(NC)\n"; \
		git push origin $$NEW_BRANCH >/dev/null 2>&1; \
	fi; \
	printf "$(GREEN)Pull request created! - $(BLUE)$$PR_URL$(NC)\n"; \
	exit 42;

_check-deps:
	@set -e; \
	RED='\033[0;31m'; NC='\033[0m'; \
	if ! command -v gh >/dev/null 2>&1; then \
		printf "$(RED)Error: GitHub CLI (gh) is not installed or not in your PATH.$(NC)\n"; \
		printf "$(RED)Please install it using 'brew install gh' or from: https://cli.github.com/manual/installation$(NC)\n"; \
		exit 1; \
	fi
