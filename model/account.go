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
	IsVerified    bool           `db:"is_verified"`
	CreatedAt     time.Time      `db:"created_at"`
	CreatedBy     string         `db:"created_by"`
	UpdatedAt     time.Time      `db:"updated_at"`
	UpdatedBy     sql.NullString `db:"updated_by"`
}

type PaginationRequest struct {
	Keywords      string `json:"q"`
	Cursor        string `json:"cursor"`
	Direction     string `json:"direction" valid:"optional,in(next|prev)"`
	Limit         int    `json:"limit" valid:"required"`
	CursorID      int64  `json:"-"`
	AccountMaskID string `json:"-"`
}

type AccountResponse struct {
	AccountMaskID string `json:"account_mask_id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	UserName      string `json:"user_name"`
	IsVerified    bool   `json:"is_verified"`
}

type ListAccountPagination struct {
	Data       []AccountResponse `json:"data"`
	LoadMore   bool              `json:"load_more"`
	NextCursor string            `json:"next_cursor"`
	PrevCursor string            `json:"prev_cursor"`
	Limit      int               `json:"limit"`
	Keywords   string            `json:"q"`
}

type PremiumPackageCheckoutRequest struct {
	AccountMaskID string `json:"-"`
	PackageUID    string `json:"package_uid" valid:"required"`
}
