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

func TestGetUser(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	r.GET("/user/:login", func(c *gin.Context) {
		handlers.GetUser(c, mockPostgres)
	})

	var login string = "testuser"
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

	mockPostgres.On("GetUserByLogin", login).Return(user)

	req, _ := http.NewRequest("GET", "/user/"+login, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)

	expectedResponse := gin.H{
		"id":       float64(user.Id),
		"login":    user.Login,
		"name":     user.Name,
		"email":    user.Email,
		"birthday": user.Birthday.Format(time.RFC3339Nano),
		"phone":    user.Phone,
		"city":     user.City,
		"image":    interface{}(nil),
	}
	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
