package userDto

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
)
