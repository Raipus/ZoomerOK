package handlers_test

import (
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

func TestGetComments(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockBroker := new(broker.MockBroker)
	mockMessageStore := new(memory.MockMessageStore)
	r.GET("/post/:post_id/comments", func(c *gin.Context) {
		handlers.GetComments(c, mockPostgres, mockBroker, mockMessageStore)
	})

	postId := 789
	date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockPostgres.On("GetComments", postId, 1).Return([]postgres.Comment{{Id: 1, PostId: postId, UserId: 1, Text: "Комментарий", Time: &date}}, nil)
	mockBroker.On("PushUsers", mock.Anything).Return(nil)
	mockMessageStore.On("ProcessPushUsers", mock.Anything).Return(&pb.GetUsersResponse{Users: []*pb.GetUserResponse{{Id: 1, Login: "testuser", Name: "Тест", Image: ""}}, Ids: []int64{1}}, nil)

	req, _ := http.NewRequest("GET", "/post/"+strconv.Itoa(postId)+"/comments?page=1", nil)
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
	assert.NotEmpty(t, actualResponse["comments"])
}
