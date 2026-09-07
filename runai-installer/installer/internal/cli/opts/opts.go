/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Package opts holds the values bound to the root command's persistent flags.
// It lives in its own package so that command packages can read them without
// importing the cli package that owns the root command.
package opts

import "time"

type Global struct {
	DevMode     bool
	LogsDir     string
	StepTimeout time.Duration
}
