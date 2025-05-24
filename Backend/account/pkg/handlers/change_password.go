package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// ChangePasswordForm представляет данные, необходимые для изменения пароля пользователя.
type ChangePasswordForm struct {
	NewPassword string `json:"password"` // NewPassword, новый пароль пользователя.
}

// ChangePassword отправляет запрос на изменение пароля пользователя.
// @Summary Отправить запрос на изменение пароля.
// @Description Позволяет пользователю отправить запрос на изменение собственного пароля, если он его забыл.
// @Accept json
// @Produce json
// @Param user body ChangePasswordForm true "Данные для изменения пароля пользователя"
// @Success 204 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /change_password/:reset_link [put]
func ChangePassword(c *gin.Context, db postgres.PostgresInterface, cache caching.CachingInterface) {
	resetLink := c.Param("reset_link")

	var newChangePasswordForm ChangePasswordForm
	if err := c.BindJSON(&newChangePasswordForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	login := cache.GetCacheResetLink(resetLink)
	if login == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	} else {
		user := db.GetUserByLogin(login)
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
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
