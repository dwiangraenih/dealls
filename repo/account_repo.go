package repo

import (
	"context"
	"database/sql"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/model"
)

type user struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) interfaces.IAccountRepo {
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

func (u *user) CreateAccount(ctx context.Context, account model.AccountBaseModel) (model.AccountBaseModel, error) {
	if err := u.db.QueryRowContext(ctx, RepoInsertAccount, account.Type, account.Name, account.UserName,
		account.Password, account.CreatedBy).Scan(&account.ID, &account.AccountMaskID); err != nil {
		return account, err
	}

	return account, nil
}

func (u *user) UpdateAccountType(ctx context.Context, account model.AccountBaseModel) (model.AccountBaseModel, error) {
	if _, err := u.db.ExecContext(ctx, RepoUpdateAccountType, account.Type, account.UpdatedBy, account.ID); err != nil {
		return account, err
	}
	return account, nil
}

func (u *user) FindOneAccountByAccountMaskID(ctx context.Context, accountMaskID string) (output model.AccountBaseModel, err error) {
	if err = u.db.QueryRowContext(ctx, RepoFindOneAccountByAccountMaskID, accountMaskID).
		Scan(&output.ID, &output.AccountMaskID, &output.Type, &output.Name, &output.UserName, &output.Password,
			&output.CreatedAt, &output.CreatedBy, &output.UpdatedAt, &output.UpdatedBy); err != nil {
		return output, err
	}
	return output, err
}
