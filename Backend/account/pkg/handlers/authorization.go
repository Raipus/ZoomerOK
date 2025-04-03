package handlers

import (
	"strings"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "API token required"})
			return
		}

		s := strings.Split(token, ":")
		if s[0] != "Token" || !postgres.UUIDExists(s[1]) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token required"})
			return
		}

		c.Next()
	}
}
