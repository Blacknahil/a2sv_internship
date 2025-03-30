package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"task_manager_api_auth/constants"
	"task_manager_api_auth/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret key not set in the environment variables"})
			c.Abort()
			return
		}

		claims := &models.CustomJWTClaims{}

		token, err := jwt.ParseWithClaims(authParts[1], claims, func(token *jwt.Token) (interface{}, error) {

			if token.Method.Alg() != constants.SigningMethod {
				return nil, fmt.Errorf("unexpeected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("Role", claims.Role)
		c.Set("UserID", claims.UserID)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		role, _ := c.Get("Role")
		if role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}
		c.Next()
	}
}
