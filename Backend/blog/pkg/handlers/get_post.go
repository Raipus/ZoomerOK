package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type GetPostForm struct {
	PostId int
}

func GetPost(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface) {
	var getPostForm GetPostForm
	if err := c.ShouldBindJSON(&getPostForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	post, err := db.GetPost(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	c.JSON(http.StatusOK, post)
}
