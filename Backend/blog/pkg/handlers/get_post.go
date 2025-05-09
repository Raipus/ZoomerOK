package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/blog/pkg/memory"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// GetPost отправляет запрос, чтобы получить Id поста
// @Summary Получить пост по ID
// @Description Возвращает пост по заданному идентификатору
// @Tags posts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param post_id path int true "ID поста"
// @Success 200 {object} gin.H{"post": {"user": {"id": float64, "login": string, "name": string, "image": string}, "body": {"id": float64, "text": string, "image": string, "time": string}}}
// @Failure 400 {object} gin.H{"error": "Неверный формат ID комментария"}
// @Failure 404 {object} gin.H{"error": "Пост не найден"}
// @Router /post/{post_id} [get]
func GetPost(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface, messageStore memory.MessageStoreInterface) {
	postIdStr := c.Param("post_id")

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

	post, err := db.GetPost(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	getUserRequest := &pb.GetUserRequest{
		Id: int64(post.UserId),
	}
	if err := broker.PushUser(getUserRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	getUserResponse, err := messageStore.ProcessPushUser(getUserRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		log.Println("Invalid response from message queue")
		return
	}

	if getUserResponse.Id == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid response"})
		log.Println("Empty response from message queue")
		return
	}

	commentCountMap, likeCountMap, err := db.GetCountCommentsAndLikes([]int{post.Id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": gin.H{
			"user": gin.H{
				"id":    float64(getUserResponse.Id),
				"login": getUserResponse.Login,
				"name":  getUserResponse.Name,
				"image": getUserResponse.Image,
			},
			"body": gin.H{
				"id":    float64(post.Id),
				"text":  post.Text,
				"image": post.Image,
				"time":  post.Time,
				"number_of_comments": float64(commentCountMap[post.Id]),
				"number_of_likes": float64(likeCountMap[post.Id]),
			},
		},
	})
}
