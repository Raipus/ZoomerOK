package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface) {
	postId := 1
	posts, err := db.GetPosts(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	var responsePosts []gin.H
	for _, post := range posts {
		responsePosts = append(responsePosts, gin.H{
			"id":      float64(post.Id),
			"user_id": float64(post.UserId),
			"text":    post.Text,
			"image":   post.Image,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": responsePosts,
	})
}
