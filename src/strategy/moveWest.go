package strategy

import (
	"rover/src/constants"
	"rover/src/models"
)

type MoveWest struct {
	coordinate *models.Coordinate
}

func NewMoveWest(coordinate *models.Coordinate) MoveCriteria {
	return &MoveWest{
		coordinate: coordinate,
	}
}

func (m *MoveWest) GetDirection() string {
	return constants.WEST
}

func (m *MoveWest) Move() error {
	m.coordinate.SetX(m.coordinate.GetX() - 1)
	return nil
}

func (m *MoveWest) Spin(command rune) MoveCriteria {
	if command == 'L' {
		return NewMoveSouth(m.coordinate)
	} else if command == 'R' {
		return NewMoveNorth(m.coordinate)
	}
	return m
}
