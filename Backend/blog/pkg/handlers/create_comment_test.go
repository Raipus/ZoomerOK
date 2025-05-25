package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/blog/pkg/handlers"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	r := router.SetupRouter(false)
	userId := 1
	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(userId))
		c.Next()
	})
	mockPostgres := new(postgres.MockPostgres)
	r.POST("/post/:post_id/create_comment", func(c *gin.Context) {
		handlers.CreateComment(c, mockPostgres)
	})

	postId := "123"
	commentText := "Это тестовый комментарий"

	// Установим ожидания для mock
	mockPostgres.On("CreateComment", userId, 123, commentText).Return(nil)

	// Создадим JSON-данные для запроса
	jsonData := gin.H{"text": commentText}
	jsonValue, _ := json.Marshal(jsonData)

	// Создадим запрос
	req, _ := http.NewRequest("POST", "/post/"+postId+"/create_comment", bytes.NewReader(jsonValue))
	req.Header.Set("Authorization", "Bearer testtoken")

	// Запишем ответ
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверим статус ответа
	assert.Equal(t, http.StatusCreated, w.Code)
	mockPostgres.AssertExpectations(t)

	// Проверка на пустой ответ
	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Empty(t, actualResponse)
}
