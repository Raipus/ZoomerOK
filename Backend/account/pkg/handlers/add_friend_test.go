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

	acceptFriendData := handlers.AddFriendForm{
		UserId:       1,
		FriendUserId: 2,
	}

	mockPostgres.On("AddFriendRequest", acceptFriendData.UserId, acceptFriendData.FriendUserId).Return(nil)

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
