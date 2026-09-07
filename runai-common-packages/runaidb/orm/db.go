package orm

import (
	"context"
	"database/sql"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm/clause"
	"gorm.io/gorm"
)

type runaiDB struct {
	gormDB *gorm.DB
}

func NewDB(gormDB *gorm.DB) DB {
	return &runaiDB{gormDB}
}

// Expr creates a raw SQL expression, mirroring gorm.Expr.
func Expr(expr string, args ...interface{}) clause.Expr {
	return clause.Expr{SQL: expr, Vars: args}
}

func (db *runaiDB) Actual() *gorm.DB {
	return db.gormDB
}

func (db *runaiDB) Close() error {
	if db.gormDB == nil {
		return nil
	}
	sqlDB, err := db.gormDB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *runaiDB) WithContext(ctx context.Context) DB {
	return &runaiDB{db.gormDB.WithContext(ctx)}
}

func (db *runaiDB) Session(config *Session) DB {
	return &runaiDB{db.gormDB.Session(config)}
}

// ---- Transactioner ----

func (db *runaiDB) Transaction(fc func(tx DB) error, opts ...*sql.TxOptions) error {
	return db.gormDB.Transaction(func(gormTx *gorm.DB) error {
		return fc(&runaiDB{gormTx})
	}, opts...)
}

func (db *runaiDB) SavePoint(name string) DB {
	return &runaiDB{db.gormDB.SavePoint(name)}
}

func (db *runaiDB) RollbackTo(name string) DB {
	return &runaiDB{db.gormDB.RollbackTo(name)}
}

func (db *runaiDB) Begin() DB {
	return &runaiDB{gormDB: db.gormDB.Begin()}
}

func (db *runaiDB) Commit() error {
	return db.gormDB.Commit().Error
}

func (db *runaiDB) Rollback() error {
	return db.gormDB.Rollback().Error
}

// ---- QueryBuilder ----

func (db *runaiDB) Model(value interface{}) DB {
	return &runaiDB{db.gormDB.Model(value)}
}

func (db *runaiDB) Clauses(conds ...clause.Expression) DB {
	return &runaiDB{db.gormDB.Clauses(conds...)}
}

func (db *runaiDB) Table(name string, args ...interface{}) DB {
	return &runaiDB{db.gormDB.Table(name, unwrapDBArgs(args)...)}
}

func (db *runaiDB) Distinct(args ...interface{}) DB {
	return &runaiDB{db.gormDB.Distinct(args...)}
}

func (db *runaiDB) Select(query interface{}, args ...interface{}) DB {
	return &runaiDB{db.gormDB.Select(query, unwrapDBArgs(args)...)}
}

func (db *runaiDB) Omit(columns ...string) DB {
	return &runaiDB{db.gormDB.Omit(columns...)}
}

func (db *runaiDB) Where(query interface{}, args ...interface{}) DB {
	return &runaiDB{db.gormDB.Where(query, unwrapDBArgs(args)...)}
}

func (db *runaiDB) Not(query interface{}, args ...interface{}) DB {
	return &runaiDB{db.gormDB.Not(unwrapQuery(query), unwrapDBArgs(args)...)}
}

func (db *runaiDB) Or(query interface{}, args ...interface{}) DB {
	return &runaiDB{db.gormDB.Or(unwrapQuery(query), unwrapDBArgs(args)...)}
}

func (db *runaiDB) Joins(query string, args ...interface{}) DB {
	return &runaiDB{db.gormDB.Joins(query, unwrapDBArgs(args)...)}
}

func (db *runaiDB) InnerJoins(query string, args ...interface{}) DB {
	return &runaiDB{db.gormDB.InnerJoins(query, unwrapDBArgs(args)...)}
}

func (db *runaiDB) Group(name string) DB {
	return &runaiDB{db.gormDB.Group(name)}
}

func (db *runaiDB) Having(query interface{}, args ...interface{}) DB {
	return &runaiDB{db.gormDB.Having(unwrapQuery(query), unwrapDBArgs(args)...)}
}

func (db *runaiDB) Order(value interface{}) DB {
	return &runaiDB{db.gormDB.Order(value)}
}

func (db *runaiDB) Limit(limit int) DB {
	return &runaiDB{db.gormDB.Limit(limit)}
}

