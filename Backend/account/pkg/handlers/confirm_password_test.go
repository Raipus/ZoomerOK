package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConfirmPasswordWithLogin(t *testing.T) {
	// Устанавливаем режим тестирования для Gin
	r := router.SetupRouter(false)

	// Создаем mock для кэширования
	mockCache := new(caching.MockCache)

	// Настраиваем тестовые данные
	resetLink := "someResetLink"
	login := "user1"

	// Тест для случая, когда username найден
	mockCache.On("GetCacheResetLink", resetLink).Return(login)

	// Регистрируем обработчик с использованием mockCache
	r.PUT("/confirm_password/:reset_link", func(c *gin.Context) {
		handlers.ConfirmPassword(c, mockCache)
	})

	// Создаем тестовый запрос
	req, _ := http.NewRequest(http.MethodPut, "/confirm_password/"+resetLink, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	mockCache.AssertExpectations(t)
}

func TestConfirmPasswordWithoutLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)

	// Создаем mock для кэширования
	mockCache := new(caching.MockCache)

	// Настраиваем тестовые данные
	resetLink := "someResetLink"
	mockCache.On("GetCacheResetLink", resetLink).Return("")

	// Регистрируем обработчик с использованием mockCache
	r.GET("/confirm_password/:reset_link", func(c *gin.Context) {
		handlers.ConfirmPassword(c, mockCache)
	})

	// Создаем тестовый запрос
	req, _ := http.NewRequest(http.MethodGet, "/confirm_password/"+resetLink, nil)

	// Выполняем запрос снова
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем статус-код для случая отсутствия username
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Проверяем, что ожидания выполнены
	mockCache.AssertExpectations(t)
}
