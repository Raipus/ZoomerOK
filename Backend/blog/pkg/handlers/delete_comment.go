package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func DeleteComment(c *gin.Context, db postgres.PostgresInterface) {
	commentIdStr := c.Param("comment_id")

	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

	userId := 1
	if err := db.DeleteComment(userId, commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении комментария"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
