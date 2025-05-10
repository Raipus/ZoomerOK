package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type FindUserForm struct {
	Login string `json:"login"` // Login
}

func FindUser(c *gin.Context, db postgres.PostgresInterface, redis memory.RedisInterface) {
	var newFindUser FindUserForm
	if err := c.BindJSON(&newFindUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	user := db.GetUserByLogin(newFindUser.Login)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "пользователь не найден!",
		})
		return
	}

	redisUser := redis.GetUser(user.Id)

	c.JSON(http.StatusOK, gin.H{
		"id":    user.Id,
		"login": user.Login,
		"name":  user.Name,
		"image": redisUser.Image,
	})
}
