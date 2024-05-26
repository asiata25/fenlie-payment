package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                 string              `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	Status             string              `gorm:"column:status;not null;default:unpaid"`
	Total              int                 `gorm:"column:total;not null"`
	OrderDate          time.Time           `gorm:"autoCreateTime"`
	UpdatedAt          time.Time           `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt      `gorm:"index"`
	CompanyID          string              `gorm:"column:company_id"`
	Company            Company             `gorm:"foreignKey:CompanyID"`
	UserId             string              `gorm:"not null"`
	DetailTransactions []DetailTransaction `gorm:"foreignKey:TransactionID;references:ID"`
	Invoices           []Invoice           `gorm:"foreignKey:TransactionID;references:ID"`
}
