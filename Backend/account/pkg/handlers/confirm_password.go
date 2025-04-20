package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/gin-gonic/gin"
)

func ConfirmPassword(c *gin.Context, cache caching.CachingInterface) {
	resetLink := c.Param("reset_link")

	login := cache.GetCacheResetLink(resetLink)
	if login == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	} else {
		cache.DeleteCacheResetLink(resetLink)
		c.JSON(http.StatusOK, gin.H{"message": "Password confirmed"})
	}
}
