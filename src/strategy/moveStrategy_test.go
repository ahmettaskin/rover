package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCriteria_ShouldReturnMoveCriteria_WhenCalled(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveStrategy := NewMoveStrategy(NewMoveSouth(coordinate))

	moveCriteria := moveStrategy.GetCriteria()

	assert.Equal(t, moveCriteria, NewMoveSouth(coordinate))
}

func TestSetCriteria_ShouldSetMoveCriteria_WhenCalled(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveStrategy := NewMoveStrategy(NewMoveSouth(coordinate))

	moveStrategy.SetCriteria(NewMoveNorth(coordinate))

	assert.Equal(t, moveStrategy.GetCriteria(), NewMoveNorth(coordinate))
}
