package interfaces

import "context"

type IAccountService interface {
	UpgradeAccount(ctx context.Context, accountMaskID string) (string, error)
}
