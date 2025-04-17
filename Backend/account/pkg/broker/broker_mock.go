package broker

import (
	"github.com/Raipus/ZoomerOK/account/pkg/broker/pb"
	"github.com/stretchr/testify/mock"
)

type MockBroker struct {
	mock.Mock
}

func (m *MockBroker) PushUser(redisUser *pb.GetUserResponse) error {
	args := m.Called(redisUser)
	return args.Error(0)
}

func (m *MockBroker) PushUserFriend(redisUserFriend *pb.GetUserFriendResponse) error {
	args := m.Called(redisUserFriend)
	return args.Error(0)
}

func (m *MockBroker) Listen() {
	m.Called()
}
