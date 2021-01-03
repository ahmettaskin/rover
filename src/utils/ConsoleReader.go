package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"rover/src"
	"rover/src/models"
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

	leftBottom := new(models.Coordinate)
	leftBottom.SetX(0)
	leftBottom.SetY(0)

	rightTop := new(models.Coordinate)
	x, _ := strconv.Atoi(words[0])
	y, _ := strconv.Atoi(words[1])
	rightTop.SetX(x)
	rightTop.SetY(y)

	return &models.Plateau{
		LeftBottom: leftBottom,
		RightTop:   rightTop,
	}
}

func ReadInitialRover(plateau *models.Plateau, roverNumber int) *src.Rover {

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

	coordinate := new(models.Coordinate)
	x, _ := strconv.Atoi(words[0])
	y, _ := strconv.Atoi(words[1])
	coordinate.SetX(x)
	coordinate.SetY(y)

	return src.NewRoverBuilder().
		SetPlateau(plateau).
		SetCoordinate(coordinate).
		SetDirection(words[2]).
		Build()
}

func ReadInstructions() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Provide initial instructions for the rover: (exp: LLRRMM) ")

	for {
		input, _, _ := reader.ReadLine()
		matched, _ := regexp.MatchString(`^[RLM]+$`, string(input))

		if matched {
			return string(input)
		}
		fmt.Print("Invalid format! Please provide instructions for the rover: (exp: LLRRMM) ")
	}
}
