package handlers

type ChangePasswordForm struct {
	Email       string `json:"email"`
	NewPassword string `json:"password"`
}

/*
func ChangePassword(c *gin.Context, db postgres.PostgresInterface) {
	var newChangePasswordForm ChangePasswordForm
	if err := c.BindJSON(&newChangePasswordForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	user := db.GetUserByEmail(newChangePasswordForm.Email)
	if user == (postgres.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Пользователь не найден",
		})
		return
	}

	if err := db.ChangePassword(user, newChangePasswordForm.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"email": newChangePasswordForm.Email,
		})
	}
}*/
