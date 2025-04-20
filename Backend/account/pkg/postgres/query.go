package postgres

import (
	"log"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
	"gorm.io/gorm"
)

// TODO: написать валидацию данных
func (Instance *RealPostgres) Login(email string, password string) (bool, string) {
	var user User
	result := Instance.instance.Where(&User{Email: email}).First(&user)

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
func (Instance *RealPostgres) Signup(name string, email string, password string) (string, bool) {
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return "", false
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Phone:    "",
		City:     "",
		Image:    config.Config.Photo.ByteImage,
	}

	// resizedImage, err := security.ResizeImage(config.Config.Photo.ByteImage)

	if err != nil {
		return "", false
	}

	// encodedImage := base64.StdEncoding.EncodeToString([]byte(resizedImage))
	userToken := security.UserToken{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	token, err := security.GenerateJWT(userToken)
	if err != nil {
		return "", false
	}

	return token, Instance.CreateUser(&user)
}

// TODO: написать валидацию данных
func (Instance *RealPostgres) ChangePassword(user *User, newPassword string) error {
	hashedPassword, err := security.HashPassword(newPassword)
	if err != nil {
		return err
	}
	return Instance.UpdateUserPassword(user, hashedPassword)
}
