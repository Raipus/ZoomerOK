package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetFriends(t *testing.T) {
	r := router.SetupRouter(false)
	mockRedis := new(memory.MockRedis)
	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(1))
		c.Next()
	})

	r.GET("/get_friends", func(c *gin.Context) {
		handlers.GetFriends(c, mockRedis)
	})

	redisUserFriend := memory.RedisUserFriend{
		UserId:    1,
		FriendIds: []int{1, 2},
	}
	redisUsers := []memory.RedisUser{
		{
			UserId: 1,
			Login:  "testuser",
			Name:   "Тестовый Пользователь",
			Image:  "image",
		},
		{
			UserId: 2,
			Login:  "testuser1",
			Name:   "Тестовый Пользователь1",
			Image:  "image1",
		},
	}

	mockRedis.On("GetUsers", redisUserFriend.FriendIds).Return(redisUsers)
	mockRedis.On("GetUserFriends", 1).Return(redisUserFriend)

	req, _ := http.NewRequest("GET", "/get_friends", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockRedis.AssertExpectations(t)

	expectedResponse := gin.H{
		"users": []interface{}{
			map[string]interface{}{
				"friend": map[string]interface{}{
					"id":    float64(1),
					"image": "image",
					"login": "testuser",
					"name":  "Тестовый Пользователь",
				},
			},
			map[string]interface{}{
				"friend": map[string]interface{}{
					"id":    float64(2),
					"image": "image1",
					"login": "testuser1",
					"name":  "Тестовый Пользователь1",
				},
			},
		},
	}
	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
