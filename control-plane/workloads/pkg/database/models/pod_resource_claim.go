/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package models

import (
	"database/sql/driver"
	"encoding/json"
)

type PodResourceClaim struct {
	Name                       string  `json:"name"`
	ResourceClaimTemplateName  *string `json:"resource_claim_template_name,omitempty"`
	ResourceClaimName          *string `json:"resource_claim_name,omitempty"`
	AllocatedResourceClaimName *string `json:"allocated_resource_claim_name,omitempty"`
}

func (c PodResourceClaim) Value() (driver.Value, error) {
	valueString, err := json.Marshal(c)
	return string(valueString), err
}

func (c *PodResourceClaim) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &c); err != nil {
		return err
	}
	return nil
}
