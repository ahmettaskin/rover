package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"rover/src/models"
	"rover/src/strategy"
	"strconv"
	"strings"
)

func ReadPlateauCoordinates() *models.Plateau {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Provide right top coordinate of the Plateau: (exp: 5 5) ")

	coordinateStr := ""

	for {
		input, _, _ := reader.ReadLine()
		matched, _ := regexp.MatchString(`^\d+?\s\d+$`, string(input))

		if matched {
			coordinateStr = string(input)
			break
		}
		fmt.Print("Invalid format! Please provide right top coordinate of the Plateau: (exp: 5 5) ")
	}

	words := strings.Fields(coordinateStr)

	leftBottom := new(strategy.Coordinate)
	leftBottom.SetX(0)
	leftBottom.SetY(0)

	rightTop := new(strategy.Coordinate)
	x, _ := strconv.Atoi(words[0])
	y, _ := strconv.Atoi(words[1])
	rightTop.SetX(x)
	rightTop.SetY(y)
	plateau := &models.Plateau{}
	plateau.SetLeftBottom(leftBottom)
	plateau.SetRightTop(rightTop)
	return plateau
}

func ReadInitialRoverState(plateau *models.Plateau, roverNumber int) models.Vehicle {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nProvide initial position of the rover %d: (exp: 1 2 N) ", roverNumber)

	coordinateStr := ""

	for {
		input, _, _ := reader.ReadLine()
		matched, _ := regexp.MatchString(`^\d+?\s\d+?\s[NSEW]$`, string(input))

		if matched {
			coordinateStr = string(input)
			break
		}
		fmt.Print("Invalid format! Please provide initial position of the rover: (exp: 1 2 N) ")

	}

	words := strings.Fields(coordinateStr)

	coordinate := new(strategy.Coordinate)
	x, _ := strconv.Atoi(words[0])
	y, _ := strconv.Atoi(words[1])
	coordinate.SetX(x)
	coordinate.SetY(y)

	return models.NewRoverBuilder().
		SetPlateau(plateau).
		SetCoordinate(coordinate).
		SetDirection(words[2]).
		Build()
}

func ReadRoverInstruction() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Provide the instructions for the rover: (exp: LLRRMM) ")

	for {
		input, _, _ := reader.ReadLine()
		matched, _ := regexp.MatchString(`^[RLM]+$`, string(input))

		if matched {
			return string(input)
		}
		fmt.Print("Invalid format! Please provide the instructions for the rover: (exp: LLRRMM) ")
	}
}
