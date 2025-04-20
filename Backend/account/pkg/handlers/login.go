package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	LoginOrEmail string `json:"login_or_email"`
	Password     string `json:"password"`
}

func Login(c *gin.Context, db postgres.PostgresInterface) {
	var newLoginForm LoginForm
	if err := c.BindJSON(&newLoginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	token, error := db.Login(newLoginForm.LoginOrEmail, newLoginForm.Password)
	if error != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
