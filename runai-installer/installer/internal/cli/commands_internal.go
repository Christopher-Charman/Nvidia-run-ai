//go:build runai_internal

/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package cli

import (
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/cli/legacy"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/cli/opts"
	"github.com/spf13/cobra"
)

// InternalBuild reports whether this binary was built with the runai_internal
// build tag.
const InternalBuild = true

func internalCommands(globals *opts.Global) []*cobra.Command {
	return legacy.Commands(globals)
}
