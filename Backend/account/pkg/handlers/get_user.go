package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, db postgres.PostgresInterface) {
	login := c.Param("login")

	user := db.GetUserByLogin(login)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "пользователь не найден!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"login":           user.Login,
		"name":            user.Name,
		"email":           user.Email,
		"confirmed_email": user.ConfirmedEmail,
		"image":           user.Image,
	})
}
