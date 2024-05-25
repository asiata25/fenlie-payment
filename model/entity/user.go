package entity

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID        string         `gorm:"column:id;primaryKey;default:gen_random_uuid()"`
		Name      string         `gorm:"column:name"`
		Email     string         `gorm:"column:email;unique"`
		Password  string         `gorm:"column:password"`
		Role      string         `gorm:"column:role"`
		CreatedAt time.Time      `gorm:"autoCreateTime"`
		UpdatedAt time.Time      `gorm:"autoUpdateTime"`
		DeletedAt gorm.DeletedAt `gorm:"index"`
		CompanyID string         `gorm:"column:company_id"`
		Company   *Company       `gorm:"foreignKey:CompanyID;references:ID"`
	}
)
