package models

import (
	"rover/src/constants"
	"rover/src/strategy"
)

type RoverBuilder interface {
	SetDirection(direction string) RoverBuilder
	SetCoordinate(coordinate *strategy.Coordinate) RoverBuilder
	SetPlateau(plateau *Plateau) RoverBuilder
	Build() Vehicle
}

func NewRoverBuilder() RoverBuilder {
	rover := &Rover{}
	return &RoverBuilderImp{
		rover: rover,
	}
}

type RoverBuilderImp struct {
	rover *Rover
}

func (r *RoverBuilderImp) SetDirection(direction string) RoverBuilder {
	r.rover.initialDirection = direction
	return r
}

func (r *RoverBuilderImp) SetCoordinate(coordinate *strategy.Coordinate) RoverBuilder {
	r.rover.coordinate = coordinate
	return r
}

func (r *RoverBuilderImp) SetPlateau(plateau *Plateau) RoverBuilder {
	r.rover.plateau = plateau
	return r
}

func (r *RoverBuilderImp) Build() Vehicle {
	switch r.rover.initialDirection {
	case constants.NORTH:
		r.rover.moveStrategy = strategy.NewMoveStrategy(strategy.NewMoveNorth(r.rover.coordinate))
	case constants.SOUTH:
		r.rover.moveStrategy = strategy.NewMoveStrategy(strategy.NewMoveSouth(r.rover.coordinate))
	case constants.WEST:
		r.rover.moveStrategy = strategy.NewMoveStrategy(strategy.NewMoveWest(r.rover.coordinate))
	case constants.EAST:
		r.rover.moveStrategy = strategy.NewMoveStrategy(strategy.NewMoveEast(r.rover.coordinate))
	}
	return r.rover
}
