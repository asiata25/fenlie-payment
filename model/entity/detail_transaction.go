package entity

import (
	"time"

	"gorm.io/gorm"
)

type DetailTransaction struct {
	ID            string         `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	Quantity      int            `gorm:"column:quantity"`
	Amount        int            `gorm:"column:amount"`
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`
	CompanyID     string         `gorm:"column:company_id"`
	Company       *Company       `gorm:"foreignKey:CompanyID;references:ID"`
	TransactionID string         `gorm:"column:transaction_id"`
	Transaction   *Transaction   `gorm:"foreignKey:TransactionID;references:ID"`
	ProductID     string         `gorm:"column:product_id"`
	Product       *Product       `gorm:"foreignKey:ProductID;references:ID"`
}
