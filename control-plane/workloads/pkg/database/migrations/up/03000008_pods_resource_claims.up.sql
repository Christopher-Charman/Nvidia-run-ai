-- SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
-- SPDX-License-Identifier: Apache-2.0

-- Per-pod DRA usage, reported as-is by cluster-sync.
-- One jsonb array per pod; each element mirrors a spec.resourceClaims[] entry
-- enriched with the actual bound ResourceClaim name once status publishes it.
-- GIN index with jsonb_path_ops supports the reverse-view containment query
-- (workloads referencing a template name), which is the only lookup pattern
-- planned for this column.
ALTER TABLE pods ADD COLUMN IF NOT EXISTS resource_claims jsonb;

CREATE INDEX IF NOT EXISTS idx_pods_resource_claims
    ON pods USING gin (resource_claims jsonb_path_ops);
