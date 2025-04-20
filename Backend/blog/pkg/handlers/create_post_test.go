package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)

	createCommentData := CreatePostForm{
		PostId: 1,
		Text:   "Новый комментарий",
		Photo:  nil,
	}

	r.POST("/create_post", func(c *gin.Context) {
		CreatePost(c, mockPostgres)
	})

	mockPostgres.On("CreatePost", 1, createCommentData.Text, createCommentData.Photo).Return(nil)

	jsonData, err := json.Marshal(createCommentData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/create_post", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	mockPostgres.AssertExpectations(t)
}
