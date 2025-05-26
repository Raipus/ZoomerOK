package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestChangePassword(t *testing.T) {
	resetLink := "someResetLink"
	birthday := time.Now()
	login := "testuser"
	user := postgres.User{
		Id:             1,
		Login:          login,
		Name:           "Тестовый Пользователь",
		Email:          "testuser@example.com",
		ConfirmedEmail: true,
		Password:       "securepassword",
		Birthday:       &birthday,
		Phone:          "123-456-7890",
		City:           "Москва",
		Image:          nil,
	}
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)

	t.Run("successful acceptance", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.PUT("/change_password/:reset_link", func(c *gin.Context) {
			handlers.ChangePassword(c, mockPostgres, mockCache)
		})

		changePasswordData := handlers.ChangePasswordForm{
			NewPassword: "newsecurepassword",
		}
		mockPostgres.On("GetUserByLogin", login).Return(user).Once()
		mockPostgres.On("ChangePassword", &user, changePasswordData.NewPassword).Return(nil).Once()
		mockCache.On("GetCacheResetLink", resetLink).Return(login).Once()
		mockCache.On("DeleteCacheResetLink", resetLink).Once()

		jsonData, err := json.Marshal(changePasswordData)
		if err != nil {
			t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
		}

		req, err := http.NewRequest("PUT", "/change_password/"+resetLink, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Ошибка при создании запроса: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)

		mockPostgres.AssertExpectations(t)
		mockCache.AssertExpectations(t)
	})
	t.Run("bad request - invalid JSON", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.PUT("/change_password/:reset_link", func(c *gin.Context) {
			handlers.ChangePassword(c, mockPostgres, mockCache)
		})

		req, _ := http.NewRequest(http.MethodPut, "/change_password/:reset_link", bytes.NewBuffer([]byte(`{"password": false}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("not found - User not found", func(t *testing.T) {
		r := router.SetupRouter(false)
		mockPostgres := new(postgres.MockPostgres)
		mockCache := new(caching.MockCache)

		changePasswordData := handlers.ChangePasswordForm{
			NewPassword: "newsecurepassword",
		}

		resetLink := "reset"
		mockCache.On("GetCacheResetLink", resetLink).Return("").Once()

		jsonData, err := json.Marshal(changePasswordData)
		if err != nil {
			t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
		}

		r.PUT("/change_password/:reset_link", func(c *gin.Context) {
			handlers.ChangePassword(c, mockPostgres, mockCache)
		})

		req, err := http.NewRequest("PUT", "/change_password/"+resetLink, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Ошибка при создании запроса: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"error": "User not found"}`, w.Body.String())

		mockPostgres.AssertExpectations(t)
		mockCache.AssertExpectations(t)
	})
	t.Run("bad request - Get User By Login", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.PUT("/change_password/:reset_link", func(c *gin.Context) {
			handlers.ChangePassword(c, mockPostgres, mockCache)
		})

		changePasswordData := handlers.ChangePasswordForm{
			NewPassword: "newsecurepassword",
		}
		mockPostgres.On("GetUserByLogin", login).Return(postgres.User{}).Once()
		mockCache.On("GetCacheResetLink", resetLink).Return(login).Once()

		jsonData, err := json.Marshal(changePasswordData)
		if err != nil {
			t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
		}

		req, err := http.NewRequest("PUT", "/change_password/"+resetLink, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Ошибка при создании запроса: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error": "Пользователь не найден"}`, w.Body.String())

		mockPostgres.AssertExpectations(t)
		mockCache.AssertExpectations(t)
	})
	t.Run("bad request - Change Password", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.PUT("/change_password/:reset_link", func(c *gin.Context) {
			handlers.ChangePassword(c, mockPostgres, mockCache)
		})

		changePasswordData := handlers.ChangePasswordForm{
			NewPassword: "newsecurepassword",
		}
		mockPostgres.On("GetUserByLogin", login).Return(user).Once()
		mockPostgres.On("ChangePassword", &user, changePasswordData.NewPassword).Return(fmt.Errorf("error")).Once()
		mockCache.On("GetCacheResetLink", resetLink).Return(login).Once()

		jsonData, err := json.Marshal(changePasswordData)
		if err != nil {
			t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
		}

		req, err := http.NewRequest("PUT", "/change_password/"+resetLink, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatalf("Ошибка при создании запроса: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "Ошибка сервера"}`, w.Body.String())

		mockPostgres.AssertExpectations(t)
		mockCache.AssertExpectations(t)
	})
}
