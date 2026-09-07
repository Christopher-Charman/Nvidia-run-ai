# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

include ../../../build/makefile/runai-cluster.mk
include ../../../build/makefile/index.mk
DOCKER_WORK_DIR=${DOCKER_WORK_DIR_CLUSTER}

SERVICE_NAME=runai-crd-upgrader
GO_LDFLAGS='-s -w'

.PHONY: build
build: build-go

.PHONY: test
test: test-go

.PHONY: lint
lint: lint-go

.PHONY: format
format: format-go

.PHONY: validate
validate:
	# no-op
