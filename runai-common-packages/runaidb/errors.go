package runaidb

import pgerrors "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/postgres/errors"

var ErrReadOnlyConnection = pgerrors.ErrReadOnlyConnection


func IsReadOnlyConnection(err error) bool {
	return pgerrors.IsReadOnlyConnection(err)
}
