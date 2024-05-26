package entity

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID            string         `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	EmailCustomer string         `gorm:"column:email_customer;not null;default:null"`
	Status        string         `gorm:"column:status;default:unpaid"`
	Amount        int            `gorm:"column:amount"`
	PaymentMethod string         `gorm:"column:payment_method;default:null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	TransactionID string         `gorm:"column:transaction_id"`
	Transaction   *Transaction   `gorm:"foreignKey:TransactionID;references:ID"`
}
