package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAcceptFriend(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

	var userId int = 1
	acceptFriendData := handlers.AcceptFriendForm{
		FriendUserId: 2,
	}

	redisUserFriend := memory.RedisUserFriend{
		UserId:    userId,
		FriendIds: []int{acceptFriendData.FriendUserId},
	}

	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(userId))
		c.Next()
	})

	mockPostgres.On("AcceptFriendRequest", userId, acceptFriendData.FriendUserId).Return(nil)
	mockRedis.On("AddUserFriend", redisUserFriend).Return(nil)

	jsonData, err := json.Marshal(acceptFriendData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	r.PUT("/accept_friend", func(c *gin.Context) {
		handlers.AcceptFriend(c, mockPostgres, mockRedis)
	})

	req, err := http.NewRequest(http.MethodPut, "/accept_friend", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	mockPostgres.AssertExpectations(t)
	mockRedis.AssertExpectations(t)
}
