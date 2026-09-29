-- SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
-- SPDX-License-Identifier: Apache-2.0

-- we do not want this index to be unique because it prevents two clusters to have
-- the same policy name for no reason. duplicate names may appear in the policies
-- grid, but this has no impact on functionality; users decide whether to reuse names.
DROP INDEX IF EXISTS by_name;
CREATE INDEX IF NOT EXISTS by_name ON policies(tenant_id, workload_type, name);
