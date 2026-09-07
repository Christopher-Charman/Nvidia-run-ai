//go:build runai_internal

/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInternalBuildRegistersLegacyCommands(t *testing.T) {
	root := NewRootCmd()

	assert.True(t, InternalBuild)
	assert.Subset(t, commandNames(root), []string{"install", "upgrade", "uninstall", "airgap"})
}

func TestLegacyCommandsReadRootPersistentFlags(t *testing.T) {
	root := NewRootCmd()
	require.NoError(t, root.PersistentFlags().Set("dev", "true"))

	assert.True(t, globals.DevMode)
}
