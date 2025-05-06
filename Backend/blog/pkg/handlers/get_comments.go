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

// GetComments отправляет запрос на получение комментариев для конкретного поста.
// @Summary Отправить запрос на получение комментариев.
// @Description Позволяет пользователю отправить запрос для получения комментариев для конкретного поста.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} gin.H {"message": "Success", "comments": []gin.H{"user": {"id": 0, "login": "", "name": "", "image": ""}, "body": {"id": 0, "post_id": 0, "text": "", "time": ""}}}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 404 {object} gin.H {"error": "Пост не найден"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /post/:post_id/comments [get]
func GetComments(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface, messageStore memory.MessageStoreInterface) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	postIdStr := c.Param("post_id")

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

	comments, err := db.GetComments(postId, page)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	userIds := make(map[int64]bool)
	for _, comment := range comments {
		userIds[int64(comment.UserId)] = true
	}

	ids := make([]int64, 0, len(userIds))
	for userId := range userIds {
		ids = append(ids, userId)
	}

	getUsersRequest := &pb.GetUsersRequest{
		Ids: ids,
	}
	if err := broker.PushUsers(getUsersRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	getUsersResponse, err := messageStore.ProcessPushUsers(getUsersRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Комментарии не найдены"})
		return
	}

	if len(getUsersResponse.Ids) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Комментарии не найдены"})
		log.Println("Empty response from message queue")
		return
	}

	userMap := make(map[int]*pb.GetUserResponse)
	for _, user := range getUsersResponse.Users {
		userMap[int(user.Id)] = user
	}

	var responseComments []gin.H
	for _, comment := range comments {
		user, exists := userMap[comment.UserId]
		if !exists {
			continue
		}

		responseComments = append(responseComments, gin.H{
			"user": gin.H{
				"id":    float64(user.Id),
				"login": user.Login,
				"name":  user.Name,
				"image": user.Image,
			},
			"body": gin.H{
				"id":      float64(comment.Id),
				"post_id": float64(comment.PostId),
				"text":    comment.Text,
				"time":    comment.Time,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": responseComments,
	})
}
