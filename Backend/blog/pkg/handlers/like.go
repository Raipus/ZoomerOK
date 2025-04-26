package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// Like отправляет запрос, чтобы поставить лайк на пост, если он уже присутствует, то снимает его с поста
// @Summary Поставить лайк на пост (или снять его)
// @Description Позволяет пользователю поставить лайк на пост (или снять его) по идентификатору
// @Tags posts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param post_id path int true "ID поста"
// @Success 204
// @Failure 400 {object} gin.H{"error": "Неверный формат ID комментария"}
// @Failure 500 {object} gin.H{"error": "User ID not found in context"}
// @Router /post/{post_id}/like [post]
func Like(c *gin.Context, db postgres.PostgresInterface) {
	postIdStr := c.Param("post_id")

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

	userIdInterface, exists := c.Get("user_id")
	log.Println("userIdInterface", userIdInterface)
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
	if err := db.Like(userId, postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
