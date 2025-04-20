package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type CreatePostForm struct {
	PostId int
	Text   string
	Photo  []byte
}

func CreatePost(c *gin.Context, db postgres.PostgresInterface) {
	var createPostForm CreatePostForm
	if err := c.ShouldBindJSON(&createPostForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	userId := 1
	if err := db.CreatePost(userId, createPostForm.Text, createPostForm.Photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании поста"})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
