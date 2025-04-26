package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// AddFriendForm представляет данные, необходимые для отправки запроса на дружбу.
type AddFriendForm struct {
	UserId       int `json:"user_id"`        // ID пользователя, который отправляет запрос.
	FriendUserId int `json:"friend_user_id"` // ID пользователя, получивший запрос на дружбу.
}

// AddFriend отправляет запрос на дружбу другому пользователю.
// @Summary Отправить запрос на дружбу
// @Description Позволяет пользователю отправить запрос на дружбу другому пользователю.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user body AddFriendForm true "Данные для отправки запроса на дружбу"
// @Success 204 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /add_friend [post]
func AddFriend(c *gin.Context, db postgres.PostgresInterface) {
	var newAddFriendForm AddFriendForm
	if err := c.BindJSON(&newAddFriendForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	if err := db.AddFriendRequest(newAddFriendForm.UserId, newAddFriendForm.FriendUserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка сервиса",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
