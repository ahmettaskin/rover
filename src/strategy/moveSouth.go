package strategy

import (
	"rover/src/constants"
	"rover/src/models"
)

type MoveSouth struct {
	coordinate *models.Coordinate
}

func NewMoveSouth(coordinate *models.Coordinate) MoveCriteria {
	return &MoveSouth{
		coordinate: coordinate,
	}
}

func (m *MoveSouth) GetDirection() string {
	return constants.SOUTH
}

func (m *MoveSouth) Move() error {
	m.coordinate.SetY(m.coordinate.GetY() - 1)
	return nil
}

func (m *MoveSouth) Spin(command rune) MoveCriteria {
	if command == 'L' {
		return NewMoveEast(m.coordinate)
	} else if command == 'R' {
		return NewMoveWest(m.coordinate)
	}
	return m
}
