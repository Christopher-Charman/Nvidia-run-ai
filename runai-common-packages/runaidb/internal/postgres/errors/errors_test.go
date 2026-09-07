package pgerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func TestTranslate_Nil(t *testing.T) {
	require.Nil(t, Translate(nil))
}

func TestTranslate_ReadOnlyTransaction(t *testing.T) {
	pgErr := &pgconn.PgError{Code: ErrCodeReadOnlyTransaction}
	wrapped := fmt.Errorf("exec failed: %w", pgErr)

	require.Equal(t, ErrReadOnlyConnection, Translate(pgErr))
	require.Equal(t, ErrReadOnlyConnection, Translate(wrapped))
	require.True(t, IsReadOnlyConnection(pgErr))
	require.True(t, IsReadOnlyConnection(wrapped))
	require.True(t, errors.Is(Translate(wrapped), ErrReadOnlyConnection))
}

func TestTranslate_OtherPostgresError(t *testing.T) {
	pgErr := &pgconn.PgError{Code: "23505"}
	require.Same(t, pgErr, Translate(pgErr))
	require.False(t, IsReadOnlyConnection(pgErr))
}

func TestTranslate_GenericError(t *testing.T) {
	err := errors.New("connection refused")
	require.Same(t, err, Translate(err))
	require.False(t, IsReadOnlyConnection(err))
}
