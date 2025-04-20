package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAcceptFriend(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)

	acceptFriendData := AcceptFriendForm{
		UserId:       1,
		FriendUserId: 2,
	}

	mockPostgres.On("AcceptFriendRequest", acceptFriendData.UserId, acceptFriendData.FriendUserId).Return(nil)

	jsonData, err := json.Marshal(acceptFriendData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	r.PUT("/accept_friend", func(c *gin.Context) {
		AcceptFriend(c, mockPostgres)
	})

	req, err := http.NewRequest(http.MethodPut, "/accept_friend", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockPostgres.AssertExpectations(t)
}
