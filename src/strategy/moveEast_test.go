package strategy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"rover/src/models"
	"testing"
)

func TestCourierClientImp_GetCourierBy_should_get_courierDetail(t *testing.T) {

	coordinate := new(models.Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveEast(coordinate)

	moveEast.Move()

	assert.Equal(t, 2, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
}

func TestXXX(t *testing.T) {

	coordinate := new(models.Coordinate)
	coordinate.SetX(1)
	coordinate.SetY(1)
	moveEast := NewMoveEast(coordinate)

	spin := moveEast.Spin('R')

	assert.Equal(t, 1, coordinate.GetX())
	assert.Equal(t, 1, coordinate.GetY())
	fmt.Println(reflect.TypeOf(spin))
	assert.Equal(t, reflect.TypeOf(NewMoveSouth(coordinate)), reflect.TypeOf(spin))

}
