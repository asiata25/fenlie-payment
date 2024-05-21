package user

import "github.com/google/uuid"

type (
	User struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password,omitempty"`
		CompanyID uuid.UUID `json:"company_id,omitempty"`
		Role      string    `json:"role,omitempty"`
	}

	Paging struct {
		Page int `json:"page"`
		Size int `json:"size"`
	}

	GetResponse struct {
		Data       []User `json:"data"`
		Pagination Paging `json:"paging"`
		TotalData  int64  `json:"totalData"`
	}

	LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	UserInfo struct {
		Email     string `json:"email"`
		CompanyID string `json:"company_id"`
		Roles     string `json:"roles,omitempty"`
	}
)
