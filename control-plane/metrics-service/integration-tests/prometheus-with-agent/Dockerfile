# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

FROM prom/prometheus
LABEL description="Prometheus with a promtool 'agent'"
USER nobody

ARG AGENT_PORT=8090
COPY ./bin/promtool-agent       /bin/promtool-agent
COPY ./prometheus_with_agent.sh /bin/prometheus_with_agent.sh

ENTRYPOINT ["/bin/prometheus_with_agent.sh"]
