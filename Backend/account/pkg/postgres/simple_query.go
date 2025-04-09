package postgres

func (Instance *RealPostgres) CreateUser(user *User) bool {
	result := Instance.instance.Create(&user)
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

/*
func CreatePhoto(db *gorm.DB, name string, data []byte) {
	photo := Photo{Name: name, Data: data}
	result := db.Create(&photo)
	if result.Error != nil {
		log.Println("Error creating photo:", result.Error)
	}
}

func GetPhoto(db *gorm.DB, id uint) Photo {
	var photo Photo
	db.First(&photo, id)
	return photo
}
*/

func (Instance *RealPostgres) GetUserByUUID(id int) User {
	var user User
	Instance.instance.Model(&User{Id: id}).First(&user)
	return user
}

func (Instance *RealPostgres) GetUserByEmail(email string) User {
	var user User
	Instance.instance.Model(&User{Email: email}).First(&user)
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

func (Instance *RealPostgres) UUIDExists(uuid string) bool {
	var exists bool
	err := Instance.instance.Model(&User{}).Where("UUID = ?", uuid).Find(&exists)
	if err != nil {
		return false
	}
	return exists
}
