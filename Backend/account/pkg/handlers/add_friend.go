package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type AddFriendForm struct {
	UserId       int `json:"user_id"`
	FriendUserId int `json:"friend_user_id"`
}

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
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
