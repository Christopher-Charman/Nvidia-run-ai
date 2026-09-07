# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

FROM registry.access.redhat.com/ubi9/ubi-minimal
LABEL name="Run.AI Karta Migration"
LABEL description="Run.AI Karta migration job"
LABEL summary="Migrates legacy ResourceInterface CRs to the new Karta CRD"
LABEL vendor="Runai Labs LTD"
LABEL maintainer=support@run.ai
ARG TARGETARCH
ENV TARGETARCH=$TARGETARCH

WORKDIR /app
ADD bin/karta-migrate-$TARGETARCH karta-migrate

USER 65532:65532

ENTRYPOINT ["/app/karta-migrate"]
