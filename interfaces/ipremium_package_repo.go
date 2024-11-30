package interfaces

import (
	"context"
	"github.com/dwiangraeni/dealls/model"
)

type IPremiumPackageRepo interface {
	// premium package
	GetListPremiumPackagePagination(ctx context.Context, req model.PaginationRequest) (output []model.PremiumPackageBaseModel, err error)

	// premium package user
	GetPremiumPackageUserByAccountID(ctx context.Context, accountID int64) (output []model.PremiumPackageUserBaseModel, err error)
}
