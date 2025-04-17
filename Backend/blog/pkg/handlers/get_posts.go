package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/broker"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface) {
	userId := c.MustGet("userId").(int)
	posts, err := db.GetPosts(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении постов"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
