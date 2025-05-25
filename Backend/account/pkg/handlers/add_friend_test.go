package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddFriend(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)

	userId := 1
	acceptFriendData := handlers.AddFriendForm{
		FriendUserId: 2,
	}

	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(userId))
		c.Next()
	})
	mockPostgres.On("AddFriendRequest", userId, acceptFriendData.FriendUserId).Return(nil)

	jsonData, err := json.Marshal(acceptFriendData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	r.POST("/add_friend", func(c *gin.Context) {
		handlers.AddFriend(c, mockPostgres)
	})

	req, err := http.NewRequest(http.MethodPost, "/add_friend", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	mockPostgres.AssertExpectations(t)
}
