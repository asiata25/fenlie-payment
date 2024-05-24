package middleware

import (
	"finpro-fenlie/helper"
	"finpro-fenlie/model/dto/auth"
	jsonDTO "finpro-fenlie/model/dto/json"
	"finpro-fenlie/model/dto/user"
	"finpro-fenlie/src/company"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuth(companyUseCase company.CompanyUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId, companySecret, ok := c.Request.BasicAuth()

		if !ok {
			jsonDTO.NewResponseUnauthorized(c, "unauthorized")
			return
		}

		company, err := companyUseCase.GetById(companyId)
		if err != nil {
			jsonDTO.NewResponseError(c, err.Error())
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(company.SecretKey), []byte(companySecret)); err != nil {
			jsonDTO.NewResponseUnauthorized(c, "unauthorized")
			return
		}
		c.Next()
	}
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
				if role == claims.Role {
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
			Role:      claims.Role,
		}
		c.Set("userInfo", userInfo)

		c.Next()
	}
}
