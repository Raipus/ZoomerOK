package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConfirmEmailWithLogin(t *testing.T) {
	r := router.SetupRouter(false)
	mockCache := new(caching.MockCache)
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

	confirmationLink := "someResetLink"
	login := "testuser"
	var token string = "new-token"

	r.Use(func(c *gin.Context) {
		c.Set("token", token)
		c.Next()
	})
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
	redisAuthorization := memory.RedisAuthorization{
		UserId:         user.Id,
		Token:          token,
		Login:          user.Login,
		Email:          user.Email,
		ConfirmedEmail: true,
	}
	// Тест для случая, когда username найден
	mockCache.On("GetCacheConfirmationLink", confirmationLink).Return(login)
	mockCache.On("DeleteCacheConfirmationLink", confirmationLink).Return()
	mockPostgres.On("ConfirmEmail", login).Return(user, true)
	mockRedis.On("SetAuthorization", redisAuthorization).Return()

	// Регистрируем обработчик с использованием mockCache
	r.PUT("/confirm_email/:confirmation_link", func(c *gin.Context) {
		handlers.ConfirmEmail(c, mockPostgres, mockCache, mockRedis)
	})

	// Создаем тестовый запрос
	req, _ := http.NewRequest(http.MethodPut, "/confirm_email/"+confirmationLink, nil)
	w := httptest.NewRecorder()

	// Выполняем запрос
	r.ServeHTTP(w, req)

	// Проверяем статус-код
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Проверяем, что ожидания выполнены
	mockCache.AssertExpectations(t)
	mockPostgres.AssertExpectations(t)
	mockRedis.AssertExpectations(t)
}

func TestConfirmEmailWithoutLogin(t *testing.T) {
	// Устанавливаем режим тестирования для Gin
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)

	// Создаем mock для кэширования
	mockCache := new(caching.MockCache)
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

	// Настраиваем тестовые данные
	confirmationLink := "someResetLink"

	// Тест для случая, когда username не найден
	mockCache.On("GetCacheConfirmationLink", confirmationLink).Return("")

	// Регистрируем обработчик с использованием mockCache
	r.GET("/confirm_email/:confirmation_link", func(c *gin.Context) {
		handlers.ConfirmEmail(c, mockPostgres, mockCache, mockRedis)
	})

	// Создаем тестовый запрос
	req, _ := http.NewRequest(http.MethodGet, "/confirm_email/"+confirmationLink, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Проверяем, что ожидания выполнены
	mockCache.AssertExpectations(t)
	mockCache.AssertCalled(t, "GetCacheConfirmationLink", confirmationLink)
}
