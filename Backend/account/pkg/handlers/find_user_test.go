package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFindUser(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)
	r.GET("/find_user", func(c *gin.Context) {
		handlers.FindUser(c, mockPostgres, mockRedis)
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
	redisUser := memory.RedisUser{
		UserId: 1,
		Login:  "testuser",
		Name:   "Тестовый Пользователь",
		Image:  "image",
	}

	mockPostgres.On("GetUserByLogin", user.Login).Return(user)
	mockRedis.On("GetUser", user.Id).Return(redisUser)

	findUserData := handlers.FindUserForm{
		Login: redisUser.Login,
	}

	jsonData, err := json.Marshal(findUserData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	req, _ := http.NewRequest("GET", "/find_user", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockPostgres.AssertExpectations(t)
	mockRedis.AssertExpectations(t)

	expectedResponse := gin.H{
		"id":    float64(user.Id),
		"login": user.Login,
		"name":  user.Name,
		"image": redisUser.Image,
	}
	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
