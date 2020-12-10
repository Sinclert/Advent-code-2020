package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {

	codeNumbers, err := parseNumbersFile("solutions/09/files/numbers.txt")
	if err != nil {
		fmt.Println("invalid code numbers")
	}

	preambleLength := 5
	numbersToCheck := codeNumbers[preambleLength:]

	for i, num := range numbersToCheck {
		preambleStart  := i
		preambleEnding := i + preambleLength
		preambleSlice  := codeNumbers[preambleStart:preambleEnding]

		var isValid = checkValidNumber(num, preambleSlice)
		if isValid == false {
			fmt.Printf("The first invalid number is: %d\n", num)
			return
		}
	}

	fmt.Printf("All code numbers are valid\n")
}


func parseNumbersFile(filePath string) ([]int, error) {

	var numbers []int

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, int(num))
	}

	return numbers, nil
}


func checkValidNumber(targetNumber int, preamble []int) bool {

	preambleSet := buildNumbersSet(preamble)

	for _, num := range preamble {

		// Ignore when the same number would be used twice
		remain := targetNumber - num
		if remain == num {
			continue
		}

		// Check that the remaining value is present
		_, ok := preambleSet[remain]
		if ok {
			return true
		}
	}

	return false
}


func buildNumbersSet(numbers []int) map[int]bool {

	numbersSet := make(map[int]bool)
	for _, num := range numbers {
		numbersSet[num] = true
	}

	return numbersSet
}
