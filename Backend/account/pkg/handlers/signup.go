package handlers

import (
	"fmt"
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
)

type SignupForm struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context, db postgres.PostgresInterface, smtp security.SMTPInterface, cache caching.CachingInterface) {
	var newSignupForm SignupForm
	if err := c.BindJSON(&newSignupForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	// Проверяем логин
	loginUser := db.GetUserByLogin(newSignupForm.Login)
	if loginUser.Login != "" {
		c.JSON(http.StatusConflict, gin.H{
			"error": fmt.Sprintf("Логин '%s' уже существует", newSignupForm.Login),
		})
		return
	}

	// Проверяем электронную почту
	emailUser := db.GetUserByEmail(newSignupForm.Email)
	if emailUser.Email != "" {
		c.JSON(http.StatusConflict, gin.H{
			"error": fmt.Sprintf("Электронная почта '%s' уже существует", newSignupForm.Email),
		})
		return
	}

	token, registered := db.Signup(newSignupForm.Login, newSignupForm.Name, newSignupForm.Email, newSignupForm.Password)

	if registered {
		smtp.SendConfirmEmail(newSignupForm.Name, newSignupForm.Email, cache)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервиса",
		})
	}
}
