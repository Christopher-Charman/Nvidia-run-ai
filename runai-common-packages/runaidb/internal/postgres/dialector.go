package postgres

import (
	pgerrors "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/postgres/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresErrorTranslator interface {
	Translate(error) error
}

type translatingDialector struct {
	gorm.Dialector
}

func newTranslatingDialector(dsn string) translatingDialector {
	return translatingDialector{
		Dialector: postgres.New(postgres.Config{
			DSN: dsn,
		}),
	}
}

func (d translatingDialector) Translate(err error) error {
	if translator, ok := d.Dialector.(postgresErrorTranslator); ok {
		err = translator.Translate(err)
	}
	return pgerrors.Translate(err)
}
