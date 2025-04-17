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
		return
	}

	db.DeleteFriendRequest(deleteFriendForm.UserId, deleteFriendForm.FriendUserId)
	c.JSON(http.StatusOK, gin.H{})
}
