package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewLocker(t *testing.T) {
	db := &mocks.DB{}
	pl := NewLocker(db)
	require.NotNil(t, pl)
}

func TestAcquire_Success(t *testing.T) {
	db := &mocks.DB{}
	db.On("Raw", "SELECT pg_try_advisory_lock(to_regnamespace($1)::integer)", "some-schema").Return(db)
	db.On("Scan", matchNullBoolScan()).Run(scanNullBool(true, true)).Return(int64(0), nil)

	pl := Locker{db: db}
	acquired, err := pl.Acquire("some-schema")

	require.NoError(t, err)
	require.True(t, acquired)
}

func TestAcquire_SuccessAfterFailed(t *testing.T) {
	db := &mocks.DB{}
	db.On("Raw", "SELECT pg_try_advisory_lock(to_regnamespace($1)::integer)", "some-schema").Return(db).Twice()
	db.On("Scan", matchNullBoolScan()).Run(scanNullBool(true, false)).Return(int64(0), nil).Once()
	db.On("Scan", matchNullBoolScan()).Run(scanNullBool(true, true)).Return(int64(0), nil).Once()

	pl := Locker{db: db, retries: 3}
	acquired, err := pl.Acquire("some-schema")

	require.NoError(t, err)
	require.True(t, acquired)
}

func TestAcquire_NotAcquired(t *testing.T) {
	db := &mocks.DB{}
	db.On("Raw", "SELECT pg_try_advisory_lock(to_regnamespace($1)::integer)", "some-schema").Return(db)
	db.On("Scan", matchNullBoolScan()).Run(scanNullBool(true, false)).Return(int64(0), nil)

	pl := Locker{db: db}
	acquired, err := pl.Acquire("some-schema")

	require.NoError(t, err)
	require.False(t, acquired)
}

func TestAcquire_NullResult(t *testing.T) {
	db := &mocks.DB{}
	db.On("Raw", "SELECT pg_try_advisory_lock(to_regnamespace($1)::integer)", "some-schema").Return(db)
	db.On("Scan", matchNullBoolScan()).Run(scanNullBool(false, false)).Return(int64(0), nil)

	pl := Locker{db: db}
	acquired, err := pl.Acquire("some-schema")

	require.NoError(t, err)
	require.False(t, acquired)
}

func TestAcquire_Failed(t *testing.T) {
	db := &mocks.DB{}
	db.On("Raw", "SELECT pg_try_advisory_lock(to_regnamespace($1)::integer)", "some-schema").Return(db)
	db.On("Scan", matchNullBoolScan()).Return(int64(0), fmt.Errorf("failed to acquire lock"))

	pl := Locker{db: db}
	acquired, err := pl.Acquire("some-schema")

	require.Error(t, err)
	require.False(t, acquired)
}

func TestRelease_Success(t *testing.T) {
	db := &mocks.DB{}
	db.On("Raw", "SELECT pg_advisory_unlock(to_regnamespace($1)::integer)", "some-schema").Return(db)
	db.On("Scan", matchNullBoolScan()).Run(scanNullBool(true, true)).Return(int64(0), nil)

	pl := Locker{db: db}
	released, err := pl.Release("some-schema")

	require.NoError(t, err)
	require.True(t, released)
}

func TestRelease_NotReleased(t *testing.T) {
	db := &mocks.DB{}
	db.On("Raw", "SELECT pg_advisory_unlock(to_regnamespace($1)::integer)", "some-schema").Return(db)
	db.On("Scan", matchNullBoolScan()).Run(scanNullBool(true, false)).Return(int64(0), nil)

	pl := Locker{db: db}
	released, err := pl.Release("some-schema")

	require.NoError(t, err)
	require.False(t, released)
}

func TestRelease_Failed(t *testing.T) {
	db := &mocks.DB{}
	db.On("Raw", "SELECT pg_advisory_unlock(to_regnamespace($1)::integer)", "some-schema").Return(db)
	db.On("Scan", matchNullBoolScan()).Return(int64(0), fmt.Errorf("failed to release lock"))

	pl := Locker{db: db}
	released, err := pl.Release("some-schema")

	require.Error(t, err)
	require.False(t, released)
}

func matchNullBoolScan() interface{} {
	return mock.MatchedBy(func(dest interface{}) bool {
		_, ok := dest.(*sql.NullBool)
		return ok
	})
}

func scanNullBool(valid, val bool) func(mock.Arguments) {
	return func(args mock.Arguments) {
		ptr := args[0].(*sql.NullBool)
		ptr.Valid = valid
		ptr.Bool = val
	}
}
