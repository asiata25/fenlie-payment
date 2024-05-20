package middlewareDto

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	JwtClaim struct {
		jwt.StandardClaims
		Username  string `json:"username"`
		CompanyID string `json:"company_id"`
		Roles     string `json:"role,omitempty"`
	}

	UserInfo struct {
		Email     string `json:"email"`
		CompanyID string `json:"company_id"`
		Roles     string `json:"roles,omitempty"`
	}

	LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)
