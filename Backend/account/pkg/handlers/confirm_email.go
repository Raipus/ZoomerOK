package handlers

import (
	"fmt"
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func ConfirmEmail(c *gin.Context, db postgres.PostgresInterface, cache caching.CachingInterface) {
	confirmationLink := c.Param("confirmation_link")

	username := cache.GetCacheConfirmationLink(confirmationLink)
	fmt.Println("asfd", username)
	if username == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	} else {
		cache.DeleteCacheConfirmationLink(confirmationLink)
		c.JSON(http.StatusOK, gin.H{})
	}
}
