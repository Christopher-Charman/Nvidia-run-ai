//go:build !runai_internal

/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseBuildOmitsLegacyCommands(t *testing.T) {
	root := NewRootCmd()

	assert.False(t, InternalBuild)
	assert.NotSubset(t, commandNames(root), []string{"upgrade"})
	assert.NotSubset(t, commandNames(root), []string{"uninstall"})
}

func TestReleaseBuildRejectsLegacyCommands(t *testing.T) {
	for _, name := range []string{"upgrade", "uninstall"} {
		root := NewRootCmd()
		root.SetArgs([]string{name})
		root.SetOut(nil)

		err := root.Execute()

		assert.ErrorContains(t, err, name)
	}
}
