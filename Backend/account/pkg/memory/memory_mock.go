package memory

import (
	"github.com/stretchr/testify/mock"
)

type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) GetUser(redisId int) RedisUser {
	args := m.Called(redisId)
	return args.Get(0).(RedisUser)
}

func (m *MockRedis) SetUser(redisUser RedisUser) {
	m.Called(redisUser)
}

func (m *MockRedis) DeleteUser(userId int) {
	m.Called(userId)
}

func (m *MockRedis) GetUsers(userIds []int) []RedisUser {
	args := m.Called(userIds)
	return args.Get(0).([]RedisUser)
}

func (m *MockRedis) GetUserFriends(userId int) RedisUserFriend {
	args := m.Called(userId)
	return args.Get(0).(RedisUserFriend)
}

func (m *MockRedis) AddUserFriend(redisUserFriend RedisUserFriend) {
	m.Called(redisUserFriend)
}

func (m *MockRedis) DeleteUserFriend(userId, userFriendId int) {
	m.Called(userId, userFriendId)
}

func (m *MockRedis) DeleteAllUserFriend(userId int) {
	m.Called(userId)
}

func (m *MockRedis) GetAuthorization(token string) RedisAuthorization {
	args := m.Called(token)
	return args.Get(0).(RedisAuthorization)
}

func (m *MockRedis) SetAuthorization(redisAuthorization RedisAuthorization) {
	m.Called(redisAuthorization)
}

func (m *MockRedis) DeleteAuthorization(token string) {
	m.Called(token)
}
