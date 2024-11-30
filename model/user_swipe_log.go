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
	SwiperAccountMaskID string `json:"-" valid:"required"`
	SwipeeAccountMaskID string `json:"swipee_id" valid:"required"`
	SwipeType           string `json:"swipe_type" valid:"required,in(LIKE|PASS)"`
}

type SwipeCountBaseModel struct {
	ID             int64 `db:"id"`
	AccountID      int64 `db:"account_id"`
	TotalSwipeADay int   `db:"total_swipe_a_day"`
	TotalSwipe     int   `db:"total_swipe"`
}
