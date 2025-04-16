package postgres

import "gorm.io/gorm"

type PostgresInterface interface {
	Login(email string, password string) (bool, string)
	Signup(name string, email string, password string) (string, bool)
	ChangePassword(user *User, newPassword string) error
	CreateUser(user *User) bool
	UpdateUserPassword(user *User, newPassword string) error
	ChangeUser(user *User) bool
	GetUserById(id int) User
	GetUserByEmail(email string) User
	DeleteUser(id int)
	AcceptFriendRequest(id1 int, id2 int)
	DeleteFriendRequest(id1 int, id2 int)
}

var ProductionPostgresInterface PostgresInterface = &RealPostgres{instance: initPostgres()}

type RealPostgres struct {
	instance *gorm.DB
}
