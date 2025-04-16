package broker

import (
	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
	"github.com/stretchr/testify/mock"
)

type MockBroker struct {
	mock.Mock
}

func (m *MockBroker) PushGetUser(redisUser *pb.GetUserRequest) error {
	args := m.Called(redisUser)
	return args.Error(0)
}

func (m *MockBroker) PushGetUserFriend(redisUserFriend *pb.GetUserFriendRequest) error {
	args := m.Called(redisUserFriend)
	return args.Error(0)
}

func (m *MockBroker) Listen() {
	m.Called()
}
