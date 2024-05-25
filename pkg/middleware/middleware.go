package middleware

import (
	"finpro-fenlie/helper"
	"finpro-fenlie/model/dto/auth"
	jsonDTO "finpro-fenlie/model/dto/json"
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
			c.Abort()
			return
		}

		company, err := companyUseCase.GetById(companyId)
		if err != nil {
			jsonDTO.NewResponseUnauthorized(c, err.Error())
			c.Abort()
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(company.SecretKey), []byte(companySecret)); err != nil {
			jsonDTO.NewResponseUnauthorized(c, "unauthorized")
			c.Abort()
			return
		}

		c.Request.Header.Add("companyId", companyId)

		c.Next()
	}
}

func JWTAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("accessToken")
		if !strings.Contains(authHeader, "Bearer") {
			jsonDTO.NewResponseUnauthorized(c, "invalid token")
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &auth.JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return helper.JwtSignatureKey, nil
		})
		if err != nil {
			jsonDTO.NewResponseUnauthorized(c, "invalid token")
			c.Abort()
			return
		}
		if !token.Valid {
			jsonDTO.NewResponseForbidden(c, "forbidden")
			c.Abort()
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
			c.Abort()
			return
		}

		c.Request.Header.Add("userId", claims.UserId)
		// userInfo := &user.UserJWT{
		// 	ID: claims.ID,
		// 	Role:  claims.Role,
		// }
		// c.Set("userInfo", userInfo)

		c.Next()
	}
}
