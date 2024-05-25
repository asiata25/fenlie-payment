package category

type CategoryRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty" binding:"required"`
	CompanyID string `json:"companyid"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}
