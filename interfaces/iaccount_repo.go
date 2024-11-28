package interfaces

import (
	"context"
	"github.com/dwiangraeni/dealls/model"
)

type IAccountRepo interface {
	FindOneAccountByAccountUserName(ctx context.Context, userName string) (model.AccountBaseModel, error)
	CreateAccount(ctx context.Context, account model.AccountBaseModel) (model.AccountBaseModel, error)
	UpdateAccountType(ctx context.Context, account model.AccountBaseModel) (model.AccountBaseModel, error)
	FindOneAccountByAccountMaskID(ctx context.Context, accountMaskID string) (output model.AccountBaseModel, err error)
}
