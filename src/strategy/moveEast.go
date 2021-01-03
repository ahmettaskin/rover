package strategy

import (
	"rover/src/constants"
)

type MoveEast struct {
	coordinate *Coordinate
}

func NewMoveEast(coordinate *Coordinate) MoveCriteria {
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

func (m *MoveEast) Spin(instruction rune) MoveCriteria {
	if instruction == constants.LEFT {
		return NewMoveNorth(m.coordinate)
	} else if instruction == constants.RIGHT {
		return NewMoveSouth(m.coordinate)
	}
	return m
}
