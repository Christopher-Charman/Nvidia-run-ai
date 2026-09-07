package postgres

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	pgerrors "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/postgres/errors"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm"
)

type RWSessionCreator struct {
	db orm.DB
}

func NewRWSessionCreator(db orm.DB) *RWSessionCreator {
	return &RWSessionCreator{
		db: db,
	}
}

func (s *RWSessionCreator) Create(user, password, schema string) error {
	sanitizedUser := pgx.Identifier{user}.Sanitize()
	sanitizedSchema := pgx.Identifier{schema}.Sanitize()
	sanitizedPassword := escapeLiteral(password)

	_, err := s.db.Exec(fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s';", sanitizedUser, sanitizedPassword))
	if err != nil {
		if !pgerrors.IsUserAlreadyExistsErr(err) {
			return err
		}
		// User already exists, so we update the password in case it changed since the user was first created
		_, err = s.db.Exec(fmt.Sprintf("ALTER USER %s WITH LOGIN PASSWORD '%s';", sanitizedUser, sanitizedPassword))
		if err != nil {
			return err
		}
	}

	_, err = s.db.Exec(fmt.Sprintf("GRANT USAGE ON SCHEMA %s TO %s;", sanitizedSchema, sanitizedUser))
	if err != nil {
		return err
	}

	_, err = s.db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA %s TO %s;", sanitizedSchema, sanitizedUser))
	if err != nil {
		return err
	}

	_, err = s.db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA %s TO %s;", sanitizedSchema, sanitizedUser))
	if err != nil {
		return err
	}

	_, err = s.db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA %s TO %s;", sanitizedSchema, sanitizedUser))
	if err != nil {
		return err
	}

	return nil
}

func escapeLiteral(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
