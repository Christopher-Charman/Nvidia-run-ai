/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package steps

import (
	"context"
	"fmt"

	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/config"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/model"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/orchestrator2"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/orchestrator2/actions"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/orchestrator2/spec"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/step"
)

type ApplyAirgapRegistrySecrets struct {
	*step.RunAIInstallStepImpl
	upgradeCfg *config.UpgradeConfig
	k8s        orchestrator2.K8sExecutor
}

func NewApplyAirgapRegistrySecrets(upgradeCfg *config.UpgradeConfig, k8s orchestrator2.K8sExecutor) *ApplyAirgapRegistrySecrets {
	return &ApplyAirgapRegistrySecrets{
		RunAIInstallStepImpl: step.NewRunAIInstallStepImpl("Apply airgap registry secrets", upgradeCfg.StepTimeout),
		upgradeCfg:           upgradeCfg,
		k8s:                  k8s,
	}
}

func (s *ApplyAirgapRegistrySecrets) GetStatusAsync(_ *model.RunAIInstallData) model.InstallPhase {
	return s.GetCachedStatus()
}

func (s *ApplyAirgapRegistrySecrets) Execute(ctx context.Context, _ *model.RunAIInstallData) error {
	registry := &spec.Registry{
		Host:     s.upgradeCfg.RegistryURL,
		Username: s.upgradeCfg.RegistryUsername,
		Password: s.upgradeCfg.RegistryPassword,
	}
	dockerConfig, ok := registry.ImagePullConfig()
	if !ok {
		return fmt.Errorf("registry credentials are required for an airgap bundle upgrade")
	}

	if s.k8s == nil {
		return fmt.Errorf("kubernetes executor is required for an airgap bundle upgrade")
	}

	deps := orchestrator2.Deps{K8sExecutor: s.k8s}

	for _, namespace := range []string{step.CPNamespace, step.ClusterNamespace} {
		action := actions.RegistrySecret{
			SecretName: "runai-reg-creds",
			Namespace:  namespace,
			Token:      dockerConfig,
		}
		if _, err := action.Execute(ctx, deps); err != nil {
			return err
		}
	}

	s.SetStatus(model.Installed, nil)

	return nil
}
