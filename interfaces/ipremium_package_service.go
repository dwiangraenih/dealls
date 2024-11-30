package interfaces

import (
	"context"
	"github.com/dwiangraeni/dealls/model"
)

type IPremiumPackageService interface {
	GetListPremiumPackagePagination(ctx context.Context, req model.PaginationRequest) (output model.ListPackagePagination, err error)
	PremiumPackageCheckout(ctx context.Context, req model.PremiumPackageCheckoutRequest) error
}
