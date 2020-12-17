package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


// Function type to be called within the generateGrid loops
type cellValueGetter func(*[][][]string, int, int, int) string


func main() {

	cellsInitialState, err := parseInitialStateFile("solutions/17/files/initial_state.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var cellsGrid [][][]string
	var iterations = 6

	// Save the initial state matrix at Z=0
	cellsGrid = append(cellsGrid, cellsInitialState)

	// The algorithm adds padding to allowing alive cells to propagate,
	// independently of whether the active cells expand or not,
	// making the complexity cubic...
	for i := 0; i < iterations; i ++ {
		nextGrid := expandGrid(cellsGrid)
		nextGrid  = generateGrid(nextGrid, getEvolutionCell)
		cellsGrid = nextGrid
	}

	numAliveCells := countAliveCellsTotal(cellsGrid)
	fmt.Printf("The number of alive cells after %d iterations is: %d\n", iterations, numAliveCells)
}


func parseInitialStateFile(filePath string) ([][]string, error) {

	var initialState [][]string

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cells := strings.Split(line, "")
		initialState = append(initialState, cells)
	}

	return initialState, nil
}


func expandGrid(currentGrid [][][]string) [][][]string {

	// Add space for padding matrices / rows / columns
	paddedGridDepth  := len(currentGrid)+2
	paddedGridHeight := len(currentGrid[0])+2
	paddedGridWidth  := len(currentGrid[0][0])+2

	var paddedGrid [][][]string
	for d := 0; d < paddedGridDepth; d++ {

		if (d == 0) || (d == paddedGridDepth-1)  {
			emptyMatrix := generateEmptyMatrix(paddedGridWidth, paddedGridHeight, ".")
			paddedGrid   = append(paddedGrid, emptyMatrix)
			continue
		}

		var paddedMatrix [][]string
		for r := 0; r < paddedGridHeight; r++ {

			if (r == 0) || (r == paddedGridHeight-1) {
				emptyArray  := generateEmptyArray(paddedGridWidth, ".")
				paddedMatrix = append(paddedMatrix, emptyArray)
				continue
			}

			var paddedArray []string
			for c := 0; c < paddedGridWidth; c++ {

				if (c == 0) || (c == paddedGridWidth-1) {
					emptyColumn := "."
					paddedArray = append(paddedArray, emptyColumn)
					continue
				}

				paddedArray = append(paddedArray, currentGrid[d-1][r-1][c-1])
			}

			paddedMatrix = append(paddedMatrix, paddedArray)
		}

		paddedGrid = append(paddedGrid, paddedMatrix)
	}

	return paddedGrid
}


func generateGrid(currentGrid [][][]string, getterFunc cellValueGetter) [][][]string {

	var newGrid [][][]string
	var gridDepth = len(currentGrid)

	for d := 0; d < gridDepth; d++ {
		newMatrix := generateMatrix(currentGrid, getterFunc, d)
		newGrid    = append(newGrid, newMatrix)
	}

	return newGrid
}


func generateMatrix(currentGrid [][][]string, getterFunc cellValueGetter, d int) [][]string {

	var newMatrix [][]string
	var gridHeight = len(currentGrid[0])

	for r := 0; r < gridHeight; r++ {
		newArray := generateArray(currentGrid, getterFunc, d, r)
		newMatrix = append(newMatrix, newArray)
	}

	return newMatrix
}


func generateEmptyMatrix(dimX int, dimY int, fillingChar string) [][]string {

	var emptyMatrix [][]string

	for r := 0; r < dimY; r++ {
		emptyArray := generateEmptyArray(dimX, fillingChar)
		emptyMatrix = append(emptyMatrix, emptyArray)
	}

	return emptyMatrix
}


func generateArray(currentGrid [][][]string, getterFunc cellValueGetter, d int, r int) []string {

	var newArray []string
	var gridWidth = len(currentGrid[0][0])

	for c := 0; c < gridWidth; c++ {
		newArray = append(newArray, getterFunc(&currentGrid, d, r, c))
	}

	return newArray
}


func generateEmptyArray(dimX int, fillingChar string) []string {

	var emptyArray []string

	for c := 0; c < dimX; c++ {
		emptyArray = append(emptyArray, fillingChar)
	}

	return emptyArray
}


func getEvolutionCell(currentGrid *[][][]string, dep int, row int, col int) string {

	var aliveAround = countAliveCellsAround(currentGrid, dep, row, col)
	var currentCell = (*currentGrid)[dep][row][col]

	if currentCell == "#" && aliveAround != 2 && aliveAround != 3 {
		return "."
	}

	if currentCell == "." && aliveAround == 3 {
		return "#"
	}

	return currentCell
}


func countAliveCellsAround(currentGrid *[][][]string, dep int, row int, col int) int {

	gridDepth  := len(*currentGrid)
	gridHeight := len((*currentGrid)[0])
	gridWidth  := len((*currentGrid)[0][0])

	aliveCells := 0

	startDepIndex := dep-1
	startRowIndex := row-1
	startColIndex := col-1


	for d := startDepIndex; d < (startDepIndex + 3); d++ {

		// Avoid access if we go out of the grid
		if (d < 0) || (d == gridDepth)  {
			continue
		}

		for r := startRowIndex; r < (startRowIndex + 3); r++ {

			// Avoid access if we go out of the grid
			if (r < 0) || (r == gridHeight) {
				continue
			}

			for c := startColIndex; c < (startColIndex + 3); c++ {

				// Avoid access if we go out of the grid
				if (c < 0) || (c == gridWidth) {
					continue
				}

				if (*currentGrid)[d][r][c] == "#" {
					aliveCells += 1
				}
			}
		}
	}

	// Do not count our own cell
	if (*currentGrid)[dep][row][col] == "#" {
		aliveCells -= 1
	}

	return aliveCells
}


func countAliveCellsTotal(cellsGrid [][][]string) int {

	aliveCells := 0

	gridDepth  := len(cellsGrid)
	gridHeight := len(cellsGrid[0])
	gridWidth  := len(cellsGrid[0][0])

	for d := 0; d < gridDepth; d++ {
		for r := 0; r < gridHeight; r++ {
			for c := 0; c < gridWidth; c++ {
				if cellsGrid[d][r][c] == "#" {
					aliveCells += 1
				}
			}
		}
	}

	return aliveCells
}
