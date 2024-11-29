package interfaces

import (
	"context"
	"github.com/dwiangraeni/dealls/model"
)

type IAccountService interface {
	UpgradeAccount(ctx context.Context, accountMaskID string) (string, error)
	GetListAccountNewMatchPagination(ctx context.Context, req model.PaginationRequest) (resp model.ListAccountPagination, err error)
}
