package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func GetUnacceptedFriends(c *gin.Context, db postgres.PostgresInterface) {
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

	unacceptedFriends, err := db.GetUnacceptedFriends(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	var responseUsers []gin.H
	for _, user := range unacceptedFriends {
		responseUsers = append(responseUsers, gin.H{
			"unaccepted_friend": gin.H{
				"id":    user.Id,
				"login": user.Login,
				"name":  user.Name,
				"image": user.Image,
			},
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"unaccepted_friends": responseUsers,
	})
}
