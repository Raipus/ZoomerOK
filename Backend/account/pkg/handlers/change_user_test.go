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

func TestChangeUser(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

	var login string = "testuser"
	birthday := time.Date(2025, 5, 6, 21, 50, 36, 113918233, time.UTC)
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

	r.PUT("/user/:login", func(c *gin.Context) {
		handlers.ChangeUser(c, mockPostgres, mockRedis)
	})

	// log.Println(config.Config.Photo.Base64Small)
	redisUser := memory.RedisUser{
		UserId: user.Id,
		Login:  user.Login,
		Name:   user.Name,
		Image:  config.Config.Photo.Base64Small,
	}

	mockPostgres.On("GetUserByLogin", login).Return(user)
	mockPostgres.On("ChangeUser", &user).Return(true)
	mockRedis.On("SetUser", redisUser).Return()

	changeUserData := handlers.ChangeUserForm{
		Name:     user.Name,
		Birthday: user.Birthday,
		Phone:    user.Phone,
		City:     user.City,
		Image:    user.Image,
	}

	jsonData, err := json.Marshal(changeUserData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	req, err := http.NewRequest("PUT", "/user/"+login, bytes.NewBuffer(jsonData))
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
