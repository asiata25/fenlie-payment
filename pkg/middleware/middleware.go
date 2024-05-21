package middleware

import (
	"errors"
	"finpro-fenlie/model/dto/json"
	"finpro-fenlie/model/dto/middlewareDto"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	user, password, ok := c.Request.BasicAuth()
	if !ok {
		json.NewResponseAuth(c, "Invalid Token", "01", "01")
		return
	}
	if user != os.Getenv("CLIENT_ID") || password != os.Getenv("CLIENT_SECRET") {
		json.NewResponseAuth(c, "Unauthorized", "01", "01")
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
	claims := middlewareDto.JwtClaim{
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
			json.NewResponseAuth(c, "Invalid Token", "01", "01")
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &middlewareDto.JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSignatureKey, nil
		})
		if err != nil {
			json.NewResponseAuth(c, "Invalid Token", "01", "01")
			return
		}
		if !token.Valid {
			json.NewResponseAuth(c, "Forbidden", "01", "01")
			return
		}

		userInfo := &middlewareDto.UserInfo{
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
			json.NewResponseAuth(c, "Unauthorized", "01", "01")
			c.Abort()
			return
		}

		user, ok := userInfo.(*middlewareDto.UserInfo)
		if !ok {
			json.NewResponseAuth(c, "Internal Server Error", "01", "01")
			c.Abort()
			return
		}

		isAdmin := false
		if user.Roles == "ADMIN" {
			isAdmin = true
		}

		if !isAdmin {
			json.NewResponseAuth(c, "Forbidden", "01", "01")
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserInfo(ctx *gin.Context) (*middlewareDto.UserInfo, error) {
	userInfo, exists := ctx.Get("userInfo")
	if !exists {
		return nil, errors.New("unauthorized")
	}

	user, ok := userInfo.(*middlewareDto.UserInfo)
	if !ok {
		return nil, errors.New("internal server error")
	}

	return user, nil
}