package runaihttp

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
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
	switch myType := goError.(type) {
	//   Alter the code to more specific value instead of Service Unavailable based on
	//   identifying the SQL error codes.
	case *pgconn.PgError:
		switch myType.Code {
		case ErrDuplicateKey:
			*code = http.StatusConflict
			*details = myType.Detail
		case ErrInvalidSyntax:
			*code = http.StatusBadRequest
			*details = myType.Message
		case ErrForeignKey:
			*code = http.StatusPreconditionFailed
			*details = fmt.Sprintf("The request invalidates %s dependency", myType.ConstraintName)
		case ErrTriggerFailed:
			*code = http.StatusConflict
			*message = myType.Message
		}
		return // avoid setting details with postgres error, as it exposes internal information
	case error:
		switch myType.Error() {
		case ErrTextNotFound:
			*code = http.StatusNotFound
			*details = "requested resource does not exist"
			return
		}
	}
	*details = goError.Error()
}
