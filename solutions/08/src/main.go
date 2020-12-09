package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Instruction struct {
	operation    string
	argument     int
}


func main() {

	instructions, err := parseInstructionsFile("solutions/08/files/instructions.txt")
	if err != nil {
		fmt.Println(err)
	}

	accumulator  := 0
	currentIndex := 0
	visitedMap   := buildInstructionVisitedMap(instructions)

	for true {

		if visitedMap[currentIndex] == true {
			break
		}

		visitedMap[currentIndex] = true
		instruction := instructions[currentIndex]

		if instruction.operation == "acc" {
			currentIndex += 1
			accumulator += instruction.argument
		} else if instruction.operation == "jmp" {
			currentIndex += instruction.argument
		} else if instruction.operation == "nop" {
			currentIndex += 1
		}
	}

	fmt.Printf("The accumulator value before repeating any instruction is: %d\n", accumulator)
}


func parseInstructionsFile(filePath string) ([]Instruction, error) {

	var instructions []Instruction

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := parseInstruction(scanner.Text())
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, *i)
	}

	return instructions, nil
}


func parseInstruction(fileLine string) (*Instruction, error) {

	parts := strings.Split(fileLine, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid instructions file")
	}

	operation := parts[0]
	argument, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return nil, err
	}

	return &Instruction{
		operation: operation,
		argument: int(argument),
	}, nil
}


func buildInstructionVisitedMap(instructions []Instruction) map[int]bool {

	visitedMap := make(map[int]bool)
	for index, _ := range instructions {
		visitedMap[index] = false
	}

	return visitedMap
}
