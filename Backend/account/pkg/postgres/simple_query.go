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
	Instance.instance.Model(&User{Id: id}).First(&user)
	return user
}

func (Instance *RealPostgres) GetUserByEmail(email string) User {
	var user User
	Instance.instance.Model(&User{Email: email}).First(&user)
	return user
}

func (Instance *RealPostgres) GetUserByLogin(login string) User {
	var user User
	Instance.instance.Model(&User{Login: login}).First(&user)
	return user
}

func (Instance *RealPostgres) DeleteUser(id int) {
	var user User
	Instance.instance.Where(&User{Id: id}).Find(&user)
	Instance.instance.Delete(&user)
}

func (Instance *RealPostgres) AcceptFriendRequest(id1 int, id2 int) {
	var friend Friend
	Instance.instance.Where(&Friend{User1Id: id1, User2Id: id2}).Find(&friend)
	friend.Accepted = true
	Instance.instance.Save(&friend)
}

func (Instance *RealPostgres) DeleteFriendRequest(id1 int, id2 int) {
	var friend Friend
	Instance.instance.Where(&Friend{User1Id: id1, User2Id: id2}).Find(&friend)
	Instance.instance.Delete(&friend)
}
