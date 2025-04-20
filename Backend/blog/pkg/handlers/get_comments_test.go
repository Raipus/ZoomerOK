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

func TestGetComments(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockBroker := new(broker.MockBroker)

	r.GET("/post/:post_id/comments", func(c *gin.Context) {
		GetComments(c, mockPostgres, mockBroker)
	})

	commentTime := time.Now()
	var postId int = 1
	comment1 := postgres.Comment{
		Id:     4,
		UserId: 3,
		PostId: postId,
		Text:   "Комментарий 1",
		Time:   &commentTime,
	}
	comment2 := postgres.Comment{
		Id:     6,
		UserId: 4,
		PostId: postId,
		Text:   "Комментарий 2",
		Time:   &commentTime,
	}

	mockPostgres.On("GetComments", 1).Return([]postgres.Comment{comment1, comment2}, nil)

	req, err := http.NewRequest("GET", "/post/"+strconv.Itoa(postId)+"/comments", nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockPostgres.AssertExpectations(t)

	expectedResponse := gin.H{
		"comments": []interface{}{
			map[string]interface{}{
				"id":      float64(comment1.Id),
				"user_id": float64(comment1.UserId),
				"post_id": float64(comment1.PostId),
				"text":    comment1.Text,
			},
			map[string]interface{}{
				"id":      float64(comment2.Id),
				"user_id": float64(comment2.UserId),
				"post_id": float64(comment2.PostId),
				"text":    comment2.Text,
			},
		},
	}
	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
