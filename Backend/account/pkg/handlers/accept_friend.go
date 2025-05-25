package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// AcceptFriendForm представляет данные, необходимые для принятия запроса на дружбу.
type AcceptFriendForm struct {
	FriendUserId int `json:"friend_user_id"` // ID пользователя, отправившего запрос на дружбу.
}

// AcceptFriend принимает запрос на дружбу.
// @Summary Принять запрос на дружбу
// @Description Позволяет пользователю принять запрос на дружбу от другого пользователя.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user body AcceptFriendForm true "Данные для принятия запроса на дружбу"
// @Success 204 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Router /accept_friend [put]
func AcceptFriend(c *gin.Context, db postgres.PostgresInterface, redis memory.RedisInterface) {
	var newAcceptFriendForm AcceptFriendForm
	if err := c.BindJSON(&newAcceptFriendForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
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
	if err := db.AcceptFriendRequest(userId, newAcceptFriendForm.FriendUserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	redisUserFriendFromMyUserId := memory.RedisUserFriend{
		UserId:    userId,
		FriendIds: []int{newAcceptFriendForm.FriendUserId},
	}
	redisUserFriendFromNotMyUserId := memory.RedisUserFriend{
		UserId:    newAcceptFriendForm.FriendUserId,
		FriendIds: []int{userId},
	}
	redis.AddUserFriend(redisUserFriendFromMyUserId)
	redis.AddUserFriend(redisUserFriendFromNotMyUserId)

	c.JSON(http.StatusNoContent, gin.H{})
}
