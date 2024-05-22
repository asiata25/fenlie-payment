package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          string         `gorm:"primaryKey;default:gen_random_uuid()"`
	Name        string         `json:"name"`
	Price       int            `json:"price"`
	Description string         `json:"description"`
	Status      bool           `json:"status"`
	CategoryID  string         `json:"category_id" gorm:"type:uuid"`
	CompanyID   string         `json:"company_id" gorm:"type:uuid"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
