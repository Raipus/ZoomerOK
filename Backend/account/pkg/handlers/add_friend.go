package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type AddFriendForm struct {
	UserId       int
	FriendUserId int
}

func AddFriend(c *gin.Context, db postgres.PostgresInterface) {
	var newAddFriendForm AddFriendForm
	if err := c.BindJSON(&newAddFriendForm); err != nil {
		return
	}

	db.AcceptFriendRequest(newAddFriendForm.UserId, newAddFriendForm.FriendUserId)
	c.JSON(http.StatusOK, gin.H{})
}
