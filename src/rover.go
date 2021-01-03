package src

import (
	"errors"
	"fmt"
	"rover/src/models"
	"rover/src/strategy"
)

type Rover struct {
	coordinate       *models.Coordinate
	plateau          *models.Plateau
	moveStrategy     *strategy.MoveStrategy
	initialDirection string
	roverName        string
}

func (r *Rover) Move(instructions string, output chan *Rover, err chan string) {

	for _, rune := range instructions {
		if rune == 'M' {
			r.moveStrategy.GetCriteria().Move()
		} else {
			r.SetMoveCriteria(r.moveStrategy.GetCriteria().Spin(rune))
		}
	}
	error := r.checkBoundaries()

	if error != nil {
		err <- error.Error()
	} else {
		output <- r
	}
}

func (r *Rover) SetMoveCriteria(criteria strategy.MoveCriteria) {
	r.moveStrategy.SetCriteria(criteria)
}

func (r *Rover) GetMoveCriteria() strategy.MoveCriteria {
	return r.moveStrategy.GetCriteria()
}

func (r *Rover) GetCurrentCoordinate() *models.Coordinate {
	return r.coordinate
}

func (r *Rover) GetRoverName() string {
	return r.roverName
}

func (r *Rover) SetRoverName(roverName string) {
	r.roverName = roverName
}

func (r *Rover) checkBoundaries() error {
	if r.coordinate.GetX() < r.plateau.LeftBottom.GetX() ||
		r.coordinate.GetX() > r.plateau.RightTop.GetX() ||
		r.coordinate.GetY() < r.plateau.LeftBottom.GetY() ||
		r.coordinate.GetY() > r.plateau.RightTop.GetY() {
		return errors.New( fmt.Sprintf("Rover %s exceeded plateau boundaries.", r.roverName))
	}
	return nil
}
