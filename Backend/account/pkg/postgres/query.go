package postgres

import (
	"log"

	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(user *User) bool {
	result := Instance.Create(&user)
	if result.Error != nil {
		return false
	}

	return true
}

// TODO: написать валидацию данных
func Login(email string, password string) (bool, string) {
	var user User
	result := Instance.Where(&User{Email: email}).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, "Неверный email"
		}
		log.Println("Ошибка базы данных:", result.Error)
		return false, "Ошибка сервера"
	}

	checked := security.CheckPasswordHash(password, user.Password)
	if !checked {
		return false, "Неверный пароль"
	}

	return true, ""
}

// TODO: написать валидацию данных
func Registry(name string, email string, password string) bool {
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return false
	}

	newUUID := uuid.New().String()
	user := User{
		UUID:     newUUID,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Phone:    "",
		City:     "",
	}

	return CreateUser(&user)
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
