package memory

import (
	"github.com/stretchr/testify/mock"
)

type MockMessageQueue struct {
	mock.Mock
}

func (m *MockMessageQueue) Update(msg interface{}) {
	m.Called(msg)
}

func (m *MockMessageQueue) GetLastMessage() interface{} {
	args := m.Called()
	return args.Get(0).(interface{})
}
