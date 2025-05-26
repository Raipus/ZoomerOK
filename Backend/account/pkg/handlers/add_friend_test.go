package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	mockPostgres := new(postgres.MockPostgres)

	userId := 1
	addFriendData := handlers.AddFriendForm{
		FriendUserId: 2,
	}

	t.Run("successful acceptance", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", float64(userId))
			c.Next()
		})
		r.PUT("/add_friend", func(c *gin.Context) {
			handlers.AddFriend(c, mockPostgres)
		})
		mockPostgres.On("AddFriendRequest", userId, addFriendData.FriendUserId).Return(nil).Once()

		jsonData, err := json.Marshal(addFriendData)
		if err != nil {
			t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, "/add_friend", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Ошибка при создании запроса: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)

		mockPostgres.AssertExpectations(t)
	})

	t.Run("bad request - invalid JSON", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", float64(userId))
			c.Next()
		})
		r.PUT("/add_friend", func(c *gin.Context) {
			handlers.AddFriend(c, mockPostgres)
		})
		req, _ := http.NewRequest(http.MethodPut, "/add_friend", bytes.NewBuffer([]byte(`{"friend_user_id": "invalid"}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal server error - user ID not found", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Next()
		})
		r.PUT("/add_friend", func(c *gin.Context) {
			handlers.AddFriend(c, mockPostgres)
		})

		jsonData, _ := json.Marshal(addFriendData)
		req, _ := http.NewRequest(http.MethodPut, "/add_friend", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "User ID not found in context"}`, w.Body.String())
	})

	t.Run("internal server error - invalid user ID format", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", "invalid")
			c.Next()
		})
		r.PUT("/add_friend", func(c *gin.Context) {
			handlers.AddFriend(c, mockPostgres)
		})

		jsonData, _ := json.Marshal(addFriendData)
		req, _ := http.NewRequest(http.MethodPut, "/add_friend", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "Invalid user ID format"}`, w.Body.String())
	})

	t.Run("internal server error - Ошибка сервиса", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("user_id", float64(userId))
			c.Next()
		})
		r.PUT("/add_friend", func(c *gin.Context) {
			handlers.AddFriend(c, mockPostgres)
		})
		mockPostgres.On("AddFriendRequest", userId, addFriendData.FriendUserId).Return(fmt.Errorf("error")).Once()

		jsonData, _ := json.Marshal(addFriendData)
		req, _ := http.NewRequest(http.MethodPut, "/add_friend", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "Ошибка сервиса"}`, w.Body.String())
	})
}
