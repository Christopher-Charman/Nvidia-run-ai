-- SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
-- SPDX-License-Identifier: Apache-2.0

-- Rename the legacy 'app' subject type to 'service-account' on all access rules (including soft-deleted ones).
-- idx_subject_role_resource does not include subject_type, so an app rule and a service-account rule for the same
-- subject/role/scope can never coexist and this update cannot violate it. search_vector is a generated column and
-- is recomputed automatically. updated_at is intentionally left untouched: clusters >= v2.24 already treat both
-- values the same and older clusters are served 'app' by the sync API, so no cluster re-sync is needed.
UPDATE access_rules
SET subject_type = 'service-account'
WHERE subject_type = 'app';
