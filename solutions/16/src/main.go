package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type ValuesRange struct {
	min          int
	max          int
}

type FieldRule struct {
	field       string
	ranges      []ValuesRange
}


func main() {

	fieldsRules, ticketValues, err := parseTicketValuesFile("solutions/16/files/ticket_values.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var invalidValues []int

	for _, ticket := range ticketValues {
		for _, val := range ticket {

			var isValid = checkValidValue(fieldsRules, val)
			if isValid == false {
				fmt.Printf("Invalid value: %d\n", val)
				invalidValues = append(invalidValues, val)
			}
		}
	}

	var sumError = 0
	for _, value := range invalidValues {
		sumError += value
	}

	fmt.Printf("The scanning error rate is: %d\n", sumError)
}


func parseTicketValuesFile(filePath string) ([]FieldRule, [][]int, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	fieldsRules, err := parseFieldRules(scanner)
	if err != nil {
		return nil, nil, err
	}

	// Ignore user own ticket
	for scanner.Scan() {
		if scanner.Text() == "nearby tickets:" {
			break
		}
	}

	ticketsValues, err := parseTicketVals(scanner)
	if err != nil {
		return nil, nil, err
	}

	return fieldsRules, ticketsValues, nil
}


func parseFieldRules(scanner *bufio.Scanner) ([]FieldRule, error) {

	var fieldsRules []FieldRule

	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			break
		}

		var field = strings.Split(line, ": ")
		var rules = strings.Split(field[1], " or ")
		var valueRanges []ValuesRange

		for _, r := range rules {
			ranges := strings.Split(r, "-")
			minVal, _ := strconv.ParseInt(ranges[0], 10, 32)
			maxVal, _ := strconv.ParseInt(ranges[1], 10, 32)
			valueRanges = append(valueRanges, ValuesRange{min: int(minVal), max: int(maxVal)})
		}

		fieldsRules = append(fieldsRules, FieldRule{field: field[0], ranges: valueRanges})
	}

	return fieldsRules, nil
}


func parseTicketVals(scanner *bufio.Scanner) ([][]int, error) {

	var ticketValues [][]int

	for scanner.Scan() {
		var line = scanner.Text()
		var valuesStr = strings.Split(line, ",")
		var valuesInt []int

		for _, valStr := range valuesStr {
			valInt, _ := strconv.ParseInt(valStr, 10, 32)
			valuesInt  = append(valuesInt, int(valInt))
		}

		ticketValues = append(ticketValues, valuesInt)
	}

	return ticketValues, nil
}


func checkValidValue(fieldsRules []FieldRule, val int) bool {

	var validValue = false

	for _, rule := range fieldsRules {
		for _, r := range rule.ranges {
			if val >= r.min && val <= r.max {
				validValue = true
				break
			}
		}

		if validValue {
			break
		}
	}

	return validValue
}
