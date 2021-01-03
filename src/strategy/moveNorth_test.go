package strategy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMoveNorth_ShouldMoveToNorth_WhenCalled(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveNorth(coordinate)

	moveEast.Move()

	assert.Equal(t, 1, coordinate.GetX())
	assert.Equal(t, 2, coordinate.GetY())
}

func TestSpinNorth_ShouldSpinToRight_WhenInstructionIsR(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveNorth(coordinate)

	spin := moveEast.Spin('R')

	assert.Equal(t, 1, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
	fmt.Println(reflect.TypeOf(spin))
	assert.Equal(t, reflect.TypeOf(NewMoveEast(coordinate)), reflect.TypeOf(spin))
}

func TestSpinNorth_ShouldSpinToLest_WhenInstructionIsL(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveNorth(coordinate)

	spin := moveEast.Spin('L')

	assert.Equal(t, 1, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
	fmt.Println(reflect.TypeOf(spin))
	assert.Equal(t, reflect.TypeOf(NewMoveWest(coordinate)), reflect.TypeOf(spin))
}
