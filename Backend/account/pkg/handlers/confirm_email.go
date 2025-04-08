package handlers

import (
	"net/http"
	"strings"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/gin-gonic/gin"
)

func ConfirmEmail(c *gin.Context) {
	fullURLWithParams := c.Request.URL.String()
	splitedURL := strings.Split(fullURLWithParams, "/")
	confirmationLink := splitedURL[len(splitedURL)-1]

	username := caching.GetCacheConfirmationLink(confirmationLink)
	if username == "" {
		c.JSON(http.StatusNotFound, gin.H{})
	} else {
		caching.DeleteCacheConfirmationLink(confirmationLink)
		c.JSON(http.StatusOK, gin.H{})
	}
}
