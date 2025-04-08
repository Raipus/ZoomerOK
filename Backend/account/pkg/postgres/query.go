package postgres

import (
	"encoding/base64"
	"log"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
func Signup(name string, email string, password string) (string, bool) {
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return "", false
	}

	newUUID := uuid.New().String()
	user := User{
		UUID:     newUUID,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Phone:    "",
		City:     "",
		Image:    config.Config.Photo.ByteImage,
	}

	resizedImage, err := security.ResizeImage(config.Config.Photo.ByteImage)

	if err != nil {
		return "", false
	}

	encodedImage := base64.StdEncoding.EncodeToString([]byte(resizedImage))
	userToken := security.UserToken{
		ID:    user.UUID,
		Name:  user.Name,
		Email: user.Email,
		Image: encodedImage,
	}

	token, err := security.GenerateJWT(userToken)
	if err != nil {
		return "", false
	}

	return token, CreateUser(&user)
}

// TODO: написать валидацию данных
func ChangePassword(email string, password string) {

}
