package postgres

import (
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
)

// TODO: написать валидацию данных
func (instance *RealPostgres) Login(loginOrEmail string, password string) (string, string) {
	var user User

	// Попытка найти пользователя по email
	resultByEmail := instance.instance.Where(&User{Email: loginOrEmail}).First(&user)
	if resultByEmail.Error == nil { // Пользователь найден по email
		// Проверка пароля
		if !security.CheckPasswordHash(password, user.Password) {
			return "", "Неверный пароль"
		}
		// Генерация токена
		return generateToken(user)
	}

	// Попытка найти пользователя по логину
	resultByLogin := instance.instance.Where(&User{Login: loginOrEmail}).First(&user)
	if resultByLogin.Error == nil { // Пользователь найден по логину
		// Проверка пароля
		if !security.CheckPasswordHash(password, user.Password) {
			return "", "Неверный пароль"
		}
		// Генерация токена
		return generateToken(user)
	}

	// Если пользователь не найден ни по email, ни по логину
	return "", "Неверный login или email"
}

func generateToken(user User) (string, string) {
	userToken := security.UserToken{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	token, err := security.GenerateJWT(userToken)
	if err != nil {
		return "", "Ошибка сервера"
	}

	return token, ""
}

// TODO: написать валидацию данных
func (Instance *RealPostgres) Signup(login, name, email, password string) (string, bool) {
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return "", false
	}

	user := User{
		Login:    login,
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
