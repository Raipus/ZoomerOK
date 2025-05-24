package postgres

import "gorm.io/gorm"

type PostgresInterface interface {
	Login(loginOrEmail string, password string) (User, string, string)
	Signup(login string, name string, email string, password string) (User, string, bool)
	ChangePassword(user *User, newPassword string) error
	CreateUser(user *User) (User, bool)
	UpdateUserPassword(user *User, newPassword string) error
	ChangeUser(user *User) bool
	ConfirmEmail(login string) (User, bool)
	GetUserById(id int) User
	GetUserByEmail(email string) User
	GetUserByLogin(login string) User
	DeleteUser(user *User)
	AcceptFriendRequest(id1 int, id2 int) error
	AddFriendRequest(id1 int, id2 int) error
	ExistFriendRequest(id1 int, id2 int) (Friend, error)
	CheckUserExist(id1 int, id2 int) error
	DeleteFriendRequest(id1 int, id2 int) error
	GetUnacceptedFriends(userId int) ([]User, error)
	CheckUserFriend(userId, friendId int) (string, error)
}

var ProductionPostgresInterface PostgresInterface = &RealPostgres{instance: initPostgres()}

type RealPostgres struct {
	instance *gorm.DB
}
