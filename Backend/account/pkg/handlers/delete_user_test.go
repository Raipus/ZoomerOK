package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
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

	r.DELETE("/user/:login", func(c *gin.Context) {
		DeleteUser(c, mockPostgres)
	})

	mockPostgres.On("GetUserByLogin", login).Return(user)
	mockPostgres.On("DeleteUser", &user).Return()

	req, err := http.NewRequest("DELETE", "/user/"+login, nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
