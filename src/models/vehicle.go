package models

import (
	"rover/src/strategy"
)

type Vehicle interface {
	Move(instructions string, latestCoordinate chan *Rover, err chan string)
	SetMoveCriteria(criteria strategy.MoveCriteria)
	GetMoveCriteria() strategy.MoveCriteria
	GetCurrentCoordinate() *strategy.Coordinate
	GetVehicleName() string
	SetVehicleName(roverName string)
}
