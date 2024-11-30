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

type premiumPackageRepo struct {
	db *sqlx.DB
}

func NewPremiumPackageRepo(db *sqlx.DB) interfaces.IPremiumPackageRepo {
	return &premiumPackageRepo{
		db: db,
	}
}

func (p *premiumPackageRepo) GetListPremiumPackagePagination(ctx context.Context, req model.PaginationRequest) (output []model.PremiumPackageBaseModel, err error) {
	var (
		condition, offsetLimit, orderBy string
		inputArgs                       []interface{}
		resp                            []model.PremiumPackageBaseModel
	)

	orderBy = `ORDER BY id DESC`

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

	query := fmt.Sprintf(RepoGetListPremiumPackage, condition, orderBy, offsetLimit)
	if err = p.db.SelectContext(ctx, &resp, p.db.Rebind(query), inputArgs...); err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *premiumPackageRepo) GetPremiumPackageUserByAccountID(ctx context.Context, accountID int64) (output []model.PremiumPackageUserBaseModel, err error) {
	if err = p.db.SelectContext(ctx, &output, RepoGetPremiumPackageUserByAccountID, accountID); err != nil {
		return output, err
	}

	return output, nil
}

func (p *premiumPackageRepo) InsertPremiumPackageUser(ctx context.Context, trx *sql.Tx, req *model.PremiumPackageUserBaseModel) (err error) {
	if err = trx.QueryRowContext(ctx, InsertPremiumPackageUser, req.AccountID, req.PremiumPackageID).Scan(&req.ID, &req.PurchasedDate); err != nil {
		return err
	}

	return nil
}

func (p *premiumPackageRepo) GetPremiumPackageByPackageUID(ctx context.Context, packageUID string) (output model.PremiumPackageBaseModel, err error) {
	if err = p.db.GetContext(ctx, &output, RepoGetPremiumPackageByPackageUID, packageUID); err != nil {
		return output, err
	}

	return output, nil
}
