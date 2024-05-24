package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type (
	JwtClaim struct {
		jwt.RegisteredClaims
		Username  string `json:"username"`
		CompanyID string `json:"company_id"`
		Role      string `json:"role"`
	}
)
