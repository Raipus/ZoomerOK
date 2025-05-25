package postgres

import (
	"github.com/stretchr/testify/mock"
)

type MockPostgres struct {
	mock.Mock
}

func (m *MockPostgres) CreatePost(userId int, text string, image []byte) (int, error) {
	args := m.Called(userId, text, image)
	return args.Int(0), args.Error(1)
}

func (m *MockPostgres) DeletePost(userId int, postId int) error {
	args := m.Called(userId, postId)
	return args.Error(0)
}

func (m *MockPostgres) CreateComment(userId, postId int, text string) error {
	args := m.Called(userId, postId, text)
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

func (m *MockPostgres) GetPosts(userIds []int, page int) ([]Post, error) {
	args := m.Called(userIds, page)
	return args.Get(0).([]Post), args.Error(1)
}

func (m *MockPostgres) GetCountCommentsAndLikes(postIds []int) (map[int]int, map[int]int, error) {
	args := m.Called(postIds)
	return args.Get(0).(map[int]int), args.Get(1).(map[int]int), args.Error(2)
}

func (m *MockPostgres) GetComments(postId, page int) ([]Comment, error) {
	args := m.Called(postId, page)
	return args.Get(0).([]Comment), args.Error(1)
}

func (m *MockPostgres) Like(postId int, userId int) (bool, error) {
	args := m.Called(postId, userId)
	return args.Bool(0), args.Error(1)
}
