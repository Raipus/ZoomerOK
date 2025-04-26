package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(false)
	var token string = "my_token"
	r.Use(func(c *gin.Context) {
		c.Set("token", token)
		c.Next()
	})
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

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

	mockPostgres.On("GetUserByLogin", login).Return(user)
	mockPostgres.On("DeleteUser", &user).Return()
	mockRedis.On("DeleteAuthorization", token).Return()
	mockRedis.On("DeleteUser", user.Id).Return()
	mockRedis.On("DeleteAllUserFriend", user.Id).Return()

	r.DELETE("/user/:login", func(c *gin.Context) {
		handlers.DeleteUser(c, mockPostgres, mockRedis)
	})

	req, err := http.NewRequest("DELETE", "/user/"+login, nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
