package postgres

func (Instance *RealPostgres) CreateUser(user *User) bool {
	result := Instance.instance.Create(&user)
	if result.Error != nil {
		return false
	}

	return true
}

func (Instance *RealPostgres) ChangeUser(user *User) bool {
	result := Instance.instance.Save(user)
	if result.Error != nil {
		return false
	}
	return true
}

func (Instance *RealPostgres) UpdateUserPassword(user *User, newPassword string) error {
	user.Password = newPassword
	if err := Instance.instance.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (Instance *RealPostgres) GetUserById(id int) User {
	var user User
	Instance.instance.Where("id = ?", id).First(&user)
	return user
}

func (Instance *RealPostgres) GetUserByEmail(email string) User {
	var user User
	Instance.instance.Where("email = ?", email).First(&user)
	return user
}

func (Instance *RealPostgres) GetUserByLogin(login string) User {
	var user User
	Instance.instance.Where("login = ?", login).First(&user)
	return user
}

func (Instance *RealPostgres) ConfirmEmail(login string) (User, bool) {
	user := Instance.GetUserByLogin(login)
	if user.Login == "" {
		return user, false
	}

	user.ConfirmedEmail = true
	result := Instance.instance.Save(user)
	if result.Error != nil {
		return user, false
	}
	return user, true
}

func (Instance *RealPostgres) DeleteUser(user *User) {
	Instance.instance.Delete(&user)
}
