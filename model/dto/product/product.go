package product

type ProductCre struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required,num"`
	Description string `json:"description"`
	Status      bool   `json:"status" binding:"boolean"`
	CategoryID  string `json:"category_id" binding:"required"`
	CompanyID   string `json:"company_id" binding:"required"`
}
