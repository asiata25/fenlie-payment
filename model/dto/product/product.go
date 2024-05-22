package product

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required,num"`
	Description string `json:"description"`
	Status      bool   `json:"status" binding:"required,boolean"`
	CategoryID  string `json:"category_id,omitempty"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Category    string `json:"category"`
}
