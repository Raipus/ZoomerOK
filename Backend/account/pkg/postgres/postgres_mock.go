package postgres

import (
	"github.com/stretchr/testify/mock"
)

type MockPostgres struct {
	mock.Mock
}

func (m *MockPostgres) Login(loginOrEmail string, password string) (string, string) {
	args := m.Called(loginOrEmail, password)
	return args.String(0), args.String(1)
}

func (m *MockPostgres) Signup(login, name, email, password string) (string, bool) {
	args := m.Called(login, name, email, password)
	return args.String(0), args.Bool(1)
}

func (m *MockPostgres) ChangePassword(user *User, newPassword string) error {
	args := m.Called(user, newPassword)
	return args.Error(0)
}

func (m *MockPostgres) CreateUser(user *User) bool {
	args := m.Called(user)
	return args.Bool(0)
}

func (m *MockPostgres) UpdateUserPassword(user *User, newPassword string) error {
	args := m.Called(user, newPassword)
	return args.Error(0)
}

func (m *MockPostgres) ChangeUser(user *User) bool {
	args := m.Called(user)
	return args.Bool(0)
}

func (m *MockPostgres) ConfirmEmail(login string) bool {
	args := m.Called(login)
	return args.Bool(0)
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

func (m *MockPostgres) DeleteUser(id int) {
	m.Called(id)
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
