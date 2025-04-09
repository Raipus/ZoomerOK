package handlers

import (
	"net/http"
	"strings"

	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware проверяет наличие и валидность JWT в заголовке Authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // Если токен не был найден
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		token, err := security.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Если токен валиден, продолжаем обработку запроса
		c.Next()
	}
}
