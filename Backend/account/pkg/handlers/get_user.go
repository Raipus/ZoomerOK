package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type GetUserForm struct {
	Login string
}

func GetUser(c *gin.Context, db postgres.PostgresInterface) {
	var newUserForm GetUserForm
	if err := c.BindJSON(&newUserForm); err != nil {
		return
	}

	user := db.GetUserByLogin(newUserForm.Login)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Пользователь не найден",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
	})
}
