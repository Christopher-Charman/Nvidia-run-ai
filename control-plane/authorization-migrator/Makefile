# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

include ../../build/makefile/control-plane.mk
include ../../build/makefile/index.mk

# Change for every new service
SERVICE_NAME=authorization-migrator
##@ Basic Commands

test: test-go ## Run unit-tests.
.PHONY: test

lint: lint-go ## Run linter.
.PHONY: lint

validate: ## Run validate.
	# no-op
.PHONY: validate

purify: purify-go ## Purify the code.
.PHONY: purify

build: build-go ## Build the service.
.PHONY: build
