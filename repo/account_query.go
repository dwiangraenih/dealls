package repo

var (
	RepoFindOneAccountByAccountUserName = `
	SELECT id, account_mask_id, type, name, user_name, password, created_at, created_by, updated_at, updated_by
		FROM account where user_name = $1;`

	RepoInsertAccount = `
	INSERT INTO account (type, name, user_name, password, created_by)
	    		VALUES ($1, $2, $3, $4, $5) RETURNING id,account_mask_id;`

	RepoUpdateAccountType = `
	UPDATE account SET type = $1, updated_by = $2, updated_at = now() WHERE id = $3;`

	RepoFindOneAccountByAccountMaskID = `
	SELECT id, account_mask_id, type, name, user_name, password, created_at, created_by, updated_at, updated_by
		FROM account where account_mask_id = $1;`

	RepoGetListAccountNewMatchPagination = `
	SELECT id, account_mask_id, type, name, user_name, created_at, created_by, updated_at, updated_by
		FROM account WHERE id NOT IN (
		SELECT swipee_id FROM user_swipe_log WHERE user_swipe_log.swipee_id=account.id AND DATE(created_at) = (CURRENT_TIMESTAMP)::DATE)
	%s %s %s;`
)
