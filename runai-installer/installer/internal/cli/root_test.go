/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package cli

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func commandNames(root *cobra.Command) []string {
	names := make([]string, 0, len(root.Commands()))
	for _, cmd := range root.Commands() {
		names = append(names, cmd.Name())
	}
	return names
}

func TestRootRegistersReleasedCommands(t *testing.T) {
	root := NewRootCmd()

	assert.Subset(t, commandNames(root), []string{"install", "airgap"})
}

func TestInstallCommandKeepsInstallNewAlias(t *testing.T) {
	root := NewRootCmd()

	cmd, _, err := root.Find([]string{"install-new"})

	require.NoError(t, err)
	assert.Equal(t, "install", cmd.Name())
}
