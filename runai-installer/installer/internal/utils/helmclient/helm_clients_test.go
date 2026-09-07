/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package helmclient

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOCIRegistryHost(t *testing.T) {
	assert.Equal(t, "nvcr.io", ociRegistryHost("oci://nvcr.io/nvidia/runai/control-plane"))
	assert.Equal(t, "nvcr.io", ociRegistryHost("nvcr.io/nvidia/runai/control-plane"))
	assert.Equal(t, "ghcr.io", ociRegistryHost("oci://ghcr.io/run-ai/chart"))
	assert.Equal(t, "nvcr.io", ociRegistryHost("oci://nvcr.io"))
}

func TestChartAuth_IsZero(t *testing.T) {
	assert.True(t, ChartAuth{}.IsZero())
	assert.False(t, ChartAuth{Username: "$oauthtoken"}.IsZero())
	assert.False(t, ChartAuth{Password: "secret"}.IsZero())
}

func TestHelmActionTimeout(t *testing.T) {
	t.Run("uses fallback without context deadline", func(t *testing.T) {
		assert.Equal(t, 15*time.Minute, helmActionTimeout(context.Background(), 15*time.Minute))
	})

	t.Run("uses context deadline", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Minute)
		defer cancel()

		timeout := helmActionTimeout(ctx, 15*time.Minute)

		assert.Greater(t, timeout, 24*time.Minute)
		assert.LessOrEqual(t, timeout, 25*time.Minute)
	})
}
