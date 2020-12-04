package main

import (
	"bufio"
	"fmt"
	"os"
)


type slopePosition struct {
	col  int
	row  int
}

type slopeMovement struct {
	upperMovement int
	downMovement  int
	leftMovement  int
	rightMovement int
}


func main() {

	movement := slopeMovement{
		upperMovement: 0,
		downMovement:  1,
		leftMovement:  0,
		rightMovement: 3,

	}

	initialPos := slopePosition{
		col: 0,
		row: 0,
	}

	slopeMap, err := parseSlopeFile("solutions/03/files/mountain_slope.txt")
	if err != nil {
		fmt.Printf("Invalid text file")
	}

	trees, err := traverseSlope(slopeMap, initialPos, movement)
	if err != nil {
		fmt.Printf("Invalid file format")
	}

	fmt.Printf("Number of tress encountered: %d", trees)
}


func parseSlopeFile(filePath string) (*[]string, error) {

	var slopeMap []string

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		slopeMap = append(slopeMap, scanner.Text())
	}

	return &slopeMap, nil
}


func traverseSlope(slopeMap *[]string, init slopePosition, movement slopeMovement) (int, error) {

	trees := 0

	verticalMovement   := movement.downMovement - movement.upperMovement
	horizontalMovement := movement.rightMovement - movement.leftMovement

	currentCol := init.col
	currentRow := init.row

	for true {
		currentCol += horizontalMovement
		currentRow += verticalMovement

		if shouldStop(slopeMap, slopePosition{col: currentCol, row: currentRow}) {
			break
		}

		character := (*slopeMap)[currentRow][currentCol:currentCol+1]
		if isTree(character) {
			trees += 1
		}
	}

	return trees, nil
}


func shouldStop (slopeMap *[]string, currentPos slopePosition) bool {

	outsideVerticalBounds   := (currentPos.row < 0) || (currentPos.row >= len(*slopeMap))
	outsideHorizontalBounds := (currentPos.col < 0) || (currentPos.col >= len((*slopeMap)[0]))

	return outsideVerticalBounds || outsideHorizontalBounds
}


func isTree (character string) bool {

	if character == "#" {
		return true
	}
	return false
}
