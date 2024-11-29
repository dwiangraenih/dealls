package model

type UserSwipeLogBaseModel struct {
	ID            string `db:"id"`
	SwiperID      int64  `db:"swiper_id"`
	SwipeeID      int64  `db:"swipee_id"`
	SwipeType     string `db:"swipe_type"`
	CreatedAt     string `db:"created_at"`
	ConvertedDate string `db:"converted_date"`
}

type UserSwipeRequest struct {
	SwiperAccountMaskID string `json:"-" validate:"required"`
	SwipeeAccountMaskID string `json:"swipee_id" validate:"required"`
	SwipeType           string `json:"swipe_type" validate:"required, in(LIKE, DISLIKE)"`
}

type SwipeCountBaseModel struct {
	ID             string `db:"id"`
	AccountID      string `db:"account_id"`
	TotalSwipeADay int    `db:"total_swipe_a_day"`
	TotalSwipe     int    `db:"total_swipe"`
}
