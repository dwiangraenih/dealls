package model

import (
	"database/sql"
	"time"
)

type PremiumPackageBaseModel struct {
	ID          int64          `db:"id"`
	PackageUID  string         `db:"package_uid"`
	Title       string         `db:"title"`
	Description string         `db:"description"`
	Price       float64        `db:"price"`
	IsActive    bool           `db:"is_active"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	CreatedBy   string         `db:"created_by"`
	UpdatedBy   sql.NullString `db:"updated_by"`
}

type PremiumPackageUserBaseModel struct {
	ID               int64     `db:"id"`
	PremiumPackageID int64     `db:"premium_package_id"`
	AccountID        int64     `db:"account_id"`
	PurchaseDate     time.Time `db:"purchase_date"`
}

type PremiumPackageResponse struct {
	PackageUID  string    `json:"package_uid"`
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdateBy    string    `json:"updated_by"`
	IsPurchased bool      `json:"is_purchased"`
}

type ListPackagePagination struct {
	Data       []PremiumPackageResponse `json:"data"`
	LoadMore   bool                     `json:"load_more"`
	NextCursor string                   `json:"next_cursor"`
	PrevCursor string                   `json:"prev_cursor"`
	Limit      int                      `json:"limit"`
	Keywords   string                   `json:"q"`
}
