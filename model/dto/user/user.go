package user

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	UserResponse struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		CompanyID string `json:"company_id"`
		Role      string `json:"role"`
	}

	CreateUserRequest struct {
		Name      string `json:"name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Role      string `json:"role" binding:"required,oneof=ADMIN CUSTOMER"`
		Password  string `json:"password" binding:"required,min=8,isPassword"`
		CompanyID string `json:"company_id"`
	}

	UpdateUserRequest struct {
		ID       string `json:"id"`
		Name     string `json:"name" binding:"required"`
		Role     string `json:"role" binding:"required,oneof=ADMIN CUSTOMER"`
		Password string `json:"password" binding:"required,min=8,isPassword"`
	}
)
