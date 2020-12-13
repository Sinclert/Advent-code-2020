package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)


type FerryPosition struct {
	X              int
	Y              int
}

type Instruction struct {
	operation      string
	units          int
}


var allDirectionNames = [4]string{"N", "E", "S", "W"}

var movingInstructions = map[string]bool{
	"N": true,
	"E": true,
	"S": true,
	"W": true,
	"F": true,
}

var rotatingInstructions = map[string]bool{
	"L": true,
	"R": true,
}


func main() {

	instructions, err := parseInstructionsFile("solutions/12/files/moving_instructions.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	initialPosition  := FerryPosition{X: 0, Y: 0}
	currentPosition  := FerryPosition{X: 0, Y: 0}
	currentDirection := "E"

	for _, inst := range instructions {

		// Check if it is a rotating instruction
		if _, ok := rotatingInstructions[inst.operation]; ok {
			currentDirection = computeNewDirection(inst, currentDirection)
			continue
		}

		// Check if it is a moving instruction
		if _, ok := movingInstructions[inst.operation]; ok {
			currentPosition  = computeNewPosition(inst, currentPosition, currentDirection)
			continue
		}
	}

	distance := computeManhattanDistance(initialPosition, currentPosition)
	fmt.Printf("The Manhattan distance after all the moves is: %d\n", distance)
}


func parseInstructionsFile(filePath string) ([]Instruction, error) {

	var instructions []Instruction

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := buildMovingInstruction(scanner.Text())
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, i)
	}

	return instructions, nil
}


func buildMovingInstruction(instructionString string) (Instruction, error) {

	instruction := Instruction{}
	direction   := instructionString[0]
	strUnits    := instructionString[1:]

	intUnits, err := strconv.ParseInt(strUnits, 10, 32)
	if err != nil {
		return instruction, err
	}

	instruction.operation = string(direction)
	instruction.units     = int(intUnits)

	return instruction, nil
}


func computeNewDirection(instruction Instruction, prevDirection string) string {

	var directionChanges   = (instruction.units % 360) / 90
	var newDirectionIndex  = 0
	var prevDirectionIndex = 0

	if prevDirection == "N" {
		prevDirectionIndex = 0
	} else if prevDirection == "E" {
		prevDirectionIndex = 1
	} else if prevDirection == "S" {
		prevDirectionIndex = 2
	} else if prevDirection == "W" {
		prevDirectionIndex = 3
	}

	if instruction.operation == "L" {
		newDirectionIndex = (prevDirectionIndex - directionChanges) % 4
	} else if instruction.operation == "R" {
		newDirectionIndex = (prevDirectionIndex + directionChanges) % 4
	}

	return allDirectionNames[newDirectionIndex]
}


func computeNewPosition(instruction Instruction, prevPosition FerryPosition, prevDirection string) FerryPosition {

	// If Forward direction: override the instruction direction and try again
	if instruction.operation == "F" {
		instruction.operation = prevDirection
		return computeNewPosition(instruction, prevPosition, prevDirection)
	}

	// Copy the previous ferry position
	var newFerryPosition = FerryPosition{
		X: prevPosition.X,
		Y: prevPosition.Y,
	}

	if instruction.operation == "N" {
		newFerryPosition.Y += instruction.units
	} else if instruction.operation == "E" {
		newFerryPosition.X += instruction.units
	} else if instruction.operation == "S" {
		newFerryPosition.Y -= instruction.units
	} else if instruction.operation == "W" {
		newFerryPosition.X -= instruction.units
	}

	return newFerryPosition
}


func computeManhattanDistance(posA FerryPosition, posB FerryPosition) int {

	XDiff := math.Abs(float64(posA.X - posB.X))
	YDiff := math.Abs(float64(posA.Y - posB.Y))

	return int(XDiff + YDiff)
}
