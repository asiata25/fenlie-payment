package company

type (
	CompanyCreateRequest struct {
		Name      string      `json:"name" binding:"required"`
		SecretKey string      `json:"secret_key" binding:"required"`
		User      UserCompany `json:"user" binding:"required"`
	}

	UserCompany struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"email"`
		Password string `json:"password" binding:"required,min=8,isPassword"`
		Role     string `json:"role"`
	}

	CompanyUpdateRequest struct {
		ID        string `json:"id"`
		Name      string `json:"name" binding:"required"`
		SecretKey string `json:"secret_key"`
	}

	CompanyResponse struct {
		ID        string        `json:"id"`
		Name      string        `json:"name"`
		SecretKey string        `json:"secret_key"`
		Users     []UserCompany `json:"users"`
	}
)
