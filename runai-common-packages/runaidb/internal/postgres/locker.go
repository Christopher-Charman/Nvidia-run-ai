package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/config"
	dbconfig "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/config"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm"
)

type Locker struct {
	db      orm.DB
	retries int
}

func NewLocker(db orm.DB) *Locker {
	conf := config.Get[*dbconfig.Config]()
	return &Locker{
		db:      db,
		retries: conf.Migration.LockRetries,
	}
}

func (pl *Locker) Acquire(schema string) (bool, error) {
	query := "SELECT pg_try_advisory_lock(to_regnamespace($1)::integer)"
	maxAttempts := pl.retries + 1
	var err error
	for i := 0; i < maxAttempts; i++ {
		var acquired sql.NullBool
		_, err = pl.db.Raw(query, schema).Scan(&acquired)
		if err == nil && acquired.Valid && acquired.Bool {
			return true, nil
		}

		if err != nil {
			err = fmt.Errorf("failed to acquire lock (attempt %d/%d): %w", i+1, maxAttempts, err)
		}

		if i+1 < maxAttempts {
			time.Sleep(500 * time.Millisecond)
		}
	}

	return false, err
}

func (pl *Locker) Release(schema string) (bool, error) {
	var released sql.NullBool
	query := "SELECT pg_advisory_unlock(to_regnamespace($1)::integer)"
	_, err := pl.db.Raw(query, schema).Scan(&released)
	return released.Valid && released.Bool, err
}
