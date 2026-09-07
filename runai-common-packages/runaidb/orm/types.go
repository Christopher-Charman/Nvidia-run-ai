package orm

import (
	"context"
	"database/sql"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm/clause"
	"gorm.io/gorm"
)

// GORM core types.
type (
	Session = gorm.Session
)

// DB is the main database interface composing all operation groups.
type DB interface {
	Actual() *gorm.DB

	WithContext(ctx context.Context) DB
	Session(config *Session) DB
	Close() error

	Transactioner
	QueryBuilder
	QueryReader
	QueryWriter
}

type Transactioner interface {
	Transaction(fc func(tx DB) error, opts ...*sql.TxOptions) error
	SavePoint(name string) DB
	RollbackTo(name string) DB
}

type QueryBuilder interface {
	Model(value interface{}) DB
	Table(name string, args ...interface{}) DB
	Clauses(conds ...clause.Expression) DB
	Where(query interface{}, args ...interface{}) DB
	Not(query interface{}, args ...interface{}) DB
	Or(query interface{}, args ...interface{}) DB
	Select(query interface{}, args ...interface{}) DB
	Omit(columns ...string) DB
	Joins(query string, args ...interface{}) DB
	InnerJoins(query string, args ...interface{}) DB
	Preload(query string, args ...interface{}) DB
	Group(name string) DB
	Having(query interface{}, args ...interface{}) DB
	Order(value interface{}) DB
	Limit(limit int) DB
	Offset(offset int) DB
	Distinct(args ...interface{}) DB
	Assign(attrs ...interface{}) DB
	Unscoped() DB
	Raw(sql string, values ...interface{}) DB
}

type QueryReader interface {
	First(dest interface{}, conds ...interface{}) error
	Take(dest interface{}, conds ...interface{}) error
	Last(dest interface{}, conds ...interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	Count(count *int64) error
	Row() *sql.Row
	Rows() (*sql.Rows, error)
}

type QueryWriter interface {
	Create(value interface{}) (int64, error)
	CreateInBatches(value interface{}, batchSize int) (int64, error)
	Save(value interface{}) (int64, error)
	FirstOrInit(dest interface{}, conds ...interface{}) (int64, error)
	FirstOrCreate(dest interface{}, conds ...interface{}) (int64, error)
	Update(column string, value interface{}) (int64, error)
	Updates(values interface{}) (int64, error)
	UpdateColumn(column string, value interface{}) (int64, error)
	UpdateColumns(values interface{}) (int64, error)
	Delete(value interface{}, conds ...interface{}) (int64, error)
	Scan(dest interface{}) (int64, error)
	Exec(sql string, values ...interface{}) (int64, error)
}
