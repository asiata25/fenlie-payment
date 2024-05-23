package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        string         `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	CompanyID string         `gorm:"column:company_id"`
	Company   Company        `gorm:"foreignKey:CompanyID"`
	Products  []Product      `gorm:"foreignKey:CategoryID"`
}
