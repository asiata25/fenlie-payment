package entity

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID        string         `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	Name      string         `gorm:"column:name"`
	SecretKey string         `gorm:"column:secret_key"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Users     []User         `gorm:"foreign_key:company_id;references:id"`
}
