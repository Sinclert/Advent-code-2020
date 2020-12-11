package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)


// Function type to be called within the generateGrid loops
type seatValueGetter func(*[][]string, int, int) string


func main() {

	seatsGrid, err := parseSeatsGridFile("solutions/11/files/seats_grid.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for true {
		nextGrid := generateGrid(seatsGrid, getEvolutionSeat)
		if equal := reflect.DeepEqual(seatsGrid, nextGrid); equal {
			break
		} else {
			seatsGrid = generateGrid(nextGrid, getOriginalSeat)
		}
	}

	numOccupiedSeats := countOccupiedSeatsTotal(seatsGrid)
	fmt.Printf("The number of occupied seats when stabilised is: %d\n", numOccupiedSeats)
}


func parseSeatsGridFile(filePath string) ([][]string, error) {

	var seatsGrid [][]string

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row  := strings.Split(line, "")
		seatsGrid = append(seatsGrid, row)
	}

	return seatsGrid, nil
}


func generateGrid(seatsGrid [][]string, getterFunc seatValueGetter) [][]string {

	var newGrid [][]string

	gridHeight := len(seatsGrid)
	gridWidth  := len(seatsGrid[0])

	for r := 0; r < gridHeight; r++ {
		var newRow = make([]string, gridWidth)

		for c := 0; c < gridWidth; c++ {
			newRow[c] = getterFunc(&seatsGrid, c, r)
		}

		newGrid = append(newGrid, newRow)
	}

	return newGrid
}


func getOriginalSeat(seatsGrid *[][]string, col int, row int) string {

	return (*seatsGrid)[row][col]
}


func getEvolutionSeat(seatsGrid *[][]string, col int, row int) string {

	currentSeat := (*seatsGrid)[row][col]
	if currentSeat == "." {
		return "."
	}

	occupiedAround := countOccupiedSeatsAround(seatsGrid, col, row)

	if currentSeat == "L" && occupiedAround == 0 {
		return "#"
	}

	if currentSeat == "#" && occupiedAround >= 4 {
		return "L"
	}

	return currentSeat
}


func countOccupiedSeatsAround(seatsGrid *[][]string, col int, row int) int {

	gridHeight := len(*seatsGrid)
	gridWidth  := len((*seatsGrid)[0])

	occupiedSeats := 0
	startColIndex := col-1
	startRowIndex := row-1

	for r := startRowIndex; r < (startRowIndex + 3); r++ {

		// Avoid access if we go out of the grid
		if (r < 0) || (r == gridHeight)  {
			continue
		}

		for c := startColIndex; c < (startColIndex + 3); c++ {

			// Avoid access if we go out of the grid
			if (c < 0) || (c == gridWidth) {
				continue
			}

			if (*seatsGrid)[r][c] == "#" {
				occupiedSeats += 1
			}
		}
	}

	// Do not count our own cell
	if (*seatsGrid)[row][col] == "#" {
		occupiedSeats -= 1
	}

	return occupiedSeats
}


func countOccupiedSeatsTotal(seatsGrid [][]string) int {

	occupiedSeats := 0

	gridHeight := len(seatsGrid)
	gridWidth  := len(seatsGrid[0])

	for r := 0; r < gridHeight; r++ {
		for c := 0; c < gridWidth; c++ {

			if seatsGrid[r][c] == "#" {
				occupiedSeats += 1
			}
		}
	}

	return occupiedSeats
}
