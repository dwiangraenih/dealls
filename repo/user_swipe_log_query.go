package repo

var (
	// user_swipe_log
	RepoInsertUserSwipeLog = `
	INSERT INTO user_swipe_log (swiper_id, swipee_id, swipe_type)
		VALUES ($1, $2, $3) RETURNING id;`
	RepoGetUserSwipeLogBySwiperIDAndSwpeeID = `
	SELECT id, swiper_id, swipee_id, swipe_type, created_at
		FROM user_swipe_log
		WHERE swiper_id = $1 AND swipee_id = $2 AND DATE(created_at) = (CURRENT_TIMESTAMP)::DATE
	LIMIT 1;`

	// swipe_count
	RepoGetSwipeCountByAccountMaskID = `
	SELECT account_id, total_swipe_a_day, total_swipe
		FROM swipe_count INNER JOIN account ON account.id = swipe_count.account_id
		WHERE account.account_mask_id = $1 AND DATE(last_updated_at) = (CURRENT_TIMESTAMP)::DATE;`
)
