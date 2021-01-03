package strategy

import (
	"rover/src/constants"
	"rover/src/models"
)

type MoveNorth struct {
	coordinate *models.Coordinate
}

func NewMoveNorth(coordinate *models.Coordinate) MoveCriteria {
	return &MoveNorth{
		coordinate: coordinate,
	}
}

func (m *MoveNorth) GetDirection() string {
	return constants.NORTH
}

func (m *MoveNorth) Move() error {
	m.coordinate.SetY(m.coordinate.GetY() + 1)
	return nil
}

func (m *MoveNorth) Spin(command rune) MoveCriteria {
	if command == 'L' {
		return NewMoveWest(m.coordinate)
	} else if command == 'R' {
		return NewMoveEast(m.coordinate)
	}
	return m
}