func (db *runaiDB) Offset(offset int) DB {
	return &runaiDB{db.gormDB.Offset(offset)}
}

func (db *runaiDB) Preload(query string, args ...interface{}) DB {
	return &runaiDB{db.gormDB.Preload(query, args...)}
}

func (db *runaiDB) Assign(attrs ...interface{}) DB {
	return &runaiDB{db.gormDB.Assign(attrs...)}
}

func (db *runaiDB) Unscoped() DB {
	return &runaiDB{db.gormDB.Unscoped()}
}

func (db *runaiDB) Raw(sql string, values ...interface{}) DB {
	return &runaiDB{db.gormDB.Raw(sql, unwrapDBArgs(values)...)}
}

func (db *runaiDB) WithEnricher(enrichers ...Enricher) DB {
	scopes := make([]func(*gorm.DB) *gorm.DB, len(enrichers))
	for i, e := range enrichers {
		scopes[i] = func(tx *gorm.DB) *gorm.DB {
			return e(NewDB(tx)).Actual()
		}
	}
	return &runaiDB{db.gormDB.Scopes(scopes...)}
}

// ---- QueryReader ----

func (db *runaiDB) First(dest interface{}, conds ...interface{}) error {
	return db.gormDB.First(dest, conds...).Error
}

func (db *runaiDB) Take(dest interface{}, conds ...interface{}) error {
	return db.gormDB.Take(dest, conds...).Error
}

func (db *runaiDB) Last(dest interface{}, conds ...interface{}) error {
	return db.gormDB.Last(dest, conds...).Error
}

func (db *runaiDB) Find(dest interface{}, conds ...interface{}) error {
	return db.gormDB.Find(dest, conds...).Error
}

func (db *runaiDB) Count(count *int64) error {
	return db.gormDB.Count(count).Error
}

func (db *runaiDB) Pluck(column string, dest interface{}) error {
	return db.gormDB.Pluck(column, dest).Error
}

func (db *runaiDB) Row() *sql.Row {
	return db.gormDB.Row()
}

func (db *runaiDB) Rows() (*sql.Rows, error) {
	return db.gormDB.Rows()
}

// ---- QueryWriter ----

func (db *runaiDB) Create(value interface{}) (int64, error) {
	r := db.gormDB.Create(value)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) CreateInBatches(value interface{}, batchSize int) (int64, error) {
	r := db.gormDB.CreateInBatches(value, batchSize)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) Save(value interface{}) (int64, error) {
	r := db.gormDB.Save(value)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) FirstOrInit(dest interface{}, conds ...interface{}) (int64, error) {
	r := db.gormDB.FirstOrInit(dest, conds...)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) FirstOrCreate(dest interface{}, conds ...interface{}) (int64, error) {
	r := db.gormDB.FirstOrCreate(dest, conds...)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) Update(column string, value interface{}) (int64, error) {
	r := db.gormDB.Update(column, value)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) Updates(values interface{}) (int64, error) {
	r := db.gormDB.Updates(values)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) UpdateColumn(column string, value interface{}) (int64, error) {
	r := db.gormDB.UpdateColumn(column, value)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) UpdateColumns(values interface{}) (int64, error) {
	r := db.gormDB.UpdateColumns(values)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) Delete(value interface{}, conds ...interface{}) (int64, error) {
	r := db.gormDB.Delete(value, conds...)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) Scan(dest interface{}) (int64, error) {
	r := db.gormDB.Scan(dest)
	return r.RowsAffected, r.Error
}

func (db *runaiDB) Exec(sql string, values ...interface{}) (int64, error) {
	r := db.gormDB.Exec(sql, values...)
	return r.RowsAffected, r.Error
}

// unwrapDBArgs replaces any DB in args with its underlying *gorm.DB, since
// GORM type-asserts subquery args as *gorm.DB wherever it accepts them.
func unwrapDBArgs(args []interface{}) []interface{} {
	for i, arg := range args {
		if db, ok := arg.(DB); ok {
			args[i] = db.Actual()
		}
	}
	return args
}

func unwrapQuery(query interface{}) interface{} {
	if db, ok := query.(DB); ok {
		return db.Actual()
	}
	return query
}
