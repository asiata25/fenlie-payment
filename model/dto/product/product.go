package product

type ProductCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Description string `json:"description"`
	Status      bool   `json:"status" binding:"required,boolean"`
	CategoryID  string `json:"category_id,omitempty"`
	CompanyID   string `json:"company_id"`
}

type ProductUpdateRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty" binding:"omitempty"`
	Price       int    `json:"price,omitempty" binding:"omitempty,number"`
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status,omitempty" binding:"omitempty,boolean"`
	CategoryID  string `json:"category_id,omitempty"`
	CompanyID   string `json:"company_id"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Category    string `json:"category"`
	Image       string `json:"image"`
}
