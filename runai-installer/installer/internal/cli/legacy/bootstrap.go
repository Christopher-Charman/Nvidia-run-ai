/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package legacy

import "github.com/dsx-ai-factory/run-ai/runai-installer/installer/internal/logging"

func logBanner(phase string) {
	logging.Infof("================================================================================")
	logging.Infof("=              NVIDIA Run:ai Installer - %-20s              =", phase)
	logging.Infof("================================================================================")
}
