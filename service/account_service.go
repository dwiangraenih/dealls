package service

import (
	"context"
	"errors"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/model"
	"log"
)

type serviceAccountCtx struct {
	accountRepo interfaces.IAccountRepo
	authService interfaces.IAuthService
}

func NewAccountService(accountRepo interfaces.IAccountRepo, authService interfaces.IAuthService) interfaces.IAccountService {
	return &serviceAccountCtx{accountRepo: accountRepo,
		authService: authService}
}

func (s *serviceAccountCtx) UpgradeAccount(ctx context.Context, accountMaskID string) (string, error) {
	var (
		eventName = "serviceAccountCtx.UpgradeAccount"
		logFields = map[string]interface{}{
			"_event": eventName,
			"req": map[string]interface{}{
				"accountMaskID": accountMaskID,
			},
		}
	)

	account, err := s.accountRepo.FindOneAccountByAccountMaskID(ctx, accountMaskID)
	if err != nil {
		log.Printf("%s: failed to find account by account mask with err: %s", logFields, err.Error())
		return "", errors.New("account not found")
	}

	account.Type = model.AccountTypePremium
	_, err = s.accountRepo.UpdateAccountType(ctx, account)
	if err != nil {
		log.Printf("%s: failed to upgrade account with err: %s", logFields, err.Error())
		return "", errors.New("failed to upgrade account")
	}

	newToken, err := s.authService.RefreshToken(ctx, account)
	if err != nil {
		log.Printf("%s: failed to refresh token with err: %s", logFields, err.Error())
		return "", errors.New("failed to refresh token")
	}

	// Return the new token
	return newToken.Token, nil

}
