package handlers_test

import (
	"bytes"
	"encoding/json"
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
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)

	changePasswordData := handlers.ChangePasswordForm{
		NewPassword: "newsecurepassword",
	}

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
	resetLink := "someResetLink"
	mockPostgres.On("GetUserByLogin", login).Return(user)
	mockPostgres.On("ChangePassword", &user, changePasswordData.NewPassword).Return(nil)
	mockCache.On("GetCacheResetLink", resetLink).Return(login)
	mockCache.On("DeleteCacheResetLink", resetLink)

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
	assert.Equal(t, http.StatusNoContent, w.Code)

	mockPostgres.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func TestChangePasswordWithNoLogin(t *testing.T) {
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)

	changePasswordData := handlers.ChangePasswordForm{
		NewPassword: "newsecurepassword",
	}

	resetLink := "reset"
	mockCache.On("GetCacheResetLink", resetLink).Return("")

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

	mockPostgres.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}
