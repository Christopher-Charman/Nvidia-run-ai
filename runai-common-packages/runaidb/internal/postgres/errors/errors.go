package pgerrors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ErrCodeInsufficientPrivilege = "42501"
	ErrCodeUserAlreadyExists     = "42710"
)

func IsUserAlreadyExistsErr(err error) bool {
	return is(err, ErrCodeUserAlreadyExists)
}

func IsInsufficientPrivilegeErr(err error) bool {
	return is(err, ErrCodeInsufficientPrivilege)
}

func is(err error, code string) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == code
}
