package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestChangeUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	mockPostgres := new(postgres.MockPostgres)

	var login string = "testuser"
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

	r.PUT("/user/:login", func(c *gin.Context) {
		ChangeUser(c, mockPostgres)
	})

	mockPostgres.On("GetUserByLogin", login).Return(user)
	mockPostgres.On("ChangeUser", &user).Return(true)

	changeUserData := ChangeUserForm{
		Name:     "Обновленный Пользователь",
		Birthday: &birthday,
		Phone:    "987-654-3210",
		City:     "Санкт-Петербург",
		Image:    nil,
	}

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

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Пользователь успешно обновлен"}`, w.Body.String())
}
