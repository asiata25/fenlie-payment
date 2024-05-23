package middleware

import (
	"finpro-fenlie/helper"
	"finpro-fenlie/model/dto/auth"
	jsonDTO "finpro-fenlie/model/dto/json"
	"finpro-fenlie/model/dto/user"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func BasicAuth(c *gin.Context) {
	user, password, ok := c.Request.BasicAuth()
	if !ok {
		jsonDTO.NewResponseUnauthorized(c, "unauthorized")
		return
	}

	if user != os.Getenv("CLIENT_ID") || password != os.Getenv("CLIENT_SECRET") {
		jsonDTO.NewResponseUnauthorized(c, "unauthorized")
		return
	}
	c.Next()
}

func JWTAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			jsonDTO.NewResponseUnauthorized(c, "invalid token")
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &auth.JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return helper.JwtSignatureKey, nil
		})
		if err != nil {
			jsonDTO.NewResponseUnauthorized(c, "invalid token")
			return
		}
		if !token.Valid {
			jsonDTO.NewResponseForbidden(c, "forbidden")
			return
		}

		// validation role
		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims.Roles {
					validRole = true
					break
				}
			}
		}
		if !validRole {
			jsonDTO.NewResponseForbidden(c, "forbidden")
			return
		}

		userInfo := &user.UserResponse{
			Email:     claims.Username,
			CompanyID: claims.CompanyID,
			Roles:     claims.Roles,
		}
		c.Set("userInfo", userInfo)

		c.Next()
	}
}
