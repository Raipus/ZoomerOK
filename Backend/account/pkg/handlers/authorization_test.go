package handlers_test

import (
	"fmt"
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

func TestAuthMiddleware(t *testing.T) {
	mockPostgres := new(postgres.MockPostgres)

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

	mockPostgres.On("GetUserByLogin", user.Login).Return(user)

	t.Run("successful acceptance", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(handlers.AuthMiddleware(mockPostgres))
		r.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		strToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyQGV4YW1wbGUuY29tIiwiZXhwIjoxNzQ4NTE4MTUyLCJpZCI6MSwibG9naW4iOiJ0ZXN0dXNlciJ9.oDFERIARPY5JtbiyRomCIJJYhKfZLZ-5BQmNfDqf2zQ"
		req, err := http.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+strToken)
		if err != nil {
			t.Fatalf("Ошибка при создании запроса: %v", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "success"}`, w.Body.String())
	})

	t.Run("status unauthorized - Authorization header is required", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(handlers.AuthMiddleware(mockPostgres))
		r.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		req, err := http.NewRequest(http.MethodGet, "/protected", nil)
		if err != nil {
			t.Fatalf("Ошибка при создании запроса: %v", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Authorization header is required"}`, w.Body.String())
	})

	t.Run("status unauthorized - Invalid token", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(handlers.AuthMiddleware(mockPostgres))
		r.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		strToken := "token"
		req, err := http.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+strToken)
		if err != nil {
			t.Fatalf("Ошибка при создании запроса: %v", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Invalid token"}`, w.Body.String())
	})

	errors := map[string]string{
		"Invalid id claim":    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ZmFsc2UsImV4cCI6MTc0ODUxODI3MCwiaWQiOmZhbHNlLCJsb2dpbiI6ZmFsc2V9.YhrwPiJeXhLg49m8VKlAiLOearZuGmPWxzBXpuUslxc",
		"Invalid login claim": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ZmFsc2UsImV4cCI6MTc0ODUxODMzMCwiaWQiOjEsImxvZ2luIjpmYWxzZX0.Wl76FGofSSk2qrx4L6mQaCgEv1Gh5bGwFSocSBEPDps",
		"Invalid email claim": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ZmFsc2UsImV4cCI6MTc0ODUxODM1MCwiaWQiOjEsImxvZ2luIjoidGVzdHVzZXIifQ.fZX1CycGEYkXK1O8SIhUDuhLDI-d6tHHLeh3hilOKvo",
	}
	for key, token := range errors {
		t.Run("status unauthorized - "+key, func(t *testing.T) {
			r := router.SetupRouter(false)
			r.Use(handlers.AuthMiddleware(mockPostgres))
			r.GET("/protected", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			})

			req, err := http.NewRequest(http.MethodGet, "/protected", nil)
			req.Header.Set("Authorization", "Bearer "+token)
			if err != nil {
				t.Fatalf("Ошибка при создании запроса: %v", err)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusUnauthorized, w.Code)
			expected := fmt.Sprintf(`{"error": "%s"}`, key)
			actual := w.Body.String()

			assert.JSONEq(t, expected, actual)
		})
	}
}
