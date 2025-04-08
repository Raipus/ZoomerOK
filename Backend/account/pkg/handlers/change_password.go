package handlers

type ChangePasswordForm struct {
	Email       string `json:"email"`
	NewPassword string `json:"password"`
}

/*
func ChangePassword(c *gin.Context) {
	var newChangePasswordForm ChangePasswordForm
	if err := c.BindJSON(&newChangePasswordForm); err != nil {
		return
	}

	user := GetUserByEmail(newChangePasswordForm.Email)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Пользователь не найден",
		})
	}

	if registered {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}
*/
