package main

import (
	"fmt"
	"rover/src"
	"rover/src/utils"
	"strconv"
	"time"
)

func main() {

	output := make(chan *src.Rover, 20)
	err := make(chan string, 20)

	go func() {
		for {
			select {
			case out := <-output:
				fmt.Printf("\n*** Latest coordinates for rover %s: x:%v, y:%v direction:%s ***\n",
					out.GetRoverName(), out.GetCurrentCoordinate().GetX(),
					out.GetCurrentCoordinate().GetY(), out.GetMoveCriteria().GetDirection())
			case err := <-err:
				fmt.Println(err)
			}
		}
	}()

	plateau := utils.ReadPlateauCoordinates()

	roverNumber := 1
	for {

		rover := utils.ReadInitialRover(plateau, roverNumber)
		rover.SetRoverName(strconv.Itoa(roverNumber))

		go rover.Move(utils.ReadInstructions(), output, err)
		time.Sleep(1 * time.Second)
		roverNumber++
	}
}
