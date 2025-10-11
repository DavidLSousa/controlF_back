package middlewares

import (
	"controlF_back/internal/domain"
	"controlF_back/internal/token"
	"controlF_back/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return AuthMiddlewareWithRole(nil)
}

func AuthMiddlewareWithRole(roles []token.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := GetBearerToken(c)

		if len(authToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Not authorized"})
			return
		}

		claims, err := token.IsAuthorized(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: utils.PrintError(err)})
			return
		}

		if len(roles) > 0 {
			if !containsAny(roles, claims.Roles) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Not in role"})
				return
			}
		}

		c.Set("x-user-id", claims.RegisteredClaims.Subject)
		c.Set("x-claims", claims)
		c.Next()
	}
}

func containsAny(arr1, arr2 []token.RoleType) bool {
	for _, a := range arr1 {
		for _, b := range arr2 {
			if a == b {
				return true
			}
		}
	}
	return false
}

func GetBearerToken(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")

	t := strings.Split(authHeader, " ")
	if len(t) == 2 {
		return t[1]
	}

	return ""
}
