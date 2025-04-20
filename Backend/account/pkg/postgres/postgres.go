package postgres

import "gorm.io/gorm"

type PostgresInterface interface {
	Login(loginOrEmail string, password string) (string, string)
	Signup(login string, name string, email string, password string) (string, bool)
	ChangePassword(user *User, newPassword string) error
	CreateUser(user *User) bool
	UpdateUserPassword(user *User, newPassword string) error
	ChangeUser(user *User) bool
	ConfirmEmail(login string) bool
	GetUserById(id int) User
	GetUserByEmail(email string) User
	GetUserByLogin(login string) User
	DeleteUser(id int)
	AcceptFriendRequest(id1 int, id2 int)
	DeleteFriendRequest(id1 int, id2 int)
}

var ProductionPostgresInterface PostgresInterface = &RealPostgres{instance: initPostgres()}

type RealPostgres struct {
	instance *gorm.DB
}
