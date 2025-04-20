package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context, db postgres.PostgresInterface) {
	postIdStr := c.Param("post_id")

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

	userId := 1
	if err := db.DeletePost(userId, postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении комментария"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
