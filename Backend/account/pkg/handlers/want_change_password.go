package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
)

type WantChangePasswordForm struct {
	Email string `json:"email"`
}

func WantChangePassword(c *gin.Context, db postgres.PostgresInterface, smtp security.SMTPInterface, cache caching.CachingInterface) {
	var newWantChangePasswordForm WantChangePasswordForm
	if err := c.BindJSON(&newWantChangePasswordForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	user := db.GetUserByEmail(newWantChangePasswordForm.Email)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Пользователь не найден",
		})
		return
	}

	if err := smtp.SendChangePassword(user.Name, user.Email, cache); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email не найден",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": newWantChangePasswordForm.Email,
	})
}
