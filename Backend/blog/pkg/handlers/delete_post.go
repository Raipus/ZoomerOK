package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// DeleteComment отправляет запрос на удаления поста.
// @Summary Отправить запрос на удаление поста
// @Description Позволяет пользователю отправить запрос для удаления поста.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 204 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 401 {object} gin.H {"error": "Пользователь не имеет права на данную операцию"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /post/:post_id [delete]
func DeletePost(c *gin.Context, db postgres.PostgresInterface) {
	postIdStr := c.Param("post_id")

	postId, err := strconv.Atoi(postIdStr)
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
	if err := db.DeletePost(userId, postId); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
