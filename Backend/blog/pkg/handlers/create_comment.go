package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type CreateCommentForm struct {
	PostId int
	Text   string
}

func CreateComment(c *gin.Context, db postgres.PostgresInterface) {
	var createCommentForm CreateCommentForm
	if err := c.BindJSON(&createCommentForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	userId := 1
	if err := db.CreateComment(userId, createCommentForm.Text); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании комментария"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
