/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package post_hooks

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// webhookCallFailurePrefix is how the API server reports a transport-level failure
// calling an admission webhook (k8s.io/apiserver/pkg/util/webhook.ErrCallingWebhook),
// wrapped in an InternalError when the webhook's failurePolicy is Fail.
const webhookCallFailurePrefix = "failed calling webhook"

// Retry settings for legacy object deletes that fail because the API server could
// not reach the admission webhook guarding them. Package variables so tests can
// shorten them.
var (
	webhookRetryInterval = 5 * time.Second
	webhookRetryTimeout  = 3 * time.Minute
)

// deleteLegacyObject deletes a legacy department or project, retrying while the
// API server cannot reach the admission webhook that validates the delete.
//
// During an upgrade the previous release's validating webhook (for example
// runai-project-validation) can still point at a Service port the new operator no
// longer serves; init-ca re-points it asynchronously. Until then every delete fails
// with an InternalError "failed calling webhook ...". The delete is idempotent, so
// it is retried every webhookRetryInterval for up to webhookRetryTimeout. Any other
// error is returned immediately. NotFound and NoMatch mean the object is already
// gone and count as success.
func deleteLegacyObject(ctx context.Context, k8sClient client.Client, obj client.Object, kind string, logger logr.Logger) error {
	logger.Info("Deleting legacy "+kind, "name", obj.GetName())
	deadline := time.Now().Add(webhookRetryTimeout)
	for {
		err := k8sClient.Delete(ctx, obj)
		if err == nil || apierrors.IsNotFound(err) || meta.IsNoMatchError(err) {
			return nil
		}
		if !isWebhookCallFailure(err) {
			return fmt.Errorf("failed to delete legacy %s %q: %w", kind, obj.GetName(), err)
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("failed to delete legacy %s %q: admission webhook still unreachable after %s: %w", kind, obj.GetName(), webhookRetryTimeout, err)
		}
		logger.Info("Admission webhook not reachable yet, retrying legacy delete", "kind", kind, "name", obj.GetName(), "error", err.Error())
		select {
		case <-ctx.Done():
			return fmt.Errorf("failed to delete legacy %s %q: %w", kind, obj.GetName(), ctx.Err())
		case <-time.After(webhookRetryInterval):
		}
	}
}

// isWebhookCallFailure reports whether err is the API server failing to call an
// admission webhook, as opposed to the webhook rejecting the request.
func isWebhookCallFailure(err error) bool {
	return apierrors.IsInternalError(err) && strings.Contains(err.Error(), webhookCallFailurePrefix)
}
