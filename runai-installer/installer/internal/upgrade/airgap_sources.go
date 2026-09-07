/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package upgrade

import (
	"fmt"
	"path/filepath"

	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/airgapvalues"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/artifacts"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/config"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/step"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/utils/airgap"
)

// AirgapChartAndValues resolves the chart archive and the Helm values file an
// air-gapped upgrade installs a component from. The two air-gap formats keep
// them in different places:
//
//   - a released package, extracted by this upgrade: charts sit next to a single
//     custom-env.yaml that covers both components
//   - an artifact bundle prepared by `airgap load`: charts are pinned by the
//     bundle's lock and each component has its own generated values file
//
// Everything the upgrade does with them afterwards is identical.
func AirgapChartAndValues(cfg *config.UpgradeConfig, chartName string) (chartPath, valuesPath string, err error) {
	if !cfg.UsesAirgapBundle() {
		extracted := cfg.AirgapExtractedDir
		return airgap.FindChartFile(extracted, chartName), filepath.Join(extracted, airgap.CustomEnvFile), nil
	}

	chartPath, err = artifacts.LockedChartPath(cfg.AirgapWorkingDir, chartName)
	if err != nil {
		return "", "", err
	}
	valuesFile, err := bundleValuesFile(chartName)
	if err != nil {
		return "", "", err
	}
	return chartPath, filepath.Join(cfg.EffectiveAirgapValuesDir(), valuesFile), nil
}

func bundleValuesFile(chartName string) (string, error) {
	switch chartName {
	case step.CPChartName:
		return airgapvalues.ControlPlaneValuesFile, nil
	case step.ClusterChartName:
		return airgapvalues.ClusterValuesFile, nil
	}
	return "", fmt.Errorf("no air-gapped values file for chart %q", chartName)
}
