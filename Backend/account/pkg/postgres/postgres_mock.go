package postgres

import (
	"github.com/stretchr/testify/mock"
)

type MockPostgres struct {
	mock.Mock
}

func (m *MockPostgres) Login(loginOrEmail string, password string) (User, string, string) {
	args := m.Called(loginOrEmail, password)
	return args.Get(0).(User), args.String(1), args.String(2)
}

func (m *MockPostgres) Signup(login, name, email, password string) (User, string, bool) {
	args := m.Called(login, name, email, password)
	return args.Get(0).(User), args.String(1), args.Bool(2)
}

func (m *MockPostgres) ChangePassword(user *User, newPassword string) error {
	args := m.Called(user, newPassword)
	return args.Error(0)
}

func (m *MockPostgres) CreateUser(user *User) (User, bool) {
	args := m.Called(user)
	return args.Get(0).(User), args.Bool(1)
}

func (m *MockPostgres) UpdateUserPassword(user *User, newPassword string) error {
	args := m.Called(user, newPassword)
	return args.Error(0)
}

func (m *MockPostgres) ChangeUser(user *User) bool {
	args := m.Called(user)
	return args.Bool(0)
}

func (m *MockPostgres) ConfirmEmail(login string) (User, bool) {
	args := m.Called(login)
	return args.Get(0).(User), args.Bool(1)
}

func (m *MockPostgres) GetUserById(id int) User {
	args := m.Called(id)
	return args.Get(0).(User)
}

func (m *MockPostgres) GetUserByEmail(email string) User {
	args := m.Called(email)
	return args.Get(0).(User)
}

func (m *MockPostgres) GetUserByLogin(login string) User {
	args := m.Called(login)
	return args.Get(0).(User)
}

func (m *MockPostgres) DeleteUser(user *User) {
	m.Called(user)
}

func (m *MockPostgres) AcceptFriendRequest(id1 int, id2 int) error {
	args := m.Called(id1, id2)
	return args.Error(0)
}

func (m *MockPostgres) AddFriendRequest(id1 int, id2 int) error {
	args := m.Called(id1, id2)
	return args.Error(0)
}

func (m *MockPostgres) DeleteFriendRequest(id1 int, id2 int) error {
	args := m.Called(id1, id2)
	return args.Error(0)
}

func (m *MockPostgres) ExistFriendRequest(id1 int, id2 int) (Friend, error) {
	args := m.Called(id1, id2)
	return args.Get(0).(Friend), args.Error(1)
}

func (m *MockPostgres) CheckUserExist(id1 int, id2 int) error {
	args := m.Called(id1, id2)
	return args.Error(0)
}

func (m *MockPostgres) GetUnacceptedFriends(userId int) ([]User, error) {
	args := m.Called(userId)
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockPostgres) CheckUserFriend(userId, friendId int) (string, error) {
	args := m.Called(userId, friendId)
	return args.String(0), args.Error(1)
}
