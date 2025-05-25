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

// GetPosts отправляет запрос, чтобы получить посты
// @Summary Получить посты пользователя
// @Description Возвращает список постов для заданного пользователя с пагинацией
// @Tags posts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param page query int false "Номер страницы" default(1)
// @Success 200 {object} gin.H{"posts": []gin.H{"user": {"id": float64, "login": string, "name": string, "image": string}, "body": {"id": float64, "text": string, "image": string, "time": string}}}
// @Failure 404 {object} gin.H{"error": "Пост не найден"}
// @Failure 500 {object} gin.H{"error": "User ID not found in context"}
// @Router /posts [get]
func GetPosts(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface, messageStore memory.MessageStoreInterface) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	userIdInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	userIdFloat, ok := userIdInterface.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	userId := int(userIdFloat)

	getUserFriendRequest := &pb.GetUserFriendRequest{
		Id: int64(userId),
	}
	if err := broker.PushUserFriend(getUserFriendRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	getUserFriendResponse, err := messageStore.ProcessPushUserFriend(getUserFriendRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Посты не найдены"})
		return
	}

	if getUserFriendResponse.Id == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Посты не найдены"})
		log.Println("Empty response from message queue")
		return
	}

	intIds := make([]int, len(getUserFriendResponse.Ids))
	for i, id := range getUserFriendResponse.Ids {
		intIds[i] = int(id)
	}

	intIds = append(intIds, userId)

	posts, err := db.GetPosts(intIds, page)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Посты не найдены"})
		return
	}

	userIds := make(map[int64]bool)
	postIds := make([]int, 0, len(posts))
	for _, post := range posts {
		userIds[int64(post.UserId)] = true
		postIds = append(postIds, int(post.Id))
	}

	ids := make([]int64, 0, len(userIds))
	for userId := range userIds {
		ids = append(ids, userId)
	}

	log.Println("unique users", ids)
	getUsersRequest := &pb.GetUsersRequest{
		Ids: ids,
	}
	if err := broker.PushUsers(getUsersRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	getUsersResponse, err := messageStore.ProcessPushUsers(getUsersRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Посты не найдены"})
		return
	}

	if len(getUsersResponse.Ids) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Посты не найдены"})
		log.Println("Empty response from message queue")
		return
	}

	userMap := make(map[int64]*pb.GetUserResponse)
	for _, user := range getUsersResponse.Users {
		userMap[int64(user.Id)] = user
	}

	commentCountMap, likeCountMap, err := db.GetCountCommentsAndLikes(postIds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	var responsePosts []gin.H
	for _, post := range posts {
		user, exists := userMap[int64(post.UserId)]
		if !exists {
			continue
		}

		responsePosts = append(responsePosts, gin.H{
			"user": gin.H{
				"id":    float64(user.Id),
				"login": user.Login,
				"name":  user.Name,
				"image": user.Image,
			},
			"body": gin.H{
				"id":                 float64(post.Id),
				"text":               post.Text,
				"image":              post.Image,
				"time":               post.Time,
				"number_of_comments": float64(commentCountMap[post.Id]),
				"number_of_likes":    float64(likeCountMap[post.Id]),
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": responsePosts,
	})
}
