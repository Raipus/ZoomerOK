package postgres

func CreateUser(user *User) bool {
	result := Instance.Create(&user)
	if result.Error != nil {
		return false
	}

	return true
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

func GetUserByUUID(uuid string) User {
	var user User
	Instance.Model(&User{UUID: uuid}).First(&user)
	return user
}

func GetUserByEmail(email string) User {
	var user User
	Instance.Model(&User{Email: email}).First(&user)
	return user
}

func DeleteUser(uuid string) {
	var user User
	Instance.Where(&User{UUID: uuid}).Find(&user)
	Instance.Delete(&user)
}

func AcceptFriendRequest(uuid1 string, uuid2 string) {
	var friend Friend
	Instance.Where(&Friend{User1UUID: uuid1, User2UUID: uuid2}).Find(&friend)
	friend.Accepted = true
	Instance.Save(&friend)
}

func DeleteFriendRequest(uuid1 string, uuid2 string) {
	var friend Friend
	Instance.Where(&Friend{User1UUID: uuid1, User2UUID: uuid2}).Find(&friend)
	Instance.Delete(&friend)
}

func UUIDExists(uuid string) bool {
	var exists bool
	err := Instance.Model(&User{}).Where("UUID = ?", uuid).Find(&exists)
	if err != nil {
		return false
	}
	return exists
}
