package memory

import (
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/stretchr/testify/mock"
)

type MockMessageStore struct {
	mock.Mock
}

func (m *MockMessageStore) SaveMessage(key string, message interface{}) error {
	args := m.Called(key, message)
	return args.Error(0)
}

func (m *MockMessageStore) GetMessage(key string) (interface{}, bool) {
	args := m.Called(key)
	return args.Get(0), args.Bool(1)
}

func (m *MockMessageStore) ProcessAuthorization(authorizationRequest *pb.AuthorizationRequest) (pb.AuthorizationResponse, error) {
	args := m.Called(authorizationRequest)
	return args.Get(0).(pb.AuthorizationResponse), args.Error(1)
}

func (m *MockMessageStore) ProcessPushUserFriend(getUserFriendRequest *pb.GetUserFriendRequest) (pb.GetUserFriendResponse, error) {
	args := m.Called(getUserFriendRequest)
	return args.Get(0).(pb.GetUserFriendResponse), args.Error(1)
}

func (m *MockMessageStore) ProcessPushUsers(getUsersRequest *pb.GetUsersRequest) (pb.GetUsersResponse, error) {
	args := m.Called(getUsersRequest)
	return args.Get(0).(pb.GetUsersResponse), args.Error(1)
}

func (m *MockMessageStore) ProcessPushUser(getUserRequest *pb.GetUserRequest) (pb.GetUserResponse, error) {
	args := m.Called(getUserRequest)
	return args.Get(0).(pb.GetUserResponse), args.Error(1)
}
