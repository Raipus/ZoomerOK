package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type DeleteCommentForm struct {
	CommentId int
}

func DeleteComment(c *gin.Context, db postgres.PostgresInterface) {
	var deleteCommentForm DeleteCommentForm
	if err := c.ShouldBindJSON(&deleteCommentForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	userId := 1
	if err := db.DeleteComment(userId, deleteCommentForm.CommentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении комментария"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
