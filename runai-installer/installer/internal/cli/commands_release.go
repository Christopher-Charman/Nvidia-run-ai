//go:build !runai_internal

/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package cli

import (
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/cli/opts"
	"github.com/spf13/cobra"
)

// InternalBuild reports whether this binary was built with the runai_internal
// build tag.
const InternalBuild = false

// internalCommands returns nothing in a released build: the legacy commands are
// not registered, and the legacy packages behind them are not linked in.
func internalCommands(*opts.Global) []*cobra.Command {
	return nil
}
