package handlers

import (
	"log"
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// ConfirmEmail подтверждает почту посредством перехода по ссылке.
// @Summary Подтверждает почту
// @Description Позволяет пользователю подтвердить свою почту посредством перехода через соответствующую ссылку.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 204 {object} gin.H {}
// @Failure 404 {object} gin.H{"error": "User not found"}
// @Router /confirm_email/:confirmation_link [put]
func ConfirmEmail(c *gin.Context, db postgres.PostgresInterface, cache caching.CachingInterface, redis memory.RedisInterface) {
	confirmationLink := c.Param("confirmation_link")

	login := cache.GetCacheConfirmationLink(confirmationLink)
	if login == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	} else {
		user, confirmed := db.ConfirmEmail(login)
		if !confirmed {
			c.JSON(http.StatusNotFound, gin.H{"error": "Login not found"})
		}
		cache.DeleteCacheConfirmationLink(confirmationLink)

		token, exists := c.Get("token")
		log.Println(token)
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not found in context"})
			return
		}

		tokenStr, ok := token.(string)
		log.Println(tokenStr)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token type"})
			return
		}

		redisAuthorization := memory.RedisAuthorization{
			UserId:         user.Id,
			Token:          tokenStr,
			Login:          user.Login,
			Email:          user.Email,
			ConfirmedEmail: user.ConfirmedEmail,
		}
		redis.SetAuthorization(redisAuthorization)
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
