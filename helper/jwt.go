package helper

import (
	"finpro-fenlie/model/dto/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: move this file to pkg

var (
	applicationName  = "finpro-fenlie"
	jwtSigningMethod = jwt.SigningMethodHS256
	JwtSignatureKey  = []byte("finpro-fenlie")
)

func GenerateTokenJwt(userId, role string, expiredAt int64) (string, error) {
	claims := auth.JwtClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiredAt) * time.Minute)),
			Issuer:    applicationName,
		},
		UserId: userId,
		Role:   role,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}
