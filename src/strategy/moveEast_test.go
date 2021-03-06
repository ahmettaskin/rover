package strategy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMoveEast_ShouldMoveToEast_WhenCalled(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveEast(coordinate)

	moveEast.Move()

	assert.Equal(t, 2, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
}

func TestSpinEast_ShouldSpinToRight_WhenInstructionIsR(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveEast(coordinate)

	spin := moveEast.Spin('R')

	assert.Equal(t, 1, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
	fmt.Println(reflect.TypeOf(spin))
	assert.Equal(t, reflect.TypeOf(NewMoveSouth(coordinate)), reflect.TypeOf(spin))
}

func TestSpinEast_ShouldSpinToLest_WhenInstructionIsL(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveEast(coordinate)

	spin := moveEast.Spin('L')

	assert.Equal(t, 1, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
	fmt.Println(reflect.TypeOf(spin))
	assert.Equal(t, reflect.TypeOf(NewMoveNorth(coordinate)), reflect.TypeOf(spin))
}
