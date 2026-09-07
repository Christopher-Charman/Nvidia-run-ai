package runaihttp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm"
)

const ErrForeignKey = "23503"
const ErrDuplicateKey = "23505"
const ErrInvalidSyntax = "22P02"
const ErrTextNotFound = "record not found" // first() in some cases returns plain error
const ErrTriggerFailed = "P0001"

func LoadErrorDetails(code *int32, message *string, details *string, goError error) {
	if goError == nil {
		return
	}

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(goError, orm.ErrRecordNotFound):
		*code = http.StatusNotFound
		*details = "requested resource does not exist"
	case errors.Is(goError, orm.ErrDuplicatedKey):
		*code = http.StatusConflict
		*details = orm.ErrDuplicatedKey.Error()
	case errors.Is(goError, orm.ErrForeignKeyViolated):
		*code = http.StatusPreconditionFailed
		*details = orm.ErrForeignKeyViolated.Error()
	case errors.Is(goError, orm.ErrCheckConstraintViolated):
		*code = http.StatusConflict
		*details = orm.ErrCheckConstraintViolated.Error()
	case errors.Is(goError, orm.ErrReadOnlyConnection):
		*code = http.StatusServiceUnavailable
		*details = orm.ErrReadOnlyConnection.Error()
	case errors.Is(goError, orm.ErrUndefinedTable):
		*code = http.StatusInternalServerError
		*details = orm.ErrUndefinedTable.Error()
	case errors.As(goError, &pgErr):
		//   Alter the code to more specific value instead of Service Unavailable based on
		//   identifying the SQL error codes.
		switch pgErr.Code {
		case ErrDuplicateKey:
			*code = http.StatusConflict
			*details = pgErr.Detail
		case ErrInvalidSyntax:
			*code = http.StatusBadRequest
			*details = pgErr.Message
		case ErrForeignKey:
			*code = http.StatusPreconditionFailed
			*details = fmt.Sprintf("The request invalidates %s dependency", pgErr.ConstraintName)
		case ErrTriggerFailed:
			*code = http.StatusConflict
			*message = pgErr.Message
		}
	case goError.Error() == ErrTextNotFound:
		*code = http.StatusNotFound
		*details = "requested resource does not exist"
	default:
		*details = goError.Error()
	}
}
