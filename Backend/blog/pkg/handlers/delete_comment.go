package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/gin-gonic/gin"
)

func DeleteComment(c *gin.Context, db postgres.PostgresInterface, broker broker.BrokerInterface) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID комментария"})
		return
	}

	userId := c.MustGet("userId").(int)
	if err := db.DeleteComment(userId, commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении комментария"})
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
