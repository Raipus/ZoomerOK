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
	mockMessageQueue := new(memory.MockMessageQueue)
	r.GET("/post/:post_id", func(c *gin.Context) {
		handlers.GetPost(c, mockPostgres, mockBroker, mockMessageQueue)
	})

	postId := 123
	date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockPostgres.On("GetPost", postId).Return(&postgres.Post{Id: postId, Text: "Тестовый пост", Image: []byte{}, Time: &date}, nil)
	mockBroker.On("PushUser", mock.Anything).Return(nil)
	mockMessageQueue.On("GetLastMessage").Return(&pb.GetUserResponse{Id: 1, Login: "testuser", Name: "Тест", Image: ""})

	req, _ := http.NewRequest("GET", "/post/"+strconv.Itoa(postId), nil)
	req.Header.Set("Authorization", "Bearer testtoken")
	req = req.WithContext(context.WithValue(req.Context(), "user_id", float64(1)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)
	mockBroker.AssertExpectations(t)
	mockMessageQueue.AssertExpectations(t)

	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, actualResponse["post"])
}
