package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// CreatePostForm представляет данные, необходимые для отправки создания поста
type CreatePostForm struct {
	Text  string `json:"text"`            // Text, поста.
	Image []byte `json:"image,omitempty"` // Image, base64 кодированное изображение.
}

// CreatePost отправляет запрос на создание поста (можно добавить одно изображение).
// @Summary Отправить запрос на создания поста
// @Description Позволяет пользователю отправить форму для создания поста.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user body CreatePostForm true "Данные для отправки формы для поста"
// @Success 201 {object} gin.H {"id": "Id поста"}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router create_post [post]
func CreatePost(c *gin.Context, db postgres.PostgresInterface) {
	var createPostForm CreatePostForm
	if err := c.ShouldBindJSON(&createPostForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
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
	postId, err := db.CreatePost(userId, createPostForm.Text, createPostForm.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании поста"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": postId,
	})
}
