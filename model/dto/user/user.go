package user

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	UserResponse struct {
		Email     string `json:"email"`
		CompanyID string `json:"company_id"`
		Roles     string `json:"roles,omitempty"`
	}

	CreateUserRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Roles    string `json:"roles" binding:"required,oneof=ADMIN CUSTOMER"`
		Password string `json:"password" binding:"required,min=8,isPassword"`
	}
)
