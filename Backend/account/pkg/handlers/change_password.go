package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type ChangePasswordForm struct {
	Email       string `json:"email"`
	NewPassword string `json:"password"`
}

func ChangePassword(c *gin.Context, db postgres.PostgresInterface, cache caching.CachingInterface) {
	resetLink := c.Param("reset_link")

	var newChangePasswordForm ChangePasswordForm
	if err := c.BindJSON(&newChangePasswordForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	user := db.GetUserByEmail(newChangePasswordForm.Email)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Пользователь не найден",
		})
		return
	}

	if err := db.ChangePassword(&user, newChangePasswordForm.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера",
		})
	} else {
		cache.DeleteCacheResetLink(resetLink)
		c.JSON(http.StatusOK, gin.H{
			"email": newChangePasswordForm.Email,
		})
	}
}
