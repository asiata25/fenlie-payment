package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          string         `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	Name        string         `gorm:"column:name"`
	Price       int            `gorm:"column:price"`
	Description sql.NullString `gorm:"column:description"`
	Status      bool           `gorm:"column:status"`
	CategoryID  sql.NullString `gorm:"type:uuid"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
	CompanyID   string         `gorm:"column:company_id"`
	Company     Company        `gorm:"foreignKey:CompanyID"`
	Category    Category
}
