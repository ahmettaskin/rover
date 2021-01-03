package strategy

import (
	"rover/src/constants"
)

type MoveSouth struct {
	coordinate *Coordinate
}

func NewMoveSouth(coordinate *Coordinate) MoveCriteria {
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

func (m *MoveSouth) Spin(instruction rune) MoveCriteria {
	if instruction == constants.LEFT {
		return NewMoveEast(m.coordinate)
	} else if instruction == constants.RIGHT {
		return NewMoveWest(m.coordinate)
	}
	return m
}
