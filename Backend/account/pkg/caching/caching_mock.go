package caching

import (
	"github.com/stretchr/testify/mock"
)

// MockCache - структура для мокирования кэша
type MockCache struct {
	mock.Mock
}

// Реализация методов интерфейса CachingInterface для MockCache
func (m *MockCache) SetCacheResetLink(login, resetLink string) {
	m.Called(login, resetLink)
}

func (m *MockCache) SetCacheConfirmationLink(login, confirmationLink string) {
	m.Called(login, confirmationLink)
}

func (m *MockCache) GetCacheResetLink(resetLink string) string {
	args := m.Called(resetLink)
	return args.String(0)
}

func (m *MockCache) GetCacheConfirmationLink(confirmationLink string) string {
	args := m.Called(confirmationLink)
	return args.String(0)
}

func (m *MockCache) DeleteCacheResetLink(resetLink string) {
	m.Called(resetLink)
}

func (m *MockCache) DeleteCacheConfirmationLink(confirmationLink string) {
	m.Called(confirmationLink)
}
