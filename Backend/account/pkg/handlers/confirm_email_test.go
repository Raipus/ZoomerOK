package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/handlers"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConfirmEmailWithLogin(t *testing.T) {
	mockCache := new(caching.MockCache)
	mockPostgres := new(postgres.MockPostgres)
	mockRedis := new(memory.MockRedis)

	confirmationLink := "someResetLink"
	login := "testuser"
	token := "new-token"

	birthday := time.Now()
	byteImage := config.Config.Photo.ByteImage
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
		Image:          byteImage,
	}
	redisAuthorization := memory.RedisAuthorization{
		UserId:         user.Id,
		Token:          token,
		Login:          user.Login,
		Email:          user.Email,
		ConfirmedEmail: true,
	}
	t.Run("successful acceptance", func(t *testing.T) {
		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("token", token)
			c.Next()
		})
		r.PUT("/confirm_email/:confirmation_link", func(c *gin.Context) {
			handlers.ConfirmEmail(c, mockPostgres, mockCache, mockRedis)
		})

		mockCache.On("GetCacheConfirmationLink", confirmationLink).Return(login).Once()
		mockCache.On("DeleteCacheConfirmationLink", confirmationLink).Return().Once()
		mockPostgres.On("ConfirmEmail", login).Return(user, true).Once()
		mockRedis.On("SetAuthorization", redisAuthorization).Return().Once()

		req, _ := http.NewRequest(http.MethodPut, "/confirm_email/"+confirmationLink, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)

		mockCache.AssertExpectations(t)
		mockPostgres.AssertExpectations(t)
		mockRedis.AssertExpectations(t)
	})
	t.Run("not found - User Not found", func(t *testing.T) {
		mockCache.On("GetCacheConfirmationLink", confirmationLink).Return("").Once()

		r := router.SetupRouter(false)
		r.Use(func(c *gin.Context) {
			c.Set("token", token)
			c.Next()
		})
		r.PUT("/confirm_email/:confirmation_link", func(c *gin.Context) {
			handlers.ConfirmEmail(c, mockPostgres, mockCache, mockRedis)
		})

		req, _ := http.NewRequest(http.MethodPut, "/confirm_email/"+confirmationLink, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)

		mockCache.AssertExpectations(t)
		mockCache.AssertCalled(t, "GetCacheConfirmationLink", confirmationLink)
	})
}
