package interfaces

import (
	"context"
	"database/sql"
)

type ITransactionRepo interface {
	BeginTrx(ctx context.Context) (trx *sql.Tx, err error)
	CommitTrx(ctx context.Context, trx *sql.Tx) (err error)
	RollbackTrx(ctx context.Context, trx *sql.Tx) (err error)
}
