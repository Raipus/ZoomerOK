package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/gin-gonic/gin"
)

func Like(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface) {
	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID поста"})
		return
	}

	userId := c.MustGet("userId").(int)
	if err := db.Like(postId, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении лайка"})
		return
	}

	// Отправляем сообщение в брокер
	getUserRequest := &pb.GetUserRequest{Id: int32(userId)}
	if err := broker.PushGetUser(getUserRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при отправке сообщения в брокер"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
