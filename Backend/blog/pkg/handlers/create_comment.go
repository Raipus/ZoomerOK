package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface) {
	var comment postgres.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	userId := c.MustGet("userId").(int)
	if err := db.CreateComment(userId, &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании комментария"})
		return
	}

	getUserRequest := &pb.GetUserRequest{Id: int32(userId)}
	if err := broker.PushGetUser(getUserRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при отправке сообщения в брокер"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}
