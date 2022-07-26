package mocks

import (
	models "REVAMPS-CMI-APPS/internal/entity/model"

	"github.com/golang/mock/gomock"
)

// MockGamesRepository is a mock of GamesRepository interface
type MockGamesRepository struct {
	ctrl *gomock.Controller
}

// NewMockGamesRepository creates a new mock instance
func NewMockGamesRepository(ctrl *gomock.Controller) *MockGamesRepository {
	mock := &MockGamesRepository{ctrl: ctrl}
	return mock
}

// Func get repo SQL
func (m *MockGamesRepository) Retreive(id string) (models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
