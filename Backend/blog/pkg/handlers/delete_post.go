package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type DeletePostForm struct {
	PostId int
}

func DeletePost(c *gin.Context, db postgres.PostgresInterface) {
	var deletePostForm DeletePostForm
	if err := c.ShouldBindJSON(&deletePostForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	userId := 1
	if err := db.DeletePost(userId, commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении комментария"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
