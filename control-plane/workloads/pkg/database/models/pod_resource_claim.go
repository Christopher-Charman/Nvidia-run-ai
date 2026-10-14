/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package models

import (
	"database/sql/driver"
	"encoding/json"
)

// PodResourceClaim is the DB-side mirror of the OpenAPI PodResourceClaim
// schema. It is stored as an element of the jsonb array in pods.resource_claims
// and carries exactly the fields cluster-sync reports: the pod-local claim
// name, the spec reference (either template-generated or standalone), and the
// actual bound ResourceClaim name once it is known.
//
// JSON field names match the API wire shape (camelCase) so raw SQL inspection
// of the jsonb column reads the same as the API response.
type PodResourceClaim struct {
	Name                       string  `json:"name"`
	ResourceClaimTemplateName  *string `json:"resourceClaimTemplateName,omitempty"`
	ResourceClaimName          *string `json:"resourceClaimName,omitempty"`
	AllocatedResourceClaimName *string `json:"allocatedResourceClaimName,omitempty"`
}

// PodResourceClaimArray adapts the slice to database/sql so GORM can persist
// it in a single jsonb column (mirrors TolerationArray / NodeAffinity).
type PodResourceClaimArray []PodResourceClaim

func (c PodResourceClaimArray) Value() (driver.Value, error) {
	valueString, err := json.Marshal(c)
	return string(valueString), err
}

// Scan tolerates a NULL column (pre-sync rows and pods with no DRA usage) and
// zero-length payloads, mirroring NodeAffinity. TolerationArray's version
// would panic on NULL because it type-asserts the value directly.
func (c *PodResourceClaimArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	valueBytes, ok := value.([]byte)
	if !ok || len(valueBytes) == 0 {
		return nil
	}
	if err := json.Unmarshal(valueBytes, &c); err != nil {
		return err
	}
	return nil
}
