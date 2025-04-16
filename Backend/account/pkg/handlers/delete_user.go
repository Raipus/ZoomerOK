package handlers

import (
	"net/http"

	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type DeleteUserForm struct {
	Id int
}

func DeleteUser(c *gin.Context, db postgres.PostgresInterface) {
	var newDeleteUserForm DeleteUserForm
	if err := c.BindJSON(&newDeleteUserForm); err != nil {
		return
	}

	db.DeleteUser(newDeleteUserForm.Id)
	c.JSON(http.StatusOK, gin.H{})
}
