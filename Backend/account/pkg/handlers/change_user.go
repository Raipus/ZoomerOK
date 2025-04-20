package handlers

import (
	"net/http"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type ChangeUserForm struct {
	Name     string     `json:"name"`
	Birthday *time.Time `json:"birthday"`
	Phone    string     `json:"phone"`
	City     string     `json:"city"`
	Image    []byte     `json:"image"`
}

func ChangeUser(c *gin.Context, db postgres.PostgresInterface) {
	login := c.Param("login")

	var newChangeUser ChangeUserForm
	if err := c.BindJSON(&newChangeUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	user := db.GetUserByLogin(login)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "пользователь не найден!",
		})
		return
	}

	if success := db.ChangeUser(&user); !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при изменении пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно обновлен"})
}
