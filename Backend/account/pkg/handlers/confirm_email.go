package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func ConfirmEmail(c *gin.Context, db postgres.PostgresInterface, cache caching.CachingInterface) {
	confirmationLink := c.Param("confirmation_link")

	login := cache.GetCacheConfirmationLink(confirmationLink)
	if login == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	} else {
		confirmed := db.ConfirmEmail(login)
		if !confirmed {
			c.JSON(http.StatusNotFound, gin.H{"error": "Login not found"})
		}
		cache.DeleteCacheConfirmationLink(confirmationLink)
		c.JSON(http.StatusOK, gin.H{})
	}
}
