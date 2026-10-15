/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package models

import (
	"database/sql/driver"
	"encoding/json"
)

type PodResourceClaimArray []PodResourceClaim

func (c PodResourceClaimArray) Value() (driver.Value, error) {
	valueString, err := json.Marshal(c)
	return string(valueString), err
}

// Scan tolerates a SQL NULL input (unlike TolerationArray's naive version)
// because resource_claims is added to the live pods table and all pre-upgrade
// rows read as NULL until something writes them. Mirrors NodeAffinity's
// defensive handling, added here for the same reason.
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
