package mocks

import (
	"github.com/stretchr/testify/mock"
	"rover/src/strategy"
)

type MockMoveNorth struct {
	mock.Mock
}

func (m *MockMoveNorth) GetDirection() string {
	args := m.Called()
	if len(args) != 1 {
		return ""
	}
	return args.Get(0).(string)

}

func (m *MockMoveNorth) Move() error {
	m.Called()
	return nil
}

func (m *MockMoveNorth) Spin(instruction rune) strategy.MoveCriteria {
	m.Called()
	return nil
}
