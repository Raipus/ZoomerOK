package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockBroker := new(broker.MockBroker)

	r.GET("/posts", func(c *gin.Context) {
		GetPosts(c, mockPostgres, mockBroker)
	})

	postTime := time.Now()
	post1 := postgres.Post{
		Id:     4,
		UserId: 3,
		Text:   "Пост 1",
		Image:  nil,
		Time:   &postTime,
	}
	post2 := postgres.Post{
		Id:     6,
		UserId: 4,
		Text:   "Пост 2",
		Image:  nil,
		Time:   &postTime,
	}

	mockPostgres.On("GetPosts", 1).Return([]postgres.Post{post1, post2}, nil)

	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockPostgres.AssertExpectations(t)

	expectedResponse := gin.H{
		"posts": []interface{}{
			map[string]interface{}{
				"id":      float64(post1.Id),
				"user_id": float64(post1.UserId),
				"text":    post1.Text,
				"image":   interface{}(nil),
			},
			map[string]interface{}{
				"id":      float64(post2.Id),
				"user_id": float64(post2.UserId),
				"text":    post2.Text,
				"image":   interface{}(nil),
			},
		},
	}
	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
