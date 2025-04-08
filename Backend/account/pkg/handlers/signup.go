package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type SignupForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var newSignupForm SignupForm
	if err := c.BindJSON(&newSignupForm); err != nil {
		return
	}

	token, registered := postgres.Registry(newSignupForm.Name, newSignupForm.Email, newSignupForm.Password)

	if registered {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}
