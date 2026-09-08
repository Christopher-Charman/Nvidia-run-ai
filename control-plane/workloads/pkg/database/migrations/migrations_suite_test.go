/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package migrations

import (
	"testing"

	"github.com/dsx-ai-factory/run-ai/control-plane/workloads/pkg/config"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/integdb"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"
)

// TestMigrations starts an empty Postgres and initializes runaidb without the
// automatic migration, so each spec drives the migrator through the exact
// version path it wants to check.
func TestMigrations(t *testing.T) {
	RegisterFailHandler(Fail)
	config.Init()
	conf, cleanup, err := integdb.SetupPostgresDB(&integdb.IntegdbConfig{
		Schema:           "workloads",
		MigrationDirPath: "file://../migrations",
	})
	if err != nil {
		t.Fatalf("failed to set up test postgres: %v", err)
	}
	t.Cleanup(cleanup)
	conf.Migration.Auto = ptr.To(false)
	conf.UseServiceSession = new(bool)
	conf.SkipDefaultTransaction = true
	if err := runaidb.Init(conf); err != nil {
		t.Fatalf("failed to init db: %v", err)
	}
	t.Cleanup(func() { _ = runaidb.Close() })

	RunSpecs(t, "Workloads SQL migrations")
}
