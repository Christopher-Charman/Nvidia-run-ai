/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package controller

import (
	"context"
	"errors"

	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/sdk/apis/run/v1alpha1"
	"github.com/rs/zerolog/log"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const credentialBindingsIndex = "run.ai/application-credential-secret" //nolint:gosec // Cache index, not a secret.

var errCredentialBindingsUnavailable = errors.New("application credential Secrets unavailable")

func credentialSecretNames(obj client.Object) []string {
	application, ok := obj.(*v1alpha1.Application)
	if !ok || application.Spec.Helm == nil {
		return nil
	}
	names := make([]string, 0, len(application.Spec.Helm.CredentialBindings))
	for _, binding := range application.Spec.Helm.CredentialBindings {
		names = append(names, binding.SecretName)
	}
	return names
}

func (c *ApplicationController) mapCredentialSecretToApplications(
	ctx context.Context,
	secret client.Object,
) []reconcile.Request {
	var applications v1alpha1.ApplicationList
	if err := c.client.List(ctx, &applications,
		client.InNamespace(secret.GetNamespace()),
		client.MatchingFields{credentialBindingsIndex: secret.GetName()},
	); err != nil {
		log.Error().Err(err).Msg("failed to list applications for credential Secret")
		return nil
	}
	requests := make([]reconcile.Request, 0, len(applications.Items))
	for i := range applications.Items {
		requests = append(requests, reconcile.Request{NamespacedName: client.ObjectKeyFromObject(&applications.Items[i])})
	}
	return requests
}
