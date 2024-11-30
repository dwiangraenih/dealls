package repo

import (
	"context"
	"database/sql"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/jmoiron/sqlx"
)

type transactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) interfaces.ITransactionRepo {
	return &transactionRepo{db: db}
}

func (t *transactionRepo) BeginTrx(ctx context.Context) (trx *sql.Tx, err error) {
	trx, err = t.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return trx, err
	}

	return trx, nil
}

func (t *transactionRepo) CommitTrx(ctx context.Context, trx *sql.Tx) error {
	return trx.Commit()
}

func (t *transactionRepo) RollbackTrx(ctx context.Context, trx *sql.Tx) error {
	return trx.Rollback()
}
