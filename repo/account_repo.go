package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/model"
	"github.com/dwiangraeni/dealls/utils"
	"github.com/jmoiron/sqlx"
)

type user struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) interfaces.IAccountRepo {
	return &user{db: db}
}

func (u *user) FindOneAccountByAccountUserName(ctx context.Context, userName string) (output model.AccountBaseModel, err error) {
	if err = u.db.QueryRowContext(ctx, RepoFindOneAccountByAccountUserName, userName).
		Scan(&output.ID, &output.AccountMaskID, &output.Type, &output.Name, &output.UserName, &output.Password,
			&output.CreatedAt, &output.CreatedBy, &output.UpdatedAt, &output.UpdatedBy); err != nil {
		return output, err
	}
	return output, err
}

func (u *user) InsertAccount(ctx context.Context, account model.AccountBaseModel) (model.AccountBaseModel, error) {
	if err := u.db.QueryRowContext(ctx, RepoInsertAccount, account.Type, account.Name, account.UserName,
		account.Password, account.CreatedBy).Scan(&account.ID, &account.AccountMaskID); err != nil {
		return account, err
	}

	return account, nil
}

func (u *user) UpdateAccountType(ctx context.Context, trx *sql.Tx, account model.AccountBaseModel) (model.AccountBaseModel, error) {
	if _, err := trx.ExecContext(ctx, RepoUpdateAccount, account.ID, account.Type, account.Name, account.UserName, account.UpdatedBy, account.IsVerified); err != nil {
		return account, err
	}
	return account, nil
}

func (u *user) FindOneAccountByAccountMaskID(ctx context.Context, accountMaskID string) (output model.AccountBaseModel, err error) {
	if err = u.db.QueryRowContext(ctx, RepoFindOneAccountByAccountMaskID, accountMaskID).
		Scan(&output.ID, &output.AccountMaskID, &output.Type, &output.Name, &output.UserName, &output.IsVerified,
			&output.CreatedAt, &output.CreatedBy, &output.UpdatedAt, &output.UpdatedBy); err != nil {
		return output, err
	}
	return output, err
}

func (u *user) GetListAccountNewMatchPagination(ctx context.Context, req model.PaginationRequest) (output []model.AccountBaseModel, err error) {
	var (
		condition, offsetLimit, orderBy string
		inputArgs                       []interface{}
		resp                            []model.AccountBaseModel
	)

	orderBy = `ORDER BY id DESC`
	// Set condition
	if req.AccountMaskID != "" {
		condition += `AND account_mask_id != ? `
		inputArgs = append(inputArgs, req.AccountMaskID)
	}

	if req.CursorID != 0 && req.Direction == utils.DirectionNext {
		condition += `AND id < ? `
		inputArgs = append(inputArgs, req.CursorID)
	}

	if req.CursorID != 0 && req.Direction == utils.DirectionPrev {
		condition += `AND id > ? `
		inputArgs = append(inputArgs, req.CursorID)
		orderBy = `ORDER BY id ASC`
	}

	if req.Limit != 0 {
		offsetLimit = fmt.Sprintf("LIMIT %d", req.Limit)
	}

	query := fmt.Sprintf(RepoGetListAccountNewMatchPagination, condition, orderBy, offsetLimit)
	if err = u.db.SelectContext(ctx, &resp, u.db.Rebind(query), inputArgs...); err != nil {
		return nil, err
	}

	return resp, nil
}
