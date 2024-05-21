package categoryDto

type CreateCategoryRequest struct {
	Name      string `json:"name,omitempty" gorm:"column:name;omitempty;not null"`
	CompanyID string `json:"company_id" gorm:"type:uuid"`
}

type UpdateCategoryRequest struct {
	ID        string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string `json:"name,omitempty" gorm:"column:name;omitempty;not null"`
	CompanyID string `json:"company_id" binding:"required,max=255"`
}

type CategoryResponse struct {
	ID        string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string `json:"name,omitempty" gorm:"column:name;omitempty;not null"`
	CompanyID string `json:"company_id" binding:"required,max=255"`
}
