package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/broker"
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/Raipus/ZoomerOK/blog/pkg/handlers"
	"github.com/Raipus/ZoomerOK/blog/pkg/memory"
	"github.com/Raipus/ZoomerOK/blog/pkg/postgres"
	"github.com/Raipus/ZoomerOK/blog/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPost(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockBroker := new(broker.MockBroker)
	mockMessageStore := new(memory.MockMessageStore)
	r.GET("/post/:post_id", func(c *gin.Context) {
		handlers.GetPost(c, mockPostgres, mockBroker, mockMessageStore)
	})

	likeCountMap := make(map[int]int)
	commentCountMap := make(map[int]int)
	postId := 123
	likeCountMap[postId] = 5
	commentCountMap[postId] = 6
	date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockPostgres.On("GetPost", postId).Return(&postgres.Post{Id: postId, Text: "Тестовый пост", Image: []byte{}, Time: &date}, nil)
	mockBroker.On("PushUser", mock.Anything).Return(nil)
	mockMessageStore.On("ProcessPushUser", mock.Anything).Return(pb.GetUserResponse{Id: 1, Login: "testuser", Name: "Тест", Image: ""}, nil)
	mockPostgres.On("GetCountCommentsAndLikes", []int{postId}).Return(commentCountMap, likeCountMap, nil)

	req, _ := http.NewRequest("GET", "/post/"+strconv.Itoa(postId), nil)
	req.Header.Set("Authorization", "Bearer testtoken")
	req = req.WithContext(context.WithValue(req.Context(), "user_id", float64(1)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)
	mockBroker.AssertExpectations(t)
	mockMessageStore.AssertExpectations(t)

	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, actualResponse["post"])
	assert.Equal(t, actualResponse["post"].(map[string]interface{})["body"].(map[string]interface{})["number_of_comments"], float64(6))
	assert.Equal(t, actualResponse["post"].(map[string]interface{})["body"].(map[string]interface{})["number_of_likes"], float64(5))
}
