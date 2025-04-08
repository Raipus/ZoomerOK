package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/gin-gonic/gin"
)

func ConfirmPassword(c *gin.Context, cache caching.CachingInterface) {
	resetLink := c.Param("reset_link")

	username := cache.GetCacheResetLink(resetLink)
	if username == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	cache.DeleteCacheResetLink(resetLink)

	c.JSON(http.StatusOK, gin.H{"message": "Password confirmed"})
}
