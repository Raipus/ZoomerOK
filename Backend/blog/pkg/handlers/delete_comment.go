package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// DeleteComment отправляет запрос на удаления комментария под конкретным постом.
// @Summary Отправить запрос на удаление комментария
// @Description Позволяет пользователю отправить запрос для удаления комментария для конкретного поста.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 204 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 401 {object} gin.H {"error": "Пользователь не имеет права на данную операцию"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /post/:post_id/comments/:comment_id [delete]
func DeleteComment(c *gin.Context, db postgres.PostgresInterface) {
	commentIdStr := c.Param("comment_id")

	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

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

	if err := db.DeleteComment(userId, commentId); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
