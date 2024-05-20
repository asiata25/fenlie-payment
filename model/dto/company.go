package dto

type (
	CompanyCreateRequest struct {
		Name      string `json:"name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		SecretKey string `json:"secret_key" binding:"required"`
	}

	CompanyUpdateRequest struct {
		ID        string `json:"id"`
		Name      string `json:"name" binding:"required"`
		SecretKey string `json:"secret_key"`
	}

	CompanyResponse struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		SecretKey string `json:"secret_key"`
	}
)
