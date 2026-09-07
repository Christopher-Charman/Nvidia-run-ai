/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package cli

import (
	"errors"
	"os"
	"time"

	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/cli/opts"
	"github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/specrun"
	"github.com/spf13/cobra"
)

var globals opts.Global

func NewRootCmd() *cobra.Command {
	short := "NVIDIA Run:ai Installer"
	if InternalBuild {
		short += " (internal build)"
	}

	root := &cobra.Command{
		Use:          "installer",
		Short:        short,
		SilenceUsage: true,
	}

	pf := root.PersistentFlags()
	pf.BoolVar(&globals.DevMode, "dev", false, "development mode")
	pf.StringVar(&globals.LogsDir, "logs-dir", "", "directory for installer logs")
	pf.DurationVar(&globals.StepTimeout, "step-timeout", 7*time.Minute, "timeout per step")

	root.AddCommand(newInstallCmd())
	root.AddCommand(newAirgapCmd())

	for _, cmd := range internalCommands(&globals) {
		root.AddCommand(cmd)
	}

	return root
}

func Execute() {
	err := NewRootCmd().Execute()
	if err == nil {
		return
	}
	if errors.Is(err, specrun.ErrPreflightFailed) {
		os.Exit(2)
	}
	os.Exit(1)
}
