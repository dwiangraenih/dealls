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
)

type userSwipeLogCtx struct {
	userSwipeLogRepo   interfaces.IUserSwipeLogRepo
	accountRepo        interfaces.IAccountRepo
	premiumPackageRepo interfaces.IPremiumPackageRepo
	maxSwipeADay       int
}

func NewUserSwipeLogService(userSwipeLogRepo interfaces.IUserSwipeLogRepo,
	accountRepo interfaces.IAccountRepo,
	premiumPackageRepo interfaces.IPremiumPackageRepo,
	maxSwipeADay int) interfaces.IUserSwipeLogService {
	return &userSwipeLogCtx{userSwipeLogRepo: userSwipeLogRepo,
		accountRepo:        accountRepo,
		premiumPackageRepo: premiumPackageRepo,
		maxSwipeADay:       maxSwipeADay,
	}
}

func (u *userSwipeLogCtx) ProcessUserSwipe(ctx context.Context, req model.UserSwipeRequest) error {
	var (
		eventName = "userSwipeLogCtx.ProcessUserSwipe"
		logFields = map[string]interface{}{
			"_event": eventName,
			"rew":    req,
		}
	)

	// validate request
	if _, err := govalidator.ValidateStruct(req); err != nil {
		log.Printf("%s: error validate request: %v", logFields, err)
		return err
	}

	// validate total last swipe
	// get swipe count by account mask id
	swipeCount, err := u.userSwipeLogRepo.GetSwipeCountByAccountID(ctx, req.SwiperAccountMaskID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("%s: error get swipe count by account mask id: %v", logFields, err)
		return utils.ErrInternal
	}

	// get account by account mask id
	swiperAccount, err := u.accountRepo.FindOneAccountByAccountMaskID(ctx, req.SwiperAccountMaskID)
	if err != nil {
		log.Printf("%s: error get account by account mask id: %v", logFields, err)
		return utils.ErrInternal
	}

	// get premium package user swipe limit
	premiumPackageUser, err := u.premiumPackageRepo.GetPremiumPackageUserByTitleAndAccountID(ctx, model.PremiumPackageSwipe, swiperAccount.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("%s: error get premium package user by account id: %v", logFields, err)
		return utils.ErrInternal
	}

	if swipeCount.TotalSwipeADay >= u.maxSwipeADay && premiumPackageUser.ID == 0 {
		log.Printf("%s: total swipe a day is already reach the limit", logFields)
		return errors.New("total swipe a day is already reach the limit, upgrade your account to get more swipe")
	}

	// get account by account mask id
	swipeeAccount, err := u.accountRepo.FindOneAccountByAccountMaskID(ctx, req.SwipeeAccountMaskID)
	if err != nil {
		log.Printf("%s: error get account by account mask id: %v", logFields, err)
		return utils.ErrInternal
	}

	// validate swipee user
	swipeLog, err := u.userSwipeLogRepo.GetUserSwipeLogBySwiperIDAndSwpeeID(ctx, swiperAccount.ID, swipeeAccount.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("%s: error get user swipe log by swiper id and swipee id: %v", logFields, err)
		return utils.ErrInternal
	}

	if swipeLog.ID != 0 {
		log.Printf("%s: user already swipe this user", logFields)
		return errors.New("user already swipe this user")
	}

	// insert user swipe log
	userSwipeLog := model.UserSwipeLogBaseModel{
		SwiperID:  swiperAccount.ID,
		SwipeeID:  swipeeAccount.ID,
		SwipeType: req.SwipeType,
	}

	if _, err = u.userSwipeLogRepo.InsertUserSwipeLog(ctx, userSwipeLog); err != nil {
		log.Printf("%s: error insert user swipe log: %v", logFields, err)
		return utils.ErrInternal
	}

	return nil

}
