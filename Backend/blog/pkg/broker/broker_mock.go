package broker

import (
	"github.com/Raipus/ZoomerOK/blog/pkg/broker/pb"
	"github.com/stretchr/testify/mock"
)

type MockBroker struct {
	mock.Mock
}

func (m *MockBroker) PushUser(getUserRequest *pb.GetUserRequest) error {
	args := m.Called(getUserRequest)
	return args.Error(0)
}

func (m *MockBroker) PushUsers(getUsersRequest *pb.GetUsersRequest) error {
	args := m.Called(getUsersRequest)
	return args.Error(0)
}

func (m *MockBroker) PushUserFriend(getUserFriendRequest *pb.GetUserFriendRequest) error {
	args := m.Called(getUserFriendRequest)
	return args.Error(0)
}

func (m *MockBroker) Authorization(authorizationRequest *pb.AuthorizationRequest) error {
	args := m.Called(authorizationRequest)
	return args.Error(0)
}

func (m *MockBroker) Listen() {
	m.Called()
}
