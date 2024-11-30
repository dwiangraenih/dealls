package interfaces

import (
	"context"
	"github.com/dwiangraeni/dealls/model"
)

type IAccountService interface {
	GetListAccountNewMatchPagination(ctx context.Context, req model.PaginationRequest) (resp model.ListAccountPagination, err error)
}
