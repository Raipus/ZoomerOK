package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type AcceptFriendForm struct {
	UserId       int `json:"user_id"`
	FriendUserId int `json:"friend_user_id"`
}

func AcceptFriend(c *gin.Context, db postgres.PostgresInterface) {
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

	c.JSON(http.StatusOK, gin.H{})
}
