package src

import (
	"rover/src/models"
	"rover/src/strategy"
)

type Vehicle interface {
	Move(instructions string, output chan *models.Coordinate, err chan string)
	SetMoveCriteria(moveStrategy strategy.MoveCriteria)
}
