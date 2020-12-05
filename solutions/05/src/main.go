package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)


type PassCodeSection struct {
	start int
	end   int
	len   int
}


func NewPassCodeSection(start int, end int) PassCodeSection {

	return PassCodeSection{
		start: start,
		end: end,
		len: end - start,
	}
}


func main() {

	passes, err := parseBoardingPassFile("solutions/05/files/boarding_passes.txt")
	if err != nil {
		fmt.Println(err)
	}

	rowCodeSection := NewPassCodeSection(0, 7)
	colCodeSection := NewPassCodeSection(7, 10)

	for _, passCode := range *passes {
		row, _ := decodeBoardingPassRow(passCode, rowCodeSection)
		col, _ := decodeBoardingPassCol(passCode, colCodeSection)
		seatID := row * (rowCodeSection.len+1) + col

		fmt.Printf("Boarding pass %s: row %d, column %d, seat ID %d.\n", passCode, row, col, seatID)
	}
}


func parseBoardingPassFile(filePath string) (*[]string, error) {

	var passes []string

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		boardingPas := scanner.Text()
		passes = append(passes, boardingPas)
	}

	return &passes, nil
}


func decodeBoardingPassRow(boardingPassCode string, rowCodeSection PassCodeSection) (int, error) {

	boardingPassRowCode := boardingPassCode[rowCodeSection.start:rowCodeSection.end]
	boardingPassRowCmds := strings.Split(boardingPassRowCode, "")
	return doBinarySearch(boardingPassRowCmds, "B", "F")
}


func decodeBoardingPassCol(boardingPassCode string, colCodeSection PassCodeSection) (int, error) {

	boardingPassColCode := boardingPassCode[colCodeSection.start:colCodeSection.end]
	boardingPassColCmds := strings.Split(boardingPassColCode, "")
	return doBinarySearch(boardingPassColCmds, "R", "L")
}


func doBinarySearch(searchCommands []string, upPartitionChar string, downPartitionChar string) (int, error) {

	lowerLimit := math.Pow(2, 0)
	upperLimit := math.Pow(2, float64(len(searchCommands)))

	for i, char := range searchCommands {
		binarySearchPow := i+1
		binarySearchCut := len(searchCommands)-binarySearchPow

		if char == downPartitionChar {
			upperLimit -= math.Pow(2, float64(binarySearchCut))
		} else if char == upPartitionChar {
			lowerLimit += math.Pow(2, float64(binarySearchCut))
		}
	}

	// At this point both lowerLimit and upperLimit must be the same
	if lowerLimit != upperLimit {
		return 0, errors.New("incorrect binary search")
	}

	// Correcting value, as the search started at "1", instead of "0"
	return int(lowerLimit)-1, nil
}
