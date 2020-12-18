package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


var operatorRegex = regexp.MustCompile(`[+*]`)
var operandsRegex = regexp.MustCompile(`\d+`)
var parenthesisRegex = regexp.MustCompile(`\((\d+ [+*] (\d+)*)+\)`)


func main() {

	mathOperations, err := parseMathOperationsFile("solutions/18/files/math_operations.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, op := range mathOperations {
		flatExpression := flatMathExpression(op)
		resultValueStr := calcExpressionValue(flatExpression)
		fmt.Printf("The result of operation \"%s\" is: %s\n", op, resultValueStr)
	}
}


func parseMathOperationsFile(filePath string) ([]string, error) {

	var mathOperations []string

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		mathOperations = append(mathOperations, scanner.Text())
	}

	return mathOperations, nil
}


func flatMathExpression(operations string) string {

	var flatExpression = operations

	for true {
		var parenthesisExpressions = parenthesisRegex.FindAllString(flatExpression, -1)

		// Exit condition: no more parenthesis
		if len(parenthesisExpressions) == 0 {
			break
		}

		for _, expr := range parenthesisExpressions {
			startIndex := 1
			endIndex   := len(expr)-1
			expression := expr[startIndex:endIndex]

			equivalentString := calcExpressionValue(expression)
			flatExpression = strings.ReplaceAll(flatExpression, expr, equivalentString)
		}
	}

	return flatExpression
}


func calcExpressionValue(expression string) string {

	var exprParts = strings.Split(expression, " ")
	var operands  = filterExpressionOperands(exprParts)
	var operators = filterExpressionOperators(exprParts)
	var resultValue int64

	for i, val := range operands {

		if i == 0 {
			resultValue = val
			continue
		}

		var operandIndex  = i
		var operatorIndex = i-1

		if operators[operatorIndex] == "+" {
			resultValue += operands[operandIndex]
		} else if operators[operatorIndex] == "*" {
			resultValue *= operands[operandIndex]
		}
	}

	return strconv.FormatInt(resultValue, 10)
}


func filterExpressionOperands(expressionParts []string) []int64 {

	var operands []int64

	for _, element := range expressionParts {
		if operandsRegex.MatchString(element) == false {
			continue
		}

		operand, _ := strconv.ParseInt(element, 10, 64)
		operands = append(operands, operand)
	}

	return operands
}


func filterExpressionOperators(expressionParts []string) []string {

	var operators []string

	for _, element := range expressionParts {
		if operatorRegex.MatchString(element) == false {
			continue
		}

		operators = append(operators, element)
	}

	return operators
}
