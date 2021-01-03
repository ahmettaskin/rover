package src

import (
	"rover/src/models"
	"rover/src/strategy"
)

type RoverBuilder interface {
	SetDirection(direction string) RoverBuilder
	SetCoordinate(coordinate *models.Coordinate) RoverBuilder
	SetPlateau(plateau *models.Plateau) RoverBuilder
	Build() *Rover
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

func (r *RoverBuilderImp) SetCoordinate(coordinate *models.Coordinate) RoverBuilder {
	r.rover.coordinate = coordinate
	return r
}

func (r *RoverBuilderImp) SetPlateau(plateau *models.Plateau) RoverBuilder {
	r.rover.plateau = plateau
	return r
}

func (r *RoverBuilderImp) Build() *Rover {
	switch r.rover.initialDirection {
	case "N":
		r.rover.moveStrategy = strategy.NewMoveStrategy(strategy.NewMoveNorth(r.rover.coordinate))
	case "S":
		r.rover.moveStrategy = strategy.NewMoveStrategy(strategy.NewMoveSouth(r.rover.coordinate))
	case "W":
		r.rover.moveStrategy = strategy.NewMoveStrategy(strategy.NewMoveWest(r.rover.coordinate))
	case "E":
		r.rover.moveStrategy = strategy.NewMoveStrategy(strategy.NewMoveEast(r.rover.coordinate))
	}
	return r.rover
}
