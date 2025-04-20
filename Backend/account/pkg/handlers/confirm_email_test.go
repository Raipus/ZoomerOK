package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConfirmEmailWithLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockCache := new(caching.MockCache)
	mockPostgres := new(postgres.MockPostgres)

	confirmationLink := "someResetLink"
	login := "user2"

	// Тест для случая, когда username найден
	mockCache.On("GetCacheConfirmationLink", confirmationLink).Return(login)
	mockCache.On("DeleteCacheConfirmationLink", confirmationLink).Return()
	mockPostgres.On("ConfirmEmail", login).Return(true)

	// Регистрируем обработчик с использованием mockCache
	r.GET("/confirm_email/:confirmation_link", func(c *gin.Context) {
		ConfirmEmail(c, mockPostgres, mockCache)
	})

	// Создаем тестовый запрос
	req, _ := http.NewRequest(http.MethodGet, "/confirm_email/"+confirmationLink, nil)
	w := httptest.NewRecorder()

	// Выполняем запрос
	r.ServeHTTP(w, req)

	// Проверяем статус-код
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что ожидания выполнены
	mockCache.AssertExpectations(t)
	mockPostgres.AssertExpectations(t)
	mockPostgres.Calls = nil
	mockCache.Calls = nil
}

func TestConfirmEmailWithoutLogin(t *testing.T) {
	// Устанавливаем режим тестирования для Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Создаем mock для кэширования
	mockCache := new(caching.MockCache)
	mockPostgres := new(postgres.MockPostgres)

	// Настраиваем тестовые данные
	confirmationLink := "someResetLink"

	// Тест для случая, когда username не найден
	mockCache.On("GetCacheConfirmationLink", confirmationLink).Return("")

	// Регистрируем обработчик с использованием mockCache
	router.GET("/confirm_email/:confirmation_link", func(c *gin.Context) {
		ConfirmEmail(c, mockPostgres, mockCache)
	})

	// Создаем тестовый запрос
	req, _ := http.NewRequest(http.MethodGet, "/confirm_email/"+confirmationLink, nil)

	// Выполняем запрос снова
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверяем статус-код для случая отсутствия username
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Проверяем, что ожидания выполнены
	mockCache.AssertExpectations(t)
	mockCache.AssertCalled(t, "GetCacheConfirmationLink", confirmationLink)
	mockCache.Calls = nil
	mockPostgres.Calls = nil
}
