/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package helm_deployer

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/sdk/apis/run/v1alpha1"
)

const (
	releaseDesiredFingerprintLabel = "runai-desired-fingerprint"
	releaseConfigFingerprintLabel  = "runai-config-fingerprint"
)

func fingerprint(value interface{}) (string, error) {
	if values, ok := value.(map[string]interface{}); ok && len(values) == 0 {
		value = nil
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("encode Helm release fingerprint: %w", err)
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:16]), nil
}

func desiredFingerprint(
	chartURL string,
	values map[string]interface{},
	bindings []v1alpha1.HelmCredentialBinding,
) (string, error) {
	if len(values) == 0 {
		values = nil
	}
	if len(bindings) == 0 {
		bindings = nil
	}
	return fingerprint(struct {
		ChartURL string
		Values   map[string]interface{}
		Bindings []v1alpha1.HelmCredentialBinding
	}{ChartURL: chartURL, Values: values, Bindings: bindings})
}

func fingerprintLabels(params *InstallParams, appliedValues map[string]interface{}) (map[string]string, error) {
	chartURL := params.chartUrl
	if chartURL == "" {
		chartURL = ChartUrl(&v1alpha1.HelmSpec{ChartRef: params.chartRef})
	}
	desired, err := desiredFingerprint(chartURL, params.values, params.credentialBindings)
	if err != nil {
		return nil, err
	}
	config, err := fingerprint(appliedValues)
	if err != nil {
		return nil, err
	}
	return map[string]string{releaseDesiredFingerprintLabel: desired, releaseConfigFingerprintLabel: config}, nil
}
