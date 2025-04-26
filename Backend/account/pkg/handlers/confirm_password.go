package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/gin-gonic/gin"
)

// ConfirmPassword подтверждает использование права на изменение пароля посредством перехода по ссылке.
// @Summary Подтверждает пароль
// @Description Позволяет пользователю подтвердить свое право на изменение пароля посредством перехода через соответствующую ссылку.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 204 {object} gin.H {}
// @Failure 404 {object} gin.H{"error": "User not found"}
// @Router /confirm_password/:reset_link [put]
func ConfirmPassword(c *gin.Context, cache caching.CachingInterface) {
	resetLink := c.Param("reset_link")

	login := cache.GetCacheResetLink(resetLink)
	if login == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
