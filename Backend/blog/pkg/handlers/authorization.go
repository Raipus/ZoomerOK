package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/blog/pkg/memory"
	"github.com/Raipus/ZoomerOK/blog/pkg/security"
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
func AuthMiddleware(broker broker.BrokerInterface, messageQueue memory.MessageQueue) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		_, err := security.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			log.Println("Token validation error:", err)
			return
		}

		authorizationRequest := pb.AuthorizationRequest{
			Token: tokenString,
		}
		if err := broker.Authorization(&authorizationRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
			return
		}
		time.Sleep(time.Millisecond * 200)
		message := messageQueue.GetLastMessage()
		log.Println("message:", message)
		authorizationResponse, ok := message.(*pb.AuthorizationResponse)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid response"})
			log.Println("Invalid response from message queue")
			return
		}

		log.Println("authorizationResponse", authorizationResponse)
		if authorizationResponse.Id == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid response"})
			log.Println("Empty response from message queue")
			return
		}

		if !authorizationResponse.ConfirmedEmail {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Почта не подтверждена!"})
			log.Println("Почта не подтверждена")
			return
		}

		c.Set("token", tokenString)
		c.Set("user_id", float64(authorizationResponse.Id))
		c.Set("login", authorizationResponse.Login)
		c.Set("email", authorizationResponse.Email)
		c.Next()
	}
}
