package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

	loginData := handlers.LoginForm{
		LoginOrEmail: "testuser",
		Password:     "password",
	}

	birthday := time.Now()
	byteImage := config.Config.Photo.ByteImage
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
		Image:          byteImage,
	}

	token := "new-token"
	redisAuthorization := memory.RedisAuthorization{
		UserId:         user.Id,
		Token:          token,
		Login:          user.Login,
		Email:          user.Email,
		ConfirmedEmail: user.ConfirmedEmail,
	}
	redisUser := memory.RedisUser{
		UserId: user.Id,
		Login:  user.Login,
		Name:   user.Name,
		Image:  "image",
	}

	mockPostgres.On("Login", loginData.LoginOrEmail, loginData.Password).Return(user, token, "")
	mockRedis.On("SetAuthorization", redisAuthorization).Return()
	mockRedis.On("GetUser", user.Id).Return(redisUser)

	r.POST("/login", func(c *gin.Context) {
		handlers.Login(c, mockPostgres, mockRedis)
	})

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, token, actualResponse["token"])
	assert.Equal(t, redisUser.Name, actualResponse["name"])
	assert.Equal(t, redisUser.Image, actualResponse["image"])
	mockPostgres.AssertExpectations(t)
	mockRedis.AssertExpectations(t)
}
