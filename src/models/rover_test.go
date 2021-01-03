package models

import (
	"github.com/stretchr/testify/assert"
	"rover/src/mocks"
	"rover/src/strategy"
	"testing"
)

func TestGetDirection_ShouldReturnLatestDirection_WhenCalled(t *testing.T) {

	rover := NewRoverBuilder().
		SetDirection("N").
		SetCoordinate(nil).
		SetPlateau(nil).
		Build()
	mockMoveNorth := &mocks.MockMoveNorth{}
	mockMoveNorth.On("GetDirection").Return("N")
	rover.SetMoveCriteria(mockMoveNorth)

	direction := rover.GetMoveCriteria().GetDirection()

	assert.Equal(t, "N", direction)
	mockMoveNorth.AssertNumberOfCalls(t, "GetDirection", 1)

}

func TestMove_ShouldCallMoveOrSpinMethodOfMoveCriteria_WhenCalled(t *testing.T) {

	leftBottom := &strategy.Coordinate{}
	leftBottom.SetX(0)
	leftBottom.SetY(0)
	rightTop := &strategy.Coordinate{}
	rightTop.SetX(5)
	rightTop.SetY(5)
	plateau := &Plateau{
		leftBottom: leftBottom,
		rightTop:   rightTop,
	}
	roverCoordinate := &strategy.Coordinate{}
	roverCoordinate.SetX(3)
	roverCoordinate.SetY(3)
	rover := NewRoverBuilder().
		SetDirection("N").
		SetCoordinate(roverCoordinate).
		SetPlateau(plateau).
		Build()
	mockMoveNorth := &mocks.MockMoveNorth{}
	mockMoveNorth.On("GetDirection").Return("N")
	mockMoveNorth.On("Move")
	mockMoveNorth.On("Spin")

	rover.SetMoveCriteria(mockMoveNorth)
	roverChannel := make(chan *Rover, 20)
	errorChannel := make(chan string, 20)
	rover.Move("MMMR", roverChannel, errorChannel)

	mockMoveNorth.AssertNumberOfCalls(t, "Move", 3)
	mockMoveNorth.AssertNumberOfCalls(t, "Spin", 1)
}
