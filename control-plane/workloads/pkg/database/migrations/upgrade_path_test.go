/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package migrations

import (
	"os"
	"strconv"
	"strings"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb"
	dbconfig "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/config"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"
)

const (
	// lastSharedV223Version is the newest migration v2.23 and this branch have in common.
	lastSharedV223Version = "2230018"
	// v223BackportVersion is the v2.23-only backport of RUN-39587 (forward-ported here
	// as 02240011). A database upgraded from v2.23 reports this version, and
	// golang-migrate refuses to run Up when the current version has no file.
	v223BackportVersion = "2230019"
)

var _ = Describe("Upgrade path from a v2.23 database", func() {
	It("migrates to the latest version when the database is at the v2.23 backport version", func() {
		Expect(runaidb.Migrate(&runaidb.MigrationParams{
			Direction:     dbconfig.MigrationDirectionUp,
			TargetVersion: ptr.To(lastSharedV223Version),
		})).To(Succeed())
		Expect(runaidb.Migrate(&runaidb.MigrationParams{
			Direction:     dbconfig.MigrationDirectionForce,
			TargetVersion: ptr.To(v223BackportVersion),
		})).To(Succeed())

		Expect(runaidb.Migrate(&runaidb.MigrationParams{Direction: dbconfig.MigrationDirectionUp})).To(Succeed())

		Expect(currentVersion()).To(Equal(latestUpVersion()))
	})
})

// currentVersion reads the version golang-migrate recorded for the workloads schema.
func currentVersion() uint64 {
	var version uint64
	Expect(runaidb.DB().Actual().Raw("SELECT version FROM workloads.schema_migrations").Scan(&version).Error).To(Succeed())
	return version
}

// latestUpVersion returns the highest version among the up/ migration files.
func latestUpVersion() uint64 {
	entries, err := os.ReadDir("up")
	Expect(err).ToNot(HaveOccurred())
	var latest uint64
	for _, entry := range entries {
		prefix, _, found := strings.Cut(entry.Name(), "_")
		Expect(found).To(BeTrue(), "migration file name %q has no version prefix", entry.Name())
		version, err := strconv.ParseUint(prefix, 10, 64)
		Expect(err).ToNot(HaveOccurred(), "migration file name %q has a non-numeric version", entry.Name())
		if version > latest {
			latest = version
		}
	}
	return latest
}
