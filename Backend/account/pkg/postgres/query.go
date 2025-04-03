package postgres

import (
	"log"

	"github.com/Raipus/ZoomerOK/account/pkg/security/hash"
	"gorm.io/gorm"
)

func CreateUser(user *User) {
	Instance.Create(&user)
}

func UpdateUser(user *User) {
	Instance.Model(&Settings{}).Where("UUID = ?", UUID).Updates(map[string]interface{}{"ColorBlindnessType": ColorBlindnessType, "Degree": Degree})
}

func Login(email string, password string) (bool, string) {
	var user User
	result := Instance.Where(&User{Email: email}).Find(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, "Неверный email"
		}
		log.Println("Ошибка базы данных:", result.Error)
		return false, "Ошибка сервера"
	}

	checked := hash.CheckPasswordHash(password, user.Password)
	if !checked {
		return false, "Неверный пароль"
	}

	return true, ""
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

func UUIDExists(uuid string) bool {
	var exists bool
	err := Instance.Model(&User{}).Where("UUID = ?", uuid).Find(&exists)
	if err != nil {
		return false
	}
	return exists
}
