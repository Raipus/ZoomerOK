package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// DeleteUser отправляет запрос на удалении аккаунта.
// @Summary Удалить аккаунт
// @Description Позволяет пользователю удалить свой аккаунт.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 204 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 401 {object} gin.H {"error": "Пользователь не имеет права на данную операцию"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /user/:login [delete]
func DeleteUser(c *gin.Context, db postgres.PostgresInterface, redis memory.RedisInterface) {
	login := c.Param("login")

	user := db.GetUserByLogin(login)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "пользователь не найден!",
		})
		return
	}

	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not found in context"})
		return
	}

	tokenStr, ok := token.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token type"})
		return
	}

	if user.Login != login {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не имеет права на данную операцию"})
		return
	}

	redis.DeleteAuthorization(tokenStr)
	redis.DeleteUser(user.Id)
	redis.DeleteAllUserFriend(user.Id)
	db.DeleteUser(&user)
	c.JSON(http.StatusNoContent, gin.H{})
}
