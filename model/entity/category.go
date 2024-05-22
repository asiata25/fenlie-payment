package entity

import (
	"time"

	"gorm.io/gorm"
)

// type Category struct {
// 	ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
// 	Name      string         `json:"name,omitempty" gorm:"column:name;omitempty;not null"`
// 	CompanyID string         `json:"company_id" gorm:"type:uuid"`
// 	CreatedAt time.Time      `gorm:"autoCreateTime"`
// 	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

type Category struct {
	ID        string         `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string         `gorm:"type:varchar(255);not null"`
	CompanyID string         `gorm:"type:uuid;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
