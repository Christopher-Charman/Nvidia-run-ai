package pgerrors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ErrCodeInsufficientPrivilege = "42501"
	ErrCodeUserAlreadyExists     = "42710"
	ErrCodeReadOnlyTransaction   = "25006"
)

var ErrReadOnlyConnection = errors.New("write operation not allowed on read-only database connection")

func IsUserAlreadyExistsErr(err error) bool {
	return is(err, ErrCodeUserAlreadyExists)
}

func IsInsufficientPrivilegeErr(err error) bool {
	return is(err, ErrCodeInsufficientPrivilege)
}

func IsReadOnlyConnection(err error) bool {
	return errors.Is(Translate(err), ErrReadOnlyConnection)
}

func Translate(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == ErrCodeReadOnlyTransaction {
		return ErrReadOnlyConnection
	}

	return err
}

func is(err error, code string) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == code
}
