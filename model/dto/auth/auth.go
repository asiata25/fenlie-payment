package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type (
	JwtClaim struct {
		jwt.RegisteredClaims
		UserId string `json:"user_id"`
		Role   string `json:"role"`
	}
)
