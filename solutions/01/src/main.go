package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)


func main() {

	numbersList, err := parseNumbersFile("solutions/01/files/numbers.txt")
	if err != nil {
		log.Fatal(err)
	}

	var numbersSet = createNumbersSet(numbersList)
	var currentYear = 2020

	for _, number := range numbersList {
		remaining := currentYear - number
		isPresent := numbersSet[remaining]
		if isPresent {
			fmt.Printf("Result: %d\n", number * remaining)
			return
		}
	}

	fmt.Printf("Result not found\n")
}


func createNumbersSet (numbers []int) map[int]bool {

	numbersSet := make(map[int]bool)
	for _, num := range numbers {
		numbersSet[num] = true
	}
	return numbersSet
}


func parseNumbersFile (filePath string) ([]int, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var numbers []int
	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, int(number))
	}

	return numbers, nil
}
