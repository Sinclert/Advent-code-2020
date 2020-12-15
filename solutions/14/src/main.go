package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)


type BitMaskReplacement struct {
	index        int
	value        int
}

type BitMask struct {
	replacements []BitMaskReplacement
	length       int
}


func main() {

	bitmask, bitmaskOps, err := parseBitOperationsFile("solutions/14/files/bitmask_ops.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	memoryMap := computeResultingMemory(bitmaskOps, *bitmask)

	totalSum := 0
	for address, value := range memoryMap {
		fmt.Printf("The memory address %d holds the value: %d\n", address, value)
		totalSum += value
	}

	fmt.Printf("The total value is: %d\n", totalSum)
}


func parseBitOperationsFile(filePath string) (*BitMask, map[int]int, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	bitMask, maskError := parseBitMask(scanner)
	bitOps, opsError   := parseBitOps(scanner)

	if (maskError != nil) || (opsError != nil) {
		return nil, nil, errors.New("the bitmask operations file is invalid")
	}

	return bitMask, bitOps, nil
}


func parseBitMask(scanner * bufio.Scanner) (*BitMask, error) {

	var replacements []BitMaskReplacement

	scanner.Scan()

	maskLines   := scanner.Text()
	maskValues  := strings.Split(maskLines, " = ")
	bitmaskBits := strings.Split(maskValues[1], "")
	bitmaskLen  := len(bitmaskBits)

	for i, bitStr := range bitmaskBits {
		if bitStr != "X" {
			bitInt, err := strconv.ParseInt(bitStr, 10, 32)
			if err != nil {
				return nil, err
			}
			replacement := BitMaskReplacement{index: i, value: int(bitInt)}
			replacements = append(replacements, replacement)
		}
	}

	return &BitMask{replacements: replacements, length: bitmaskLen}, nil
}


func parseBitOps(scanner * bufio.Scanner) (map[int]int, error) {

	var operations = make(map[int]int)

	for scanner.Scan() {
		memoryOpLine := scanner.Text()
		memoryOpValues := strings.Split(memoryOpLine, " = ")
		if len(memoryOpValues) != 2 {
			return nil, errors.New("the memory operations are not properly defined")
		}

		memOpReplIntValue, err := strconv.ParseInt(memoryOpValues[1], 10, 32)
		if err != nil {
			return nil, err
		}

		memOpIndexRegexExpr     := regexp.MustCompile(`mem\[(\d)\]`)
		memOpIndexRegexResults  := memOpIndexRegexExpr.FindAllStringSubmatch(memoryOpValues[0], 1)
		memOpIndexStrValue      := memOpIndexRegexResults[0][1]
		memOpIndexIntValue, err := strconv.ParseInt(memOpIndexStrValue, 10, 32)
		if err != nil {
			return nil, err
		}

		operations[int(memOpIndexIntValue)] = int(memOpReplIntValue)
	}

	return operations, nil
}


func computeResultingMemory(bitOperations map[int]int, bitMask BitMask) map[int]int {

	var resultingMemory = make(map[int]int)

	for index, value := range bitOperations {
		binaryArray  := decimalToBinaryArray(value, bitMask.length)
		for _, repl  := range bitMask.replacements {
			binaryArray[repl.index] = repl.value
		}

		resultingMemory[index] = binaryArrayToDecimal(binaryArray)
	}

	return resultingMemory
}


func decimalToBinaryArray(decimalNumber int, maxLength int) []int {

	binaryArray := make([]int, maxLength)

	binaryPower := float64(maxLength-1)
	remainValue := decimalNumber

	for i := 0; i < maxLength; i++ {
		binaryPowerValue := int(math.Pow(2, binaryPower))
		if (remainValue - binaryPowerValue) >= 0 {
			remainValue -= binaryPowerValue
			binaryArray[i] = 1
		} else {
			binaryArray[i] = 0
		}

		binaryPower -= 1
	}

	return binaryArray
}


func binaryArrayToDecimal(binaryArray []int) int {

	decimalNumber := 0

	binaryArrayLen := len(binaryArray)
	binaryArrayPow := float64(binaryArrayLen-1)

	for i := 0; i < binaryArrayLen; i++ {
		if binaryArray[i] == 1 {
			decimalNumber += int(math.Pow(2, binaryArrayPow))
		}
		binaryArrayPow -= 1
	}

	return decimalNumber
}
