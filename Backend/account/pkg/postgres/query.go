package postgres

func CreateUser(user *User) {
	Instance.Create(&user)
}

func UpdateUser(user *User) {
	Instance.Model(&Settings{}).Where("UUID = ?", UUID).Updates(map[string]interface{}{"ColorBlindnessType": ColorBlindnessType, "Degree": Degree})
}

func GetUser(uuid string) User {
	var user User
	Instance.Model(&User{UUID: uuid}).First(&user)
	return user
}

func DeleteUser(uuid string) {
	var user User
	Instance.Where(&User{UUID: uuid}).Find(&user)
	Instance.Delete(&user)
}

func CreateFriendRequest(uuid1 string, uuid2 string) {
	var user1 User
	var user2 User
	Instance.Create(&Friend{User1: user1, User2: user2, Accepted: false})
}

func AcceptFriendRequest(uuid1 string, uuid2 string) {
	var friend Friend
	var user1 User
	var user2 User
	Instance.Where(&User{UUID: uuid1}).Find(&user1)
	Instance.Where(&User{UUID: uuid2}).Find(&user2)
	Instance.Where(&Friend{User1: user1, User2: user2}).Find(&friend)
	friend.Accepted = true
	Instance.Save(&friend)
}

func DeleteFriendRequest(uuid1 string, uuid2 string) {
	var friend Friend
	var user1 User
	var user2 User
	Instance.Where(&User{UUID: uuid1}).Find(&user1)
	Instance.Where(&User{UUID: uuid2}).Find(&user2)
	Instance.Where(&Friend{User1: user1, User2: user2}).Find(&friend)
	Instance.Delete(&friend)
}
