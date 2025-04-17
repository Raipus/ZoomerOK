package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type CreateCommentForm struct {
	PostId int
	Text   string
}

func CreateComment(c *gin.Context, db postgres.PostgresInterface) {
	var createCommentForm CreateCommentForm
	if err := c.ShouldBindJSON(&createCommentForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	userId := 1
	if err := db.CreateComment(userId, &createCommentForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании комментария"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
