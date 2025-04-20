package postgres

import (
	"github.com/stretchr/testify/mock"
)

type MockPostgres struct {
	mock.Mock
}

func (m *MockPostgres) CreatePost(userId int, text string, photo []byte) error {
	args := m.Called(userId, text, photo)
	return args.Error(0)
}

func (m *MockPostgres) DeletePost(userId int, postId int) error {
	args := m.Called(userId, postId)
	return args.Error(0)
}

func (m *MockPostgres) CreateComment(userId int, text string) error {
	args := m.Called(userId, text)
	return args.Error(0)
}

func (m *MockPostgres) DeleteComment(userId int, commentId int) error {
	args := m.Called(userId, commentId)
	return args.Error(0)
}

func (m *MockPostgres) GetPost(postId int) (*Post, error) {
	args := m.Called(postId)
	return args.Get(0).(*Post), args.Error(1)
}

func (m *MockPostgres) GetPosts(userId int) ([]Post, error) {
	args := m.Called(userId)
	return args.Get(0).([]Post), args.Error(1)
}

func (m *MockPostgres) Like(postId int, userId int) error {
	args := m.Called(postId, userId)
	return args.Error(0)
}
