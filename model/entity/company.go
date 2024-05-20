package entity

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID           string `gorm:"primaryKey;default:gen_random_uuid()"`
	Name         string
	Email        string
	ClientSecret string
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
