package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type CheckLoginForm struct {
	Login string
}

func CheckLogin(c *gin.Context, db postgres.PostgresInterface) {
	var newCheckLoginForm CheckLoginForm
	if err := c.BindJSON(&newCheckLoginForm); err != nil {
		return
	}

	user := db.GetUserByLogin(newCheckLoginForm.Login)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusOK, gin.H{
			"email": user.Email,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
	})
}
