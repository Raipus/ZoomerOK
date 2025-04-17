package security

import (
	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/stretchr/testify/mock"
)

type MockSmtp struct {
	mock.Mock
}

func (m *MockSmtp) SendConfirmEmail(username, email string, cache caching.CachingInterface) error {
	args := m.Called(username, email, cache)
	return args.Error(0)
}

func (m *MockSmtp) SendChangePassword(username, email string, cache caching.CachingInterface) error {
	args := m.Called(username, email, cache)
	return args.Error(0)
}

func (m *MockSmtp) SendEmail(email string, message []byte) error {
	args := m.Called(email, message)
	return args.Error(0)
}
