/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package dao

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	authorization_config "github.com/dsx-ai-factory/run-ai/control-plane/authorization/pkg/config"
	authtesting "github.com/dsx-ai-factory/run-ai/control-plane/authorization/pkg/test"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb"
	"github.com/stretchr/testify/require"
)

const (
	appSubjectTypeMigrationUp   = "../migrations/up/03000003_rename_app_subject_type_to_service_account.up.sql"
	appSubjectTypeMigrationDown = "../migrations/down/03000003_rename_app_subject_type_to_service_account.down.sql"
)

type subjectTypeRow struct {
	SubjectID   string
	SubjectType string
	Deleted     bool
}

func TestAppSubjectTypeMigration(t *testing.T) {
	// Schema as it is just before the migration under test (03000002 is skipped too, so seedOneRole's role_type column exists).
	authtesting.SetupPostgres(t, "file://"+migrationsExcluding(t, "03000002", "03000003"))
	require.NoError(t, authorization_config.Init())

	resetDB(t)
	seedOneRole(t, seedRole{Name: "legacy-app-role", Type: "ALL"})

	seeded := []subjectTypeRow{
		{SubjectID: "agent-client", SubjectType: "app"},
		{SubjectID: "installer-client", SubjectType: "app"},
		{SubjectID: "removed-app", SubjectType: "app", Deleted: true},
		{SubjectID: "native-sa", SubjectType: "service-account"},
		{SubjectID: "alice", SubjectType: "user"},
		{SubjectID: "devs", SubjectType: "group"},
	}
	for _, row := range seeded {
		deletedAt := "NULL"
		if row.Deleted {
			deletedAt = "NOW()"
		}
		exec(t, `INSERT INTO access_rules (created_at, updated_at, created_by, subject_id,
		                             subject_type, role_id, scope_id, scope_type, scope_name,
		                             tenant_id, cluster_uuid, deleted_at)
		         SELECT '2020-01-01', '2020-01-01', 'seed', ?, ?, r.id, '1', 'tenant', 't', 1, '', `+deletedAt+`
		         FROM roles r WHERE r.name = 'legacy-app-role'`, row.SubjectID, row.SubjectType)
	}

	execFile(t, appSubjectTypeMigrationUp)

	want := []subjectTypeRow{
		{SubjectID: "agent-client", SubjectType: "service-account"},
		{SubjectID: "installer-client", SubjectType: "service-account"},
		{SubjectID: "removed-app", SubjectType: "service-account", Deleted: true},
		{SubjectID: "native-sa", SubjectType: "service-account"},
		{SubjectID: "alice", SubjectType: "user"},
		{SubjectID: "devs", SubjectType: "group"},
	}
	require.ElementsMatch(t, want, currentSubjectTypes(t), "every app row (including soft-deleted) becomes service-account, nothing else changes and no row is lost")

	var staleUpdatedAt int
	require.NoError(t, runaidb.DB().Raw(`SELECT count(*) FROM access_rules WHERE updated_at <> '2020-01-01'`).Row().Scan(&staleUpdatedAt))
	require.Zero(t, staleUpdatedAt, "updated_at must not be bumped so clusters are not forced to re-sync")

	var searchable int
	require.NoError(t, runaidb.DB().Raw(`SELECT count(*) FROM access_rules WHERE search_vector @@ plainto_tsquery('english', 'service-account')`).Row().Scan(&searchable))
	require.Equal(t, 4, searchable, "search_vector is regenerated from the new subject_type (3 migrated + 1 native)")

	t.Run("running up twice changes nothing", func(t *testing.T) {
		execFile(t, appSubjectTypeMigrationUp)
		require.ElementsMatch(t, want, currentSubjectTypes(t))
	})

	t.Run("down is a no-op that keeps all data", func(t *testing.T) {
		execFile(t, appSubjectTypeMigrationDown)
		require.ElementsMatch(t, want, currentSubjectTypes(t))
	})
}

func currentSubjectTypes(t *testing.T) []subjectTypeRow {
	t.Helper()
	rows, err := runaidb.DB().Raw(`SELECT subject_id, subject_type, deleted_at IS NOT NULL FROM access_rules`).Rows()
	require.NoError(t, err)
	defer rows.Close()
	var result []subjectTypeRow
	for rows.Next() {
		var r subjectTypeRow
		require.NoError(t, rows.Scan(&r.SubjectID, &r.SubjectType, &r.Deleted))
		result = append(result, r)
	}
	require.NoError(t, rows.Err())
	return result
}

// migrationsExcluding copies up/ and down/ into a temp dir, leaving out migrations whose file name contains any of skip.
func migrationsExcluding(t *testing.T, skip ...string) string {
	t.Helper()
	dir := t.TempDir()
	for _, sub := range []string{"up", "down"} {
		require.NoError(t, os.MkdirAll(filepath.Join(dir, sub), 0o755))
		entries, err := os.ReadDir(filepath.Join(migrationsDir, sub))
		require.NoError(t, err)
		for _, e := range entries {
			skipped := false
			for _, s := range skip {
				skipped = skipped || strings.Contains(e.Name(), s)
			}
			if skipped {
				continue
			}
			body, err := os.ReadFile(filepath.Join(migrationsDir, sub, e.Name()))
			require.NoError(t, err)
			require.NoError(t, os.WriteFile(filepath.Join(dir, sub, e.Name()), body, 0o644))
		}
	}
	return dir
}
