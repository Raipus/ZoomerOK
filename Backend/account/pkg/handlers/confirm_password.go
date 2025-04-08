package handlers

import (
	"net/http"
	"strings"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/gin-gonic/gin"
)

func ConfirmPassword(c *gin.Context) {
	fullURLWithParams := c.Request.URL.String()
	splitedURL := strings.Split(fullURLWithParams, "/")
	resetLink := splitedURL[len(splitedURL)-1]

	username := caching.GetCacheResetLink(resetLink)
	if username == "" {
		c.JSON(http.StatusNotFound, gin.H{})
	} else {
		caching.DeleteCacheResetLink(resetLink)
		c.JSON(http.StatusOK, gin.H{})
	}
}
