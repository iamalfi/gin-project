package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Role string

const (
	Admin  Role = "admin"
	Client Role = "client"
	Guest  Role = "guest"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

func AuthMiddleware(requiredRoles ...Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or invalid"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID, ok := claims["id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
			c.Abort()
			return
		}

		c.Set("id", userID)

		// Check if roles are required
		if len(requiredRoles) > 0 {
			userRole, ok := claims["role"].(string)
			if !ok {
				c.JSON(http.StatusForbidden, gin.H{"error": "Role not found in token"})
				c.Abort()
				return
			}

			for _, role := range requiredRoles {
				if userRole == string(role) {
					c.Next()
					return
				}
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access to this resource"})
			c.Abort()
		} else {
			// No roles required, just pass the middleware
			c.Next()
		}
	}
}
