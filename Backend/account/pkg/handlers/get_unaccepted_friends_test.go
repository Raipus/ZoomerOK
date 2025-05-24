package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUnacceptedFriends(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(1))
		c.Next()
	})

	r.GET("/get_unaccepted_friends", func(c *gin.Context) {
		handlers.GetUnacceptedFriends(c, mockPostgres)
	})

	birthday := time.Now()
	user := postgres.User{
		Id:             1,
		Login:          "testuser",
		Name:           "Тестовый Пользователь",
		Email:          "testuser@example.com",
		ConfirmedEmail: true,
		Password:       "securepassword",
		Birthday:       &birthday,
		Phone:          "123-456-7890",
		City:           "Москва",
		Image:          nil,
	}

	mockPostgres.On("GetUnacceptedFriends", 1).Return([]postgres.User{user}, nil)

	req, _ := http.NewRequest("GET", "/get_unaccepted_friends", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)

	expectedResponse := gin.H{
		"unaccepted_friends": []interface{}{
			map[string]interface{}{
				"unaccepted_friend": map[string]interface{}{
					"id":    float64(1),
					"image": nil,
					"login": "testuser",
					"name":  "Тестовый Пользователь",
				},
			},
		},
	}
	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetEmptyUnacceptedFriends(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	r.Use(func(c *gin.Context) {
		c.Set("user_id", float64(1))
		c.Next()
	})

	r.GET("/get_unaccepted_friends", func(c *gin.Context) {
		handlers.GetUnacceptedFriends(c, mockPostgres)
	})

	mockPostgres.On("GetUnacceptedFriends", 1).Return([]postgres.User{}, nil)

	req, _ := http.NewRequest("GET", "/get_unaccepted_friends", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)

	expectedResponse := gin.H{
		"unaccepted_friends": nil,
	}
	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
