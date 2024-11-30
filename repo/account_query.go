package repo

var (
	RepoFindOneAccountByAccountUserName = `
	SELECT id, account_mask_id, type, name, user_name, password, created_at, created_by, updated_at, updated_by
		FROM account where user_name = $1;`

	RepoInsertAccount = `
	INSERT INTO account (type, name, user_name, password, created_by)
	    		VALUES ($1, $2, $3, $4, $5) RETURNING id,account_mask_id;`

	RepoUpdateAccount = `
	UPDATE account SET type = $2, name = $3, user_name = $4, updated_by = $5,  is_verified = $6, updated_at = now()
	WHERE id = $1 ;`

	RepoFindOneAccountByAccountMaskID = `
	SELECT id, account_mask_id, type, name, user_name, is_verified, created_at, created_by, updated_at, updated_by
		FROM account where account_mask_id = $1;`

	RepoGetListAccountNewMatchPagination = `
	SELECT id, account_mask_id, type, name, user_name, is_verified, created_at, created_by, updated_at, updated_by
		FROM account WHERE id NOT IN (
		SELECT swipee_id FROM user_swipe_log WHERE user_swipe_log.swipee_id=account.id AND DATE(created_at) = (CURRENT_TIMESTAMP)::DATE)
	%s %s %s;`
)
