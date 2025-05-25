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

func TestCreatePost(t *testing.T) {
	r := router.SetupRouter(false)
	userId := 1
	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(1))
		c.Next()
	})
	mockPostgres := new(postgres.MockPostgres)
	r.POST("/create_post", func(c *gin.Context) {
		handlers.CreatePost(c, mockPostgres)
	})

	postText := "Это тестовый пост"
	mockPostgres.On("CreatePost", userId, postText, []byte(nil)).Return(1, nil)

	jsonData := gin.H{"text": postText}
	jsonValue, _ := json.Marshal(jsonData)

	req, _ := http.NewRequest("POST", "/create_post", bytes.NewReader(jsonValue))
	req.Header.Set("Authorization", "Bearer testtoken")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockPostgres.AssertExpectations(t)

	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, float64(1), actualResponse["id"])
}
