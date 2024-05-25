package user

type (
	LoginRequest struct {
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		CompanyID string `json:"company_id"`
	}

	UserResponse struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		Password string `json:"password,omitempty"`
		Company  string `json:"company,omitempty"`
	}

	CreateUserRequest struct {
		Name      string `json:"name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Role      string `json:"role" binding:"required,oneof=ADMIN EMPLOYEE"`
		Password  string `json:"password" binding:"required,min=8,isPassword"`
		CompanyID string `json:"company_id"`
	}

	UpdateUserRequest struct {
		ID        string `json:"id"`
		Name      string `json:"name" binding:"required"`
		Role      string `json:"role" binding:"required,oneof=ADMIN EMPLOYEE"`
		Password  string `json:"password,omitempty" binding:"omitempty,min=8,isPassword"`
		CompanyID string `json:"company_id"`
	}

	UserJWT struct {
		ID   string `json:"id"`
		Role string `json:"role"`
	}
)
