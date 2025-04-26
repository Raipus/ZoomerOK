package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// AcceptFriendForm представляет данные, необходимые для принятия запроса на дружбу.
type AcceptFriendForm struct {
	UserId       int `json:"user_id"`        // ID пользователя, который принимает запрос.
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

	if err := db.AcceptFriendRequest(newAcceptFriendForm.UserId, newAcceptFriendForm.FriendUserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	redisUserFriend := memory.RedisUserFriend{
		UserId:    newAcceptFriendForm.UserId,
		FriendIds: []int{newAcceptFriendForm.FriendUserId},
	}
	redis.AddUserFriend(redisUserFriend)

	c.JSON(http.StatusNoContent, gin.H{})
}
