package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// LoginForm представляет данные, необходимые для входа в социальную сеть.
type LoginForm struct {
	LoginOrEmail string `json:"login_or_email"` // LoginOrEmail пользователя, который проверяет логин или электронную почту пользователя.
	Password     string `json:"password"`       // Password пользователя, который принимает пароль пользователя.
}

// Login отправляет запрос на вход в социальную сеть пользователю.
// @Summary совершить вход в систему (социальная сеть)
// @Description Позволяет пользователю отправить запрос на вход в систему.
// @Accept json
// @Produce json
// @Param user body LoginForm true "Данные для входа в систему"
// @Success 200 {object} gin.H {"token": "JWT токен"}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /login [post]
func Login(c *gin.Context, db postgres.PostgresInterface, redis memory.RedisInterface) {
	var newLoginForm LoginForm
	if err := c.BindJSON(&newLoginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	user, token, error := db.Login(newLoginForm.LoginOrEmail, newLoginForm.Password)
	if error != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	} else {
		redisAuthorization := memory.RedisAuthorization{
			UserId:         user.Id,
			Token:          token,
			Login:          user.Login,
			Email:          user.Email,
			ConfirmedEmail: user.ConfirmedEmail,
		}
		redis.SetAuthorization(redisAuthorization)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
