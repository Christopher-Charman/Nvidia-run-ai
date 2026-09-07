package orm

import "gorm.io/gorm"

var (
	ErrRecordNotFound                = gorm.ErrRecordNotFound
	ErrInvalidTransaction            = gorm.ErrInvalidTransaction
	ErrMissingWhereClause            = gorm.ErrMissingWhereClause
	ErrUnsupportedRelation           = gorm.ErrUnsupportedRelation
	ErrPrimaryKeyRequired            = gorm.ErrPrimaryKeyRequired
	ErrModelValueRequired            = gorm.ErrModelValueRequired
	ErrModelAccessibleFieldsRequired = gorm.ErrModelAccessibleFieldsRequired
	ErrDuplicatedKey                 = gorm.ErrDuplicatedKey
	ErrForeignKeyViolated            = gorm.ErrForeignKeyViolated
	ErrCheckConstraintViolated       = gorm.ErrCheckConstraintViolated
)
