package handlers

import (
	"net/http"
	"strconv"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context, db postgres.PostgresInterface) {
	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID поста"})
		return
	}

	post, err := db.GetPost(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	c.JSON(http.StatusOK, post)
}
