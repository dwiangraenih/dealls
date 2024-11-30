package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/utils"
	"log"
	"strings"
)

type servicePremiumPackageCtx struct {
	accountRepo        interfaces.IAccountRepo
	premiumPackageRepo interfaces.IPremiumPackageRepo
	hashCursor         utils.HashInterface
	transactionRepo    interfaces.ITransactionRepo
}

func NewPremiumPackageService(accountRepo interfaces.IAccountRepo,
	premiumPackageRepo interfaces.IPremiumPackageRepo,
	transactionRepo interfaces.ITransactionRepo) interfaces.IPremiumPackageService {
	return &servicePremiumPackageCtx{
		accountRepo:        accountRepo,
		premiumPackageRepo: premiumPackageRepo,
		hashCursor:         utils.InitHash(utils.ConstCursorHashSalt, utils.ConstHashLength),
		transactionRepo:    transactionRepo,
	}
}

func (s *servicePremiumPackageCtx) GetListPremiumPackagePagination(ctx context.Context, req model.PaginationRequest) (resp model.ListPackagePagination, err error) {
	var (
		eventName = "servicePremiumPackageCtx.GetListPremiumPackagePagination"
		logFields = map[string]interface{}{
			"_event": eventName,
			"req":    req,
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

	// get list premium package
	req.Limit = req.Limit + 1
	packageList, err := s.premiumPackageRepo.GetListPremiumPackagePagination(ctx, req)
	if err != nil {
		log.Printf("%s: error get list premium package: %v", logFields, err)
		return resp, err
	}

	if packageList == nil {
		return resp, nil
	}

	if len(packageList) > actualLimit {
		loadMore = true
		packageList = packageList[:actualLimit]
	}

	// get accountID
	account, err := s.accountRepo.FindOneAccountByAccountMaskID(ctx, req.AccountMaskID)
	if err != nil {
		log.Printf("%s: error get accountID: %v", logFields, err)
		return resp, utils.ErrInternal
	}

	// get user premium package
	userPremiumPackage, err := s.premiumPackageRepo.GetPremiumPackageUserByAccountID(ctx, account.ID)
	if err != nil {
		log.Printf("%s: error get user premium package: %v", logFields, err)
		return resp, utils.ErrInternal
	}
	listUserPremiumPackage := make([]int, len(userPremiumPackage))
	for i, v := range userPremiumPackage {
		listUserPremiumPackage[i] = int(v.PremiumPackageID)
	}

	premiumPackageList := make([]model.PremiumPackageResponse, len(packageList))
	dataCursor = make([]int, len(packageList))
	for i, v := range packageList {
		dataCursor[i] = int(v.ID)

		premiumPackageList[i] = model.PremiumPackageResponse{
			PackageUID:  v.PackageUID,
			Title:       v.Title,
			Price:       v.Price,
			IsActive:    v.IsActive,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt.Time,
			CreatedBy:   v.CreatedBy,
			UpdateBy:    v.UpdatedBy.String,
			IsPurchased: utils.IsIntInSlice(listUserPremiumPackage, int(v.ID)),
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

	resp.Data = premiumPackageList
	resp.LoadMore = loadMore
	resp.NextCursor = nextCursor
	resp.PrevCursor = prevCursor
	resp.Limit = actualLimit

	return resp, nil
}

func (s *servicePremiumPackageCtx) PremiumPackageCheckout(ctx context.Context, req model.PremiumPackageCheckoutRequest) error {
	var (
		eventName = "servicePremiumPackageCtx.PremiumPackageCheckout"
		logFields = map[string]interface{}{
			"_event": eventName,
			"req":    req,
		}
	)

	account, err := s.accountRepo.FindOneAccountByAccountMaskID(ctx, req.AccountMaskID)
	if err != nil {
		log.Printf("%s: failed to find account by account mask with err: %s", logFields, err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return utils.ErrDataNotFound
		}
		return utils.ErrInternal
	}

	// get premium package
	premiumPackage, err := s.premiumPackageRepo.GetPremiumPackageByPackageUID(ctx, req.PackageUID)
	if err != nil {
		log.Printf("%s: failed to get premium package by package uid with err: %s", logFields, err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return utils.ErrDataNotFound
		}

		return utils.ErrInternal
	}

	// begin transaction
	tx, err := s.transactionRepo.BeginTrx(ctx)
	if err != nil {
		log.Printf("%s: error begin transaction: %v", logFields, err)
		return utils.ErrInternal
	}

	// insert user premium package
	userPremiumPackage := model.PremiumPackageUserBaseModel{
		PremiumPackageID: premiumPackage.ID,
		AccountID:        account.ID,
	}

	if err = s.premiumPackageRepo.InsertPremiumPackageUser(ctx, tx, &userPremiumPackage); err != nil {
		s.transactionRepo.RollbackTrx(ctx, tx)
		log.Printf("%s: failed to insert user premium package with err: %s", logFields, err.Error())
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("package already purchased")
		}

		return utils.ErrInternal
	}

	if account.Type == model.AccountTypeFree {
		account.Type = model.AccountTypePremium
	}

	if premiumPackage.Title == model.PremiumPackageVerified {
		account.IsVerified = true
	}

	_, err = s.accountRepo.UpdateAccountType(ctx, tx, account)
	if err != nil {
		s.transactionRepo.RollbackTrx(ctx, tx)
		log.Printf("%s: failed to upgrade account with err: %s", logFields, err.Error())
		return utils.ErrInternal
	}

	// commit transaction
	if err = s.transactionRepo.CommitTrx(ctx, tx); err != nil {
		log.Printf("%s: error commit transaction: %v", logFields, err)
		return utils.ErrInternal
	}

	return nil

}
