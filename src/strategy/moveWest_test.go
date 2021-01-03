package strategy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMoveWest_ShouldMoveToWest_WhenCalled(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveWest(coordinate)

	moveEast.Move()

	assert.Equal(t, 0, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
}

func TestSpinWest_ShouldSpinToRight_WhenInstructionIsR(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveWest(coordinate)

	spin := moveEast.Spin('R')

	assert.Equal(t, 1, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
	fmt.Println(reflect.TypeOf(spin))
	assert.Equal(t, reflect.TypeOf(NewMoveNorth(coordinate)), reflect.TypeOf(spin))
}

func TestSpinWest_ShouldSpinToRight_WhenInstructionIsL(t *testing.T) {

	coordinate := new(Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveWest(coordinate)

	spin := moveEast.Spin('L')

	assert.Equal(t, 1, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
	fmt.Println(reflect.TypeOf(spin))
	assert.Equal(t, reflect.TypeOf(NewMoveSouth(coordinate)), reflect.TypeOf(spin))
}
