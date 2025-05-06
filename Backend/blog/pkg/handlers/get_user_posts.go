package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// GetUserPosts отправляет запрос, чтобы получить посты определенного пользователя
// @Summary Получить посты определенного пользователя
// @Description Возвращает список постов для заданного пользователя с пагинацией
// @Tags posts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param page query int false "Номер страницы" default(1)
// @Success 200 {object} gin.H{"posts": []gin.H{"user": {"id": float64, "login": string, "name": string, "image": string}, "body": {"id": float64, "text": string, "image": string, "time": string}}}
// @Failure 404 {object} gin.H{"error": "Пост не найден"}
// @Failure 500 {object} gin.H{"error": "User ID not found in context"}
// @Router /user/:id/posts [get]
func GetUserPosts(c *gin.Context, db postgres.PostgresInterface) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID комментария"})
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	var responsePosts []gin.H

	posts, err := db.GetPosts([]int{userId}, page)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Посты не найдены"})
		return
	}

	for _, post := range posts {
		responsePosts = append(responsePosts, gin.H{
			"body": gin.H{
				"id":    float64(post.Id),
				"text":  post.Text,
				"image": post.Image,
				"time":  post.Time,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": responsePosts,
	})
}
