package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// DeleteFriendForm представляет данные, необходимые для удаления пользователя из друзей.
type DeleteFriendForm struct {
	UserId       int `json:"user_id"`        // ID пользователя, который хочет удалить друга.
	FriendUserId int `json:"friend_user_id"` // ID пользователя, который должен быть удален.
}

// DeleteFriend отправляет запрос на удалении дружбы между пользователями.
// @Summary Удалить из друзей
// @Description Позволяет пользователю удалить друга из друзей.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user body DeleteFriendForm true "Данные для удаления из друзей"
// @Success 204 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Router /delete_friend [delete]
func DeleteFriend(c *gin.Context, db postgres.PostgresInterface, redis memory.RedisInterface) {
	var deleteFriendForm DeleteFriendForm
	if err := c.BindJSON(&deleteFriendForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	if err := db.DeleteFriendRequest(deleteFriendForm.UserId, deleteFriendForm.FriendUserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Пользователи не являются друзьями",
		})
		return
	}

	redis.DeleteUserFriend(deleteFriendForm.UserId, deleteFriendForm.FriendUserId)
	c.JSON(http.StatusNoContent, gin.H{})
}
