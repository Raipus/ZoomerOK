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
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)
	login := "testuser"
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
	redisUser := memory.RedisUser{
		UserId: user.Id,
		Login:  user.Login,
		Name:   user.Name,
		Image:  config.Config.Photo.Base64Small,
	}
	changeUserData := handlers.ChangeUserForm{
		Name:     user.Name,
		Birthday: user.Birthday,
		Phone:    user.Phone,
		City:     user.City,
		Image:    user.Image,
	}
	t.Run("successful acceptance", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.PUT("/user/:login", func(c *gin.Context) {
			handlers.ChangeUser(c, mockPostgres, mockRedis)
		})

		mockPostgres.On("GetUserByLogin", login).Return(user).Once()
		mockPostgres.On("ChangeUser", &user).Return(true).Once()
		mockRedis.On("SetUser", redisUser).Return().Once()

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
	})
	t.Run("bad request - invalid JSON", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.PUT("/user/:login", func(c *gin.Context) {
			handlers.ChangeUser(c, mockPostgres, mockRedis)
		})
		req, _ := http.NewRequest(http.MethodPut, "/user/"+login, bytes.NewBuffer([]byte(`{"name": false}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("bad request - Get User By Login", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.PUT("/user/:login", func(c *gin.Context) {
			handlers.ChangeUser(c, mockPostgres, mockRedis)
		})

		mockPostgres.On("GetUserByLogin", login).Return(postgres.User{}).Once()

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

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error": "пользователь не найден!"}`, w.Body.String())

		mockPostgres.AssertExpectations(t)
		mockRedis.AssertExpectations(t)
	})
	t.Run("internal server error - Change Password", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.PUT("/user/:login", func(c *gin.Context) {
			handlers.ChangeUser(c, mockPostgres, mockRedis)
		})

		mockPostgres.On("GetUserByLogin", login).Return(user).Once()
		mockPostgres.On("ChangeUser", &user).Return(false).Once()

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

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "Ошибка сервиса"}`, w.Body.String())

		mockPostgres.AssertExpectations(t)
		mockRedis.AssertExpectations(t)
	})
}
