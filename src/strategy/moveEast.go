package strategy

import (
	"rover/src/constants"
	"rover/src/models"
)

type MoveEast struct {
	coordinate *models.Coordinate
}

func NewMoveEast(coordinate *models.Coordinate) MoveCriteria {
	return &MoveEast{
		coordinate: coordinate,
	}
}

func (m *MoveEast) GetDirection() string {
	return constants.EAST
}

func (m *MoveEast) Move() error {
	m.coordinate.SetX(m.coordinate.GetX() + 1)
	return nil
}

func (m *MoveEast) Spin(command rune) MoveCriteria {
	if command == 'L' {
		return NewMoveNorth(m.coordinate)
	} else if command == 'R' {
		return NewMoveSouth(m.coordinate)
	}
	return m
}
