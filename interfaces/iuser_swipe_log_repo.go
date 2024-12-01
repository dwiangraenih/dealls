package interfaces

import (
	"context"
	"github.com/dwiangraeni/dealls/model"
)

type IUserSwipeLogRepo interface {
	InsertUserSwipeLog(ctx context.Context, req model.UserSwipeLogBaseModel) (model.UserSwipeLogBaseModel, error)
	GetSwipeCountByAccountID(ctx context.Context, accountMaskID string) (resp model.SwipeCountBaseModel, err error)
	GetUserSwipeLogBySwiperIDAndSwpeeID(ctx context.Context, swiperID, swipeeID int64) (resp model.UserSwipeLogBaseModel, err error)
}
