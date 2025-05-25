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

func TestDeleteFriend(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

	userId := 1
	deleteFriendData := handlers.DeleteFriendForm{
		FriendUserId: 2,
	}

	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(userId))
		c.Next()
	})

	mockPostgres.On("DeleteFriendRequest", userId, deleteFriendData.FriendUserId).Return(nil)
	mockRedis.On("DeleteUserFriend", userId, deleteFriendData.FriendUserId).Return()
	mockRedis.On("DeleteUserFriend", deleteFriendData.FriendUserId, userId).Return()

	jsonData, err := json.Marshal(deleteFriendData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	r.DELETE("/delete_friend", func(c *gin.Context) {
		handlers.DeleteFriend(c, mockPostgres, mockRedis)
	})

	req, err := http.NewRequest(http.MethodDelete, "/delete_friend", bytes.NewBuffer(jsonData))
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
