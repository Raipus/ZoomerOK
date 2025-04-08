package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConfirmPassword(t *testing.T) {
	// Устанавливаем режим тестирования для Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Создаем mock для кэширования
	mockCache := new(caching.MockCache)

	// Настраиваем тестовые данные
	resetLink := "someResetLink"
	username := "user1"

	// Тест для случая, когда username найден
	mockCache.On("GetCacheResetLink", resetLink).Return(username)
	mockCache.On("DeleteCacheResetLink", resetLink).Return()

	// Регистрируем обработчик с использованием mockCache
	router.GET("/confirm/:reset_link", func(c *gin.Context) {
		ConfirmPassword(c, mockCache)
	})

	// Создаем тестовый запрос
	req, _ := http.NewRequest(http.MethodGet, "/confirm/"+resetLink, nil)
	w := httptest.NewRecorder()

	// Выполняем запрос
	router.ServeHTTP(w, req)

	// Проверяем статус-код
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что ожидания выполнены
	mockCache.AssertExpectations(t)

	// Тест для случая, когда username не найден
	mockCache.On("GetCacheResetLink", resetLink).Return("")

	// Выполняем запрос снова
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверяем статус-код для случая отсутствия username
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Проверяем, что ожидания выполнены
	mockCache.AssertExpectations(t)
}
