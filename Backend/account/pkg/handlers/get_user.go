package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

// GetUser отправляет запрос на возвращение дополнительных данных пользователя.
// @Summary Отправить запрос на возвращение дополнительных данных пользователя.
// @Description Позволяет пользователю отправить запрос на возвращение дополнительных данных о пользователе (Name, Birthday, Phone, City, Image).
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param user body ChangePasswordForm true "Данные для изменения дополнительных данных о пользователе"
//
//	@Success 200 {object} struct {
//				     Id              int    `json:"id"`
//				     Login           string `json:"login"`
//				     Name            string `json:"name"`
//				     Email           string `json:"email"`
//			         Birthday        string `json:"birthday"`
//			         Phone           string `json:"phone"`
//		             City            string `json:"city"`
//			         Image           string `json:"Image"`
//	}
//
// @Failure 404 {object} gin.H {"error": "пользователь не найден!",}
// @Router /user/:login [get]
func GetUser(c *gin.Context, db postgres.PostgresInterface) {
	login := c.Param("login")

	user := db.GetUserByLogin(login)
	if postgres.CompareUsers(user, postgres.User{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "пользователь не найден!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.Id,
		"login":    user.Login,
		"name":     user.Name,
		"email":    user.Email,
		"birthday": user.Birthday,
		"phone":    user.Phone,
		"city":     user.City,
		"image":    user.Image,
	})
}
