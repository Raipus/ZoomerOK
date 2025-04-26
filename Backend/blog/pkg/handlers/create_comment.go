package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// CreateCommentForm представляет данные, необходимые для отправки создания комментария
type CreateCommentForm struct {
	Text string `json:"text"` // Text, текст комментария.
}

// CreateComment отправляет запрос на создание комментария под конкретным постом.
// @Summary Отправить запрос на создания комментария
// @Description Позволяет пользователю отправить форму для создания комментария для конкретного поста.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user body CreateCommentForm true "Данные для отправки форма для комментария"
// @Success 201 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /post/:post_id/create_comment [post]
func CreateComment(c *gin.Context, db postgres.PostgresInterface) {
	postIdStr := c.Param("post_id")

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

	var createCommentForm CreateCommentForm
	if err := c.BindJSON(&createCommentForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
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
	if err := db.CreateComment(userId, postId, createCommentForm.Text); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
