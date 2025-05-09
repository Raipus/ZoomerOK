package handlers

import (
	"fmt"
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
)

// SignupForm представляет данные, необходимые для регистрации пользователя.
type SignupForm struct {
	Login    string `json:"login"`    // Логин пользователя
	Name     string `json:"name"`     // Имя пользователя
	Email    string `json:"email"`    // Электронная почта пользователя
	Password string `json:"password"` // Пароль пользователя
}

// Signup регистрирует нового пользователя.
// @Summary Регистрация пользователя
// @Description Позволяет зарегистрировать нового пользователя и отправить подтверждение по электронной почте.
// @Accept json
// @Produce json
// @Param user body SignupForm true "Данные для регистрации пользователя"
// @Success 200 {object} gin.H {"token": "JWT токен"}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 409 {object} gin.H {"error": "Логин или электронная почта уже существуют"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /signup [post]
func Signup(c *gin.Context, db postgres.PostgresInterface, smtp security.SMTPInterface, cache caching.CachingInterface, redis memory.RedisInterface) {
	var newSignupForm SignupForm
	if err := c.BindJSON(&newSignupForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	// Проверяем логин
	loginUser := db.GetUserByLogin(newSignupForm.Login)
	if loginUser.Login != "" {
		c.JSON(http.StatusConflict, gin.H{
			"error": fmt.Sprintf("Логин '%s' уже существует", newSignupForm.Login),
		})
		return
	}

	// Проверяем электронную почту
	emailUser := db.GetUserByEmail(newSignupForm.Email)
	if emailUser.Email != "" {
		c.JSON(http.StatusConflict, gin.H{
			"error": fmt.Sprintf("Электронная почта '%s' уже существует", newSignupForm.Email),
		})
		return
	}

	user, token, registered := db.Signup(newSignupForm.Login, newSignupForm.Name, newSignupForm.Email, newSignupForm.Password)

	if registered {
		redisAuthorization := memory.RedisAuthorization{
			UserId:         user.Id,
			Login:          user.Login,
			Token:          token,
			Email:          user.Email,
			ConfirmedEmail: false,
		}

		redisUser := memory.RedisUser{
			UserId: user.Id,
			Login:  user.Login,
			Name:   user.Name,
			Image:  config.Config.Photo.Base64Small,
		}
		redis.SetAuthorization(redisAuthorization)
		redis.SetUser(redisUser)
		smtp.SendConfirmEmail(newSignupForm.Login, newSignupForm.Email, cache)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"image": redisUser.Image,
			"name": redisUser.Name,
    	})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервиса",
		})
	}
}
