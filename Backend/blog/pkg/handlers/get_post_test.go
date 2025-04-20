package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockBroker := new(broker.MockBroker)

	r.GET("/post/:post_id", func(c *gin.Context) {
		GetPost(c, mockPostgres, mockBroker)
	})

	var postId int = 1
	postTime := time.Now()
	post := postgres.Post{
		Id:     postId,
		UserId: 2,
		Text:   "Это пост",
		Image:  nil,
		Time:   &postTime,
	}

	mockPostgres.On("GetPost", postId).Return(&post, nil)

	req, err := http.NewRequest("GET", "/post/"+strconv.Itoa(postId), nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockPostgres.AssertExpectations(t)

	expectedResponse := gin.H{
		"id":      float64(post.Id),
		"user_id": float64(post.UserId),
		"text":    post.Text,
		"image":   interface{}(nil),
	}
	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
