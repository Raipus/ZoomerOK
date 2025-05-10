package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/handlers"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserPosts(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)

	likeCountMap := make(map[int]int)
	commentCountMap := make(map[int]int)
	likeCountMap[1] = 5
	commentCountMap[1] = 6
	date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockPostgres.On("GetPosts", []int{1}, 1).Return([]postgres.Post{{Id: 1, UserId: 1, Text: "Пост 1", Image: []byte{}, Time: &date}}, nil)
	mockPostgres.On("GetCountCommentsAndLikes", []int{1}).Return(commentCountMap, likeCountMap, nil)
	r.GET("/user/:id/posts", func(c *gin.Context) {
		handlers.GetUserPosts(c, mockPostgres)
	})

	req, _ := http.NewRequest("GET", "/user/1/posts?page=1", nil)
	req.Header.Set("Authorization", "Bearer testtoken")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)

	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, actualResponse["posts"])

	postsInterface, ok := actualResponse["posts"].([]interface{})
	assert.True(t, ok)

	if len(postsInterface) > 0 {
		firstPost, ok := postsInterface[0].(map[string]interface{})
		assert.True(t, ok)

		body, ok := firstPost["body"].(map[string]interface{})
		assert.True(t, ok)

		assert.Equal(t, body["number_of_comments"], float64(6))
		assert.Equal(t, body["number_of_likes"], float64(5))
	}
}
