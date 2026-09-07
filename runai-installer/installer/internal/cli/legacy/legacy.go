/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Package legacy holds the pre-orchestrator2 upgrade and uninstall commands.
// They are kept for internal use only: the released installer is
// built without the runai_internal build tag, and nothing in the default build
// imports this package.
package legacy

import (
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/cli/opts"
	"github.com/spf13/cobra"
)

// Commands returns the legacy commands, reading the root command's persistent
// flags through g when they run.
func Commands(g *opts.Global) []*cobra.Command {
	return []*cobra.Command{
		newUpgradeCmd(g),
		newUninstallCmd(g),
	}
}
