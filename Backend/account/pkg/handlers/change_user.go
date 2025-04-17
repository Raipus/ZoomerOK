package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func ChangeUserHandler(c *gin.Context, db postgres.PostgresInterface) {
	var user postgres.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	user.Id = userId.(int)

	if success := db.ChangeUser(&user); !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при изменении пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно обновлен"})
}
