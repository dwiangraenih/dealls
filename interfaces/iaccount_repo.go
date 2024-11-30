package interfaces

import (
	"context"
	"database/sql"
	"github.com/dwiangraeni/dealls/model"
)

type IAccountRepo interface {
	FindOneAccountByAccountUserName(ctx context.Context, userName string) (model.AccountBaseModel, error)
	InsertAccount(ctx context.Context, account model.AccountBaseModel) (model.AccountBaseModel, error)
	UpdateAccountType(ctx context.Context, trx *sql.Tx, account model.AccountBaseModel) (model.AccountBaseModel, error)
	FindOneAccountByAccountMaskID(ctx context.Context, accountMaskID string) (output model.AccountBaseModel, err error)
	GetListAccountNewMatchPagination(ctx context.Context, req model.PaginationRequest) (output []model.AccountBaseModel, err error)
}
