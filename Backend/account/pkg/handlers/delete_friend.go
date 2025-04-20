package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type DeleteFriendForm struct {
	UserId       int
	FriendUserId int
}

func DeleteFriend(c *gin.Context, db postgres.PostgresInterface) {
	var deleteFriendForm DeleteFriendForm
	if err := c.BindJSON(&deleteFriendForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	if err := db.DeleteFriendRequest(deleteFriendForm.UserId, deleteFriendForm.FriendUserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
