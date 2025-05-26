package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

	userId := 1
	acceptFriendData := handlers.AcceptFriendForm{
		FriendUserId: 2,
	}

	t.Run("successful acceptance", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", float64(userId))
			c.Next()
		})
		r.PUT("/accept_friend", func(c *gin.Context) {
			handlers.AcceptFriend(c, mockPostgres, mockRedis)
		})
		mockPostgres.On("AcceptFriendRequest", userId, acceptFriendData.FriendUserId).Return(nil).Once()
		mockRedis.On("AddUserFriend", memory.RedisUserFriend{UserId: userId, FriendIds: []int{acceptFriendData.FriendUserId}}).Return(nil)
		mockRedis.On("AddUserFriend", memory.RedisUserFriend{UserId: acceptFriendData.FriendUserId, FriendIds: []int{userId}}).Return(nil)

		jsonData, _ := json.Marshal(acceptFriendData)
		req, _ := http.NewRequest(http.MethodPut, "/accept_friend", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
		mockPostgres.AssertExpectations(t)
		mockRedis.AssertExpectations(t)
	})

	t.Run("bad request - invalid JSON", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", float64(userId))
			c.Next()
		})
		r.PUT("/accept_friend", func(c *gin.Context) {
			handlers.AcceptFriend(c, mockPostgres, mockRedis)
		})
		req, _ := http.NewRequest(http.MethodPut, "/accept_friend", bytes.NewBuffer([]byte(`{"friend_user_id": "invalid"}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal server error - user ID not found", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", nil)
			c.Next()
		})
		r.PUT("/accept_friend", func(c *gin.Context) {
			handlers.AcceptFriend(c, mockPostgres, mockRedis)
		})

		jsonData, _ := json.Marshal(acceptFriendData)
		req, _ := http.NewRequest(http.MethodPut, "/accept_friend", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("internal server error - invalid user ID format", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", "invalid")
			c.Next()
		})
		r.PUT("/accept_friend", func(c *gin.Context) {
			handlers.AcceptFriend(c, mockPostgres, mockRedis)
		})

		jsonData, _ := json.Marshal(acceptFriendData)
		req, _ := http.NewRequest(http.MethodPut, "/accept_friend", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("bad request - friend request acceptance error", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", float64(userId))
			c.Next()
		})
		r.PUT("/accept_friend", func(c *gin.Context) {
			handlers.AcceptFriend(c, mockPostgres, mockRedis)
		})
		mockPostgres.On("AcceptFriendRequest", userId, acceptFriendData.FriendUserId).Return(fmt.Errorf("error")).Once()

		jsonData, _ := json.Marshal(acceptFriendData)
		req, _ := http.NewRequest(http.MethodPut, "/accept_friend", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
