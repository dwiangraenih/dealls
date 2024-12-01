package interfaces

import (
	"context"
	"database/sql"
	"github.com/dwiangraeni/dealls/model"
)

type IPremiumPackageRepo interface {
	// premium package
	GetListPremiumPackagePagination(ctx context.Context, req model.PaginationRequest) (output []model.PremiumPackageBaseModel, err error)

	// premium package user
	GetPremiumPackageUserByAccountMaskID(ctx context.Context, accountMaskID string) (output []model.PremiumPackageUserBaseModel, err error)
	InsertPremiumPackageUser(ctx context.Context, trx *sql.Tx, req *model.PremiumPackageUserBaseModel) (err error)
	GetPremiumPackageByPackageUID(ctx context.Context, packageUID string) (output model.PremiumPackageBaseModel, err error)
	GetPremiumPackageUserByTitleAndAccountID(ctx context.Context, title string, accountID int64) (output model.PremiumPackageUserBaseModel, err error)
}
