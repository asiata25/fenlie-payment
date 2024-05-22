package entity

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID        string         `gorm:"primaryKey;default:gen_random_uuid()"`
		Name      string         `json:"name"`
		Email     string         `json:"email"`
		Password  string         `json:"password,omitempty"`
		CompanyID string         `json:"company_id,omitempty"`
		Role      string         `json:"role,omitempty"`
		CreatedAt time.Time      `gorm:"autoCreateTime"`
		UpdatedAt time.Time      `gorm:"autoUpdateTime"`
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)
