package models

import (
	"errors"
	"fmt"
	"rover/src/constants"
	"rover/src/strategy"
)

type Rover struct {
	coordinate       *strategy.Coordinate
	plateau          *Plateau
	moveStrategy     *strategy.MoveStrategy
	initialDirection string
	roverName        string
}

func (r *Rover) Move(instructions string, roverChannel chan *Rover, errorChannel chan string) {

	for _, instruction := range instructions {
		if instruction == constants.MOVE {
			r.moveStrategy.GetCriteria().Move()
		} else {
			r.SetMoveCriteria(r.moveStrategy.GetCriteria().Spin(instruction))
		}
	}

	err := r.CheckBoundaries()

	if err != nil {
		errorChannel <- err.Error()
	} else {
		roverChannel <- r
	}
}

func (r *Rover) SetMoveCriteria(criteria strategy.MoveCriteria) {
	r.moveStrategy.SetCriteria(criteria)
}

func (r *Rover) GetMoveCriteria() strategy.MoveCriteria {
	return r.moveStrategy.GetCriteria()
}

func (r *Rover) GetCurrentCoordinate() *strategy.Coordinate {
	return r.coordinate
}

func (r *Rover) GetVehicleName() string {
	return r.roverName
}

func (r *Rover) SetVehicleName(roverName string) {
	r.roverName = roverName
}

func (r *Rover) CheckBoundaries() error {
	if r.coordinate.GetX() < r.plateau.GetLeftBottom().GetX() ||
		r.coordinate.GetX() > r.plateau.GetRightTop().GetX() ||
		r.coordinate.GetY() < r.plateau.GetLeftBottom().GetY() ||
		r.coordinate.GetY() > r.plateau.GetRightTop().GetY() {
		return errors.New(fmt.Sprintf("Rover %s exceeded the plateau boundaries.", r.roverName))
	}
	return nil
}
