package repo

import (
	"context"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/model"
	"github.com/jmoiron/sqlx"
)

type userSwipeLog struct {
	db *sqlx.DB
}

func NewUserSwipeLogRepo(db *sqlx.DB) interfaces.IUserSwipeLogRepo {
	return &userSwipeLog{db: db}
}

func (u *userSwipeLog) InsertUserSwipeLog(ctx context.Context, req model.UserSwipeLogBaseModel) (model.UserSwipeLogBaseModel, error) {
	if _, err := u.db.ExecContext(ctx, RepoInsertUserSwipeLog, req.SwiperID, req.SwipeeID, req.SwipeType); err != nil {
		return req, err
	}
	return req, nil
}

func (u *userSwipeLog) GetSwipeCountByAccountID(ctx context.Context, accountMaskID string) (resp model.SwipeCountBaseModel, err error) {
	if err = u.db.QueryRowContext(ctx, RepoGetSwipeCountByAccountMaskID, accountMaskID).
		Scan(&resp.AccountID, &resp.TotalSwipeADay, &resp.TotalSwipe); err != nil {
		return resp, err
	}
	return resp, err
}

func (u *userSwipeLog) GetUserSwipeLogBySwiperIDAndSwpeeID(ctx context.Context, swiperID, swipeeID int64) (resp model.UserSwipeLogBaseModel, err error) {
	if err = u.db.QueryRowContext(ctx, RepoGetUserSwipeLogBySwiperIDAndSwpeeID, swiperID, swipeeID).
		Scan(&resp.ID, &resp.SwiperID, &resp.SwipeeID, &resp.SwipeType, &resp.CreatedAt); err != nil {
		return resp, err
	}
	return resp, err
}
