package handlers

import (
	"log"
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// AddFriendForm представляет данные, необходимые для отправки запроса на дружбу.
type AddFriendForm struct {
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

	log.Println("userId", userIdFloat)
	userId := int(18)
	if err := db.AddFriendRequest(userId, newAddFriendForm.FriendUserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка сервиса",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
