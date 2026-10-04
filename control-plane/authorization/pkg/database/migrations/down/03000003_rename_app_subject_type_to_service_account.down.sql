-- SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
-- SPDX-License-Identifier: Apache-2.0

-- Intentionally a no-op. 'app' and 'service-account' rows are indistinguishable after the up migration, and
-- previous releases already read, filter and delete both values as equivalent (and the sync API maps
-- service-account back to 'app' for clusters older than v2.24), so rolling back the application code needs no
-- data change. Converting every service-account row back to 'app' would also rewrite rows that were created
-- natively as service-account.
SELECT 1;
