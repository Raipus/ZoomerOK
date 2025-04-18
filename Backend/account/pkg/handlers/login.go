package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context, db postgres.PostgresInterface) {
	var newLoginForm LoginForm
	if err := c.BindJSON(&newLoginForm); err != nil {
		return
	}

	logined, error := db.Login(newLoginForm.Email, newLoginForm.Password)
	if !logined {
		if error == "Ошибка сервера" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": error})
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}
