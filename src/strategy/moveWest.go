package strategy

import (
	"rover/src/constants"
)

type MoveWest struct {
	coordinate *Coordinate
}

func NewMoveWest(coordinate *Coordinate) MoveCriteria {
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

func (m *MoveWest) Spin(instruction rune) MoveCriteria {
	if instruction == constants.LEFT {
		return NewMoveSouth(m.coordinate)
	} else if instruction == constants.RIGHT {
		return NewMoveNorth(m.coordinate)
	}
	return m
}
