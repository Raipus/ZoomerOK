package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware проверяет JWT токен и аутентифицирует пользователя
// @Summary Проверка авторизации
// @Description Проверяет JWT токен в заголовке Authorization
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router / [get]
func AuthMiddleware(db postgres.PostgresInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := security.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			log.Println("Token validation error:", err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			log.Println("Invalid token claims")
			return
		}

		id, ok := claims["id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid id claim"})
			log.Println("Invalid id claim")
			return
		}

		login, ok := claims["login"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid login claim"})
			log.Println("Invalid login claim")
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email claim"})
			log.Println("Invalid email claim")
			return
		}

		loginUser := db.GetUserByLogin(login)
		if loginUser.Login == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден!"})
			log.Println("Пользователь не найден")
			return
		} else if loginUser.ConfirmedEmail == false && !strings.Contains(c.Request.URL.Path, "confirm_email") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Почта не подтверждена!"})
			log.Println("Почта не подтверждена")
			return
		}

		c.Set("token", tokenString)
		c.Set("user_id", id)
		c.Set("login", login)
		c.Set("email", email)
		c.Next()
	}
}
