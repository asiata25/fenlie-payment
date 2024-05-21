package middleware

import (
	"errors"
	"finpro-fenlie/model/dto/auth"
	jsonDTO "finpro-fenlie/model/dto/json"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	user, password, ok := c.Request.BasicAuth()
	if !ok {
		jsonDTO.NewResponseAuth(c, "Invalid Token")
		return
	}
	if user != os.Getenv("CLIENT_ID") || password != os.Getenv("CLIENT_SECRET") {
		jsonDTO.NewResponseAuth(c, "Unauthorized")
		return
	}
	c.Next()
}

var (
	applicationName  = "finpro-fenlie"
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte("finpro-fenlie")
)

func GenerateTokenJwt(username, role, companyID string, expiredAt int64) (string, error) {
	loginExpDuration := time.Duration(expiredAt) * time.Minute
	myExpiresAt := time.Now().Add(loginExpDuration).Unix()
	claims := auth.JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationName,
			ExpiresAt: myExpiresAt,
		},
		CompanyID: companyID,
		Username:  username,
		Roles:     role,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			jsonDTO.NewResponseAuth(c, "Invalid Token")
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &auth.JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSignatureKey, nil
		})
		if err != nil {
			jsonDTO.NewResponseAuth(c, "Invalid Token")
			return
		}
		if !token.Valid {
			jsonDTO.NewResponseAuth(c, "Forbidden")
			return
		}

		userInfo := &auth.UserInfo{
			Email:     claims.Username,
			CompanyID: claims.CompanyID,
			Roles:     claims.Roles,
		}
		c.Set("userInfo", userInfo)

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo, exists := c.Get("userInfo")
		if !exists {
			jsonDTO.NewResponseAuth(c, "Unauthorized")
			c.Abort()
			return
		}

		user, ok := userInfo.(*auth.UserInfo)
		if !ok {
			jsonDTO.NewResponseAuth(c, "Internal Server Error")
			c.Abort()
			return
		}

		isAdmin := false
		if user.Roles == "ADMIN" {
			isAdmin = true
		}

		if !isAdmin {
			jsonDTO.NewResponseAuth(c, "Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserInfo(ctx *gin.Context) (*auth.UserInfo, error) {
	userInfo, exists := ctx.Get("userInfo")
	if !exists {
		return nil, errors.New("unauthorized")
	}

	user, ok := userInfo.(*auth.UserInfo)
	if !ok {
		return nil, errors.New("internal server error")
	}

	return user, nil
}
