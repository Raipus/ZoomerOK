package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface) {
	postIdStr := c.Param("post_id")

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

	comments, err := db.GetComments(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	var responseComments []gin.H
	for _, comment := range comments {
		responseComments = append(responseComments, gin.H{
			"id":      float64(comment.Id),
			"user_id": float64(comment.UserId),
			"post_id": float64(comment.PostId),
			"text":    comment.Text,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": responseComments,
	})
}
