package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestWantChangePassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockSmtp := new(security.MockSmtp)
	mockPostgres := new(postgres.MockPostgres)
	mockCache := new(caching.MockCache)

	wantChangePasswordData := WantChangePasswordForm{
		Email: "test@example.com",
	}

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

	mockPostgres.On("GetUserByEmail", wantChangePasswordData.Email).Return(user)
	mockSmtp.On("SendChangePassword", user.Name, user.Email, mockCache).Return(nil)
	r.POST("/want_change_password", func(c *gin.Context) {
		WantChangePassword(c, mockPostgres, mockSmtp, mockCache)
	})

	jsonData, err := json.Marshal(wantChangePasswordData)
	if err != nil {
		t.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/want_change_password", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockPostgres.AssertExpectations(t)
	mockSmtp.AssertExpectations(t)
}
