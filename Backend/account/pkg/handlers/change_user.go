package handlers

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/memory"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/gin-gonic/gin"
)

// ChangeUserForm представляет данные, необходимые для изменения дополнительных пользовательских данных.
type ChangeUserForm struct {
	Name     string     `json:"name,omitempty"`     // Name, Полное имя пользователя.
	Birthday *time.Time `json:"birthday,omitempty"` // Birthday, день рождения пользователя.
	Phone    string     `json:"phone,omitempty"`    // Phone, телефон пользователя.
	City     string     `json:"city,omitempty"`     // City, город пользователя.
	Image    []byte     `json:"image,omitempty"`    // Image, электронная почта пользователя.
}

// ChangeUser отправляет запрос на изменение дополнительх данных пользователя.
// @Summary Отправить запрос на изменение дополнительных данных пользователя.
// @Description Позволяет пользователю отправить запрос на изменение дополнительных данные (Name, Birthday, Phone, City, Image).
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user body ChangePasswordForm true "Данные для изменения дополнительных данных о пользователе"
// @Success 204 {object} gin.H {}
// @Failure 400 {object} gin.H {"error": "Некорректный формат данных"}
// @Failure 401 {object} gin.H {"error": "Пользователь не имеет права на данную операцию"}
// @Failure 500 {object} gin.H {"error": "Ошибка сервиса"}
// @Router /user/:login [put]
func ChangeUser(c *gin.Context, db postgres.PostgresInterface, redis memory.RedisInterface) {
	login := c.Param("login")

	var newChangeUser ChangeUserForm
	if err := c.BindJSON(&newChangeUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат данных: " + err.Error(),
		})
		return
	}

	user := db.GetUserByLogin(login)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "пользователь не найден!",
		})
		return
	}

	if user.Login != login {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не имеет права на данную операцию"})
		return
	}

	user.Name = newChangeUser.Name
	user.Birthday = newChangeUser.Birthday
	user.Phone = newChangeUser.Phone
	user.City = newChangeUser.City
	user.Image = newChangeUser.Image
	if success := db.ChangeUser(&user); !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}

	byteImage, err := security.ResizeImage(user.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервиса"})
		return
	}
	encoded := base64.StdEncoding.EncodeToString(byteImage)

	redisUser := memory.RedisUser{
		UserId: user.Id,
		Login:  user.Login,
		Name:   user.Name,
		Image:  encoded,
	}
	redis.SetUser(redisUser)

	c.JSON(http.StatusNoContent, gin.H{})
}
