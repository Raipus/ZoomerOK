package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type RegistryForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Registry(c *gin.Context) {
	var newRegistryForm RegistryForm
	if err := c.BindJSON(&newRegistryForm); err != nil {
		return
	}

	registered := postgres.Registry(newRegistryForm.Name, newRegistryForm.Email, newRegistryForm.Password)
	if registered {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}
