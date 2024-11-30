package service

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/utils"
	"log"
)

type serviceAccountCtx struct {
	accountRepo interfaces.IAccountRepo
	authService interfaces.IAuthService
	hashCursor  utils.HashInterface
}

func NewAccountService(accountRepo interfaces.IAccountRepo, authService interfaces.IAuthService) interfaces.IAccountService {
	return &serviceAccountCtx{accountRepo: accountRepo,
		authService: authService,
		hashCursor:  utils.InitHash(utils.ConstCursorHashSalt, utils.ConstHashLength)}
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
		return "", utils.ErrDataNotFound
	}

	account.Type = model.AccountTypePremium
	_, err = s.accountRepo.UpdateAccountType(ctx, account)
	if err != nil {
		log.Printf("%s: failed to upgrade account with err: %s", logFields, err.Error())
		return "", utils.ErrInternal
	}

	newToken, err := s.authService.RefreshToken(ctx, account)
	if err != nil {
		log.Printf("%s: failed to refresh token with err: %s", logFields, err.Error())
		return "", utils.ErrInternal
	}

	// Return the new token
	return newToken.Token, nil

}

func (s *serviceAccountCtx) GetListAccountNewMatchPagination(ctx context.Context, req model.PaginationRequest) (resp model.ListAccountPagination, err error) {
	var (
		eventName = "serviceAccountCtx.GetListAccountNewMatchPagination"
		logFields = map[string]interface{}{
			"_event": eventName,
			"req": map[string]interface{}{
				"req": req,
			},
		}
		actualLimit            = req.Limit
		loadMore               bool
		dataCursor             []int
		prevCursor, nextCursor string
	)

	// validate req
	if _, err = govalidator.ValidateStruct(req); err != nil {
		log.Printf("%s: error validate request: %v", logFields, err)
		return resp, err
	}

	if req.Cursor != "" {
		req.CursorID = s.hashCursor.DecodePublicID(req.Cursor)
	}

	// get list order
	req.Limit = req.Limit + 1
	accounts, err := s.accountRepo.GetListAccountNewMatchPagination(ctx, req)
	if err != nil {
		log.Printf("%s: failed to get list account with err: %s", logFields, err.Error())
		return resp, utils.ErrInternal
	}

	if accounts == nil {
		return resp, nil
	}

	accountList := make([]model.AccountResponse, len(accounts))
	dataCursor = make([]int, len(accounts))

	for i, account := range accounts {
		dataCursor[i] = int(account.ID)

		accountList[i] = model.AccountResponse{
			AccountMaskID: account.AccountMaskID,
			Type:          account.Type,
			Name:          account.Name,
			UserName:      account.UserName,
		}
	}

	prevCursorID, nextCursorID := utils.GetPaginationCursor(dataCursor, req.Direction == utils.DirectionPrev)
	nextCursor = s.hashCursor.EncodePublicID(nextCursorID)
	prevCursor = s.hashCursor.EncodePublicID(prevCursorID)
	if !loadMore && req.Direction != utils.DirectionPrev {
		nextCursor = ""
	}

	if req.CursorID == 0 || (!loadMore && req.Direction == utils.DirectionPrev) {
		prevCursor = ""
	}

	resp.Data = accountList
	resp.LoadMore = loadMore
	resp.NextCursor = nextCursor
	resp.PrevCursor = prevCursor
	resp.Limit = actualLimit

	return resp, nil
}
