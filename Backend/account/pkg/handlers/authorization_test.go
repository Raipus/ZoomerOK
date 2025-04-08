package security

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Создаем маршрут с использованием AuthMiddleware
	router.GET("/protected", security.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You are authorized"})
	})

	// Тест без токена
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Генерация тестового токена
	user := security.UserToken{Id: 1, Name: "Test User", Email: "test@example.com", ConfirmedEmail: true, Image: ""}
	token, err := GenerateJWT(user)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Тест с валидным токеном
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "You are authorized"}`, w.Body.String())

	// Тест с недействительным токеном
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
