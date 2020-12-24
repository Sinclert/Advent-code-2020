package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type TilePosition struct {
	X             int
	Y             int
}


func main() {

	references, err := parseReferenceFile("solutions/24/files/tiles_references.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var positions = make(map[string]int)

	for _, ref := range references {
		directions   := readTileDirections(ref)
		tilePosition := convertDirectionsToPosition(directions)
		positionID   := convertPositionToString(tilePosition)

		positions[positionID] += 1
	}

	var tileFlips = reduceCountersByNumber(positions)
	fmt.Printf("There will be %d white tiles\n", tileFlips[0])
	fmt.Printf("There will be %d black tiles\n", tileFlips[1])
}



func parseReferenceFile(filePath string) ([]string, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var tiles []string

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tiles = append(tiles, line)
	}

	return tiles, nil
}


func readTileDirections(reference string) []string {

	var directions []string
	var currentDir string

	var referenceChars = strings.Split(reference, "")

	for _, char := range referenceChars {

		if char == "n" {
			currentDir = char
		}

		if char == "s" {
			currentDir = char
		}

		if char == "e" {
			currentDir += char
			directions = append(directions, currentDir)
			currentDir = ""
		}

		if char == "w" {
			currentDir += char
			directions = append(directions, currentDir)
			currentDir = ""
		}
	}

	return directions
}


func convertDirectionsToPosition(directions []string) TilePosition {

	var X = 0
	var Y = 0

	for _, dir := range directions {

		if dir == "e" {
			X -= 2
			continue
		}

		if dir == "w" {
			X += 2
			continue
		}

		if dir == "ne" {
			X -= 1
			Y += 1
			continue
		}

		if dir == "nw" {
			X += 1
			Y += 1
			continue
		}

		if dir == "se" {
			X -= 1
			Y -= 1
			continue
		}

		if dir == "sw" {
			X += 1
			Y -= 1
			continue
		}
	}

	return TilePosition{X: X, Y: Y}
}


func convertPositionToString(position TilePosition) string {
	return fmt.Sprintf("%d_%d", position.X, position.Y)
}


func reduceCountersByNumber(counters map[string]int) map[int]int {

	var counter = make(map[int]int)
	for _, num := range counters {
		flipsNum := num % 2
		counter[flipsNum] += 1
	}

	return counter
}
