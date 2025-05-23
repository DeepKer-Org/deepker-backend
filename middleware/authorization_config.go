package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RoleAuthorization(requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		if len(secretKey) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT secret key not set"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userRoles, roleExists := claims["roles"].([]interface{})
			if !roleExists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Roles not found in token"})
				c.Abort()
				return
			}

			for _, role := range userRoles {
				for _, requiredRole := range requiredRoles {
					if role == requiredRole {
						c.Next()
						return
					}
				}
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
