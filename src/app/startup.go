package app

import (
	"fmt"
	"rover/src/models"
	"rover/src/utils"
	"strconv"
	"time"
)

func Start() {
	roverChannel := make(chan *models.Rover, 20)
	errorChannel := make(chan string, 20)

	go func() {
		for {
			select {
			case rover := <-roverChannel:
				fmt.Printf("\n*** Latest coordinates for rover %s: x:%v, y:%v direction:%s ***\n",
					rover.GetVehicleName(), rover.GetCurrentCoordinate().GetX(),
					rover.GetCurrentCoordinate().GetY(), rover.GetMoveCriteria().GetDirection())
			case err := <-errorChannel:
				fmt.Println(err)
			}
		}
	}()

	plateau := utils.ReadPlateauCoordinates()

	roverNumber := 1

	for {
		rover := utils.ReadInitialRoverState(plateau, roverNumber)
		rover.SetVehicleName(strconv.Itoa(roverNumber))

		go rover.Move(utils.ReadRoverInstruction(), roverChannel, errorChannel)
		time.Sleep(100 * time.Millisecond)
		roverNumber++
	}
}
