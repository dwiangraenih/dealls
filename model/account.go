package model

import (
	"database/sql"
	"time"
)

type AccountBaseModel struct {
	ID            int64          `db:"id"`
	AccountMaskID string         `db:"account_mask_id"`
	Type          string         `db:"type"`
	Name          string         `db:"name"`
	UserName      string         `db:"user_name"`
	Password      string         `db:"password"`
	CreatedAt     time.Time      `db:"created_at"`
	CreatedBy     string         `db:"created_by"`
	UpdatedAt     time.Time      `db:"updated_at"`
	UpdatedBy     sql.NullString `db:"updated_by"`
}
