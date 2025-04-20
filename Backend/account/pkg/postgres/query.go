package postgres

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/Raipus/ZoomerOK/account/pkg/security"
)

func (instance *RealPostgres) Login(loginOrEmail string, password string) (string, string) {
	var user User

	resultByEmail := instance.instance.Where(&User{Email: loginOrEmail}).First(&user)
	if resultByEmail.Error == nil {
		if !security.CheckPasswordHash(password, user.Password) {
			return "", "Неверный пароль"
		}
		return generateToken(user)
	}

	resultByLogin := instance.instance.Where(&User{Login: loginOrEmail}).First(&user)
	if resultByLogin.Error == nil {
		if !security.CheckPasswordHash(password, user.Password) {
			return "", "Неверный пароль"
		}
		return generateToken(user)
	}

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

func (Instance *RealPostgres) CheckUserExist(id1 int, id2 int) error {
	var user1 User
	result1 := Instance.instance.First(&user1, id1)
	if result1.Error != nil {
		if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("пользователь с id %d не найден", id1)
		}
		return fmt.Errorf("ошибка при поиске пользователя с id %d: %w", id1, result1.Error)
	}

	var user2 User
	result2 := Instance.instance.First(&user2, id2)
	if result2.Error != nil {
		if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("пользователь с id %d не найден", id2)
		}
		return fmt.Errorf("ошибка при поиске пользователя с id %d: %w", id2, result2.Error)
	}

	return nil
}

func (Instance *RealPostgres) AcceptFriendRequest(id1 int, id2 int) error {
	if err := Instance.CheckUserExist(id1, id2); err != nil {
		return err
	}

	var friend Friend
	result := Instance.instance.Where(
		"(user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)",
		id1, id2, id2, id1,
	).First(&friend)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("запрос в друзья между пользователями %d и %d не найден", id1, id2)
		}
		return fmt.Errorf("ошибка при поиске запроса в друзья: %w", result.Error)
	}
	friend.Accepted = true
	result = Instance.instance.Save(&friend)
	return nil
}

func (Instance *RealPostgres) AddFriendRequest(id1 int, id2 int) error {
	if err := Instance.CheckUserExist(id1, id2); err != nil {
		return err
	}

	newFriend := Friend{
		User1Id:  id1,
		User2Id:  id2,
		Accepted: false,
	}

	Instance.instance.Create(&newFriend)
	return nil
}

func (Instance *RealPostgres) ExistFriendRequest(id1 int, id2 int) (Friend, error) {
	var friend Friend
	result := Instance.instance.Where(
		"(user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)",
		id1, id2, id2, id1,
	).First(&friend)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Friend{}, fmt.Errorf("запрос в друзья между пользователями %d и %d не найден", id1, id2)
		}
		return Friend{}, fmt.Errorf("ошибка при поиске запроса в друзья: %w", result.Error)
	}

	return friend, nil
}

func (Instance *RealPostgres) DeleteFriendRequest(id1 int, id2 int) error {
	if err := Instance.CheckUserExist(id1, id2); err != nil {
		return err
	}

	friend, err := Instance.ExistFriendRequest(id1, id2)
	if err != nil {
		return err
	}
	Instance.instance.Delete(&friend)
	return nil
}
