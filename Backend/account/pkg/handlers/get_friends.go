package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/gin-gonic/gin"
)

func GetFriends(c *gin.Context, redis memory.RedisInterface) {
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

	userFriends := redis.GetUserFriends(userId)
	if len(userFriends.FriendIds) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "У вас сейчас нет друзей!"})
		return
	}

	users := redis.GetUsers(userFriends.FriendIds)
	if len(users) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	var responseUsers []gin.H
	for _, user := range users {
		responseUsers = append(responseUsers, gin.H{
			"user": gin.H{
				"id":    user.UserId,
				"login": user.Login,
				"name":  user.Name,
				"image": user.Image,
			},
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"users": responseUsers,
	})
}
