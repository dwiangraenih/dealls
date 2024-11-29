package interfaces

import (
	"context"
	"github.com/dwiangraeni/dealls/model"
)

type IUserSwipeLogService interface {
	ProcessUserSwipe(ctx context.Context, req model.UserSwipeRequest) error
}
