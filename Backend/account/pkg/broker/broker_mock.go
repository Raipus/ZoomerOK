package broker

import (
	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
	"github.com/stretchr/testify/mock"
)

type MockBroker struct {
	mock.Mock
}

func (m *MockBroker) PushUser(getUserResponse *pb.GetUserResponse) error {
	args := m.Called(getUserResponse)
	return args.Error(0)
}

func (m *MockBroker) PushUsers(getUsersResponse *pb.GetUsersResponse) error {
	args := m.Called(getUsersResponse)
	return args.Error(0)
}

func (m *MockBroker) PushUserFriend(getUserFriendRequest *pb.GetUserFriendResponse) error {
	args := m.Called(getUserFriendRequest)
	return args.Error(0)
}

func (m *MockBroker) Authorization(authorizationResponse *pb.AuthorizationRequest) error {
	args := m.Called(authorizationResponse)
	return args.Error(0)
}

func (m *MockBroker) Listen() {
	m.Called()
}
