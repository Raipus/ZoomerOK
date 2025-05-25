package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestGetPosts(t *testing.T) {
	r := router.SetupRouter(false)
	userId := 3
	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(userId))
		c.Next()
	})
	mockPostgres := new(postgres.MockPostgres)
	mockBroker := new(broker.MockBroker)
	mockMessageStore := new(memory.MockMessageStore)
	r.GET("/posts", func(c *gin.Context) {
		handlers.GetPosts(c, mockPostgres, mockBroker, mockMessageStore)
	})

	likeCountMap := make(map[int]int)
	commentCountMap := make(map[int]int)
	likeCountMap[1] = 5
	commentCountMap[1] = 6
	date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockBroker.On("PushUserFriend", mock.Anything).Return(nil)
	mockMessageStore.On("ProcessPushUserFriend", mock.Anything).Return(pb.GetUserFriendResponse{Ids: []int64{1, 2}, Id: 1}, nil).Once()
	mockPostgres.On("GetPosts", []int{1, 2, 3}, 1).Return([]postgres.Post{{Id: 1, UserId: 1, Text: "Пост 1", Image: []byte{}, Time: &date}}, nil)
	mockBroker.On("PushUsers", mock.Anything).Return(nil)
	mockPostgres.On("GetCountCommentsAndLikes", []int{1}).Return(commentCountMap, likeCountMap, nil)
	mockMessageStore.On("ProcessPushUsers", mock.Anything).Return(pb.GetUsersResponse{
		Users: []*pb.GetUserResponse{
			&pb.GetUserResponse{
				Image: "",
				Name:  "username1",
				Login: "testuser1",
				Id:    1,
			},
			&pb.GetUserResponse{
				Image: "",
				Name:  "username2",
				Login: "testuser2",
				Id:    2,
			},
		},
		Ids: []int64{1, 2},
	}, nil).Once()

	req, _ := http.NewRequest("GET", "/posts?page=1", nil)
	req.Header.Set("Authorization", "Bearer testtoken")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)
	mockBroker.AssertExpectations(t)
	mockMessageStore.AssertExpectations(t)

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
