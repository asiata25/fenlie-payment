package entity

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID           string         `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	Name         string         `gorm:"column:name"`
	Email        string         `gorm:"column:email;unique"`
	ClientSecret string         `gorm:"column:client_secret"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
