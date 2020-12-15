package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {

	startingNumbers, err := parseStartingNumbersFile("solutions/15/files/starting_numbers.txt")
	if err != nil {
		fmt.Printf("The starting numbers file is invalid")
		return
	}

	var targetIteration = 2020

	for i, set := range startingNumbers {
		number := computeIterationNumber(set, targetIteration)
		fmt.Printf("The iteration %dth of the set #%d is: %d\n", targetIteration, i, number)
	}
}


func parseStartingNumbersFile(filePath string) ([][]int, error) {

	var startingNumbers [][]int
	var listNums []int

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		fileLine := scanner.Text()
		strNums  := strings.Split(fileLine, ",")

		for _, str := range strNums {
			n, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				return nil, err
			}

			listNums = append(listNums, int(n))
		}

		startingNumbers = append(startingNumbers, listNums)
		listNums = nil
	}

	return startingNumbers, nil
}


func computeIterationNumber(startingSet []int, iteration int) int {

	var turnsMap = make(map[int][]int)

	var prevNum  int
	var isNewNum bool

	// Load initial values
	for i, num := range startingSet {
		isNewNum = updateTurnsMap(&turnsMap, num, i+1)
		prevNum  = num
	}

	// Generate next values until the desired iteration
	for i := len(startingSet); i < iteration; i++ {
		nextNum := genNextNumber(&turnsMap, prevNum, isNewNum)
		isNewNum = updateTurnsMap(&turnsMap, nextNum, i+1)
		prevNum  = nextNum
	}

	return prevNum
}


func updateTurnsMap(turnsMap *map[int][]int, number int, turn int) bool {

	if (*turnsMap)[number] == nil {
		(*turnsMap)[number] = []int{turn}
		return true
	} else {
		(*turnsMap)[number] = append((*turnsMap)[number], turn)
		return false
	}
}


func genNextNumber(numbersTurns *map[int][]int, prevNumber int, isNew bool) int {

	turns, ok := (*numbersTurns)[prevNumber]
	if (ok == false) || isNew {
		return 0
	}

	turnsLen := len(turns)
	lastTurn := turns[turnsLen-1]
	prevTurn := turns[turnsLen-2]

	return lastTurn - prevTurn
}
