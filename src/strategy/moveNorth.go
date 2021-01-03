package strategy

import (
	"rover/src/constants"
)

type MoveNorth struct {
	coordinate *Coordinate
}

func NewMoveNorth(coordinate *Coordinate) MoveCriteria {
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

func (m *MoveNorth) Spin(instruction rune) MoveCriteria {
	if instruction == constants.LEFT {
		return NewMoveWest(m.coordinate)
	} else if instruction == constants.RIGHT {
		return NewMoveEast(m.coordinate)
	}
	return m
}
