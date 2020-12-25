package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)


var cardPublicKeyRegex = regexp.MustCompile(`card: (\d+)`)
var doorPublicKeyRegex = regexp.MustCompile(`door: (\d+)`)


func main() {

	publicKeys, err := parsePublicKeysFile("solutions/25/files/public_keys.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var subjectNum = 7
	var maximumNum = 20201227
	var loopSizes  = make([]int, len(publicKeys))

	for i, key := range publicKeys {
		loopSize := inferLoopSize(key, subjectNum, maximumNum)
		loopSizes[i] = loopSize
	}

	var encryptionKeyA = calcEncryptionKey(publicKeys[0], loopSizes[1], maximumNum)
	var encryptionKeyB = calcEncryptionKey(publicKeys[1], loopSizes[0], maximumNum)
	if encryptionKeyA != encryptionKeyB {
		fmt.Println("Something went wrong on the decrypting procedure")
		return
	}

	fmt.Printf("The encryption key is: %d\n", encryptionKeyB)
}


func parsePublicKeysFile(filePath string) ([]int, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var publicKeys []int

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if cardPublicKeyRegex.MatchString(line) {
			publicKeyStr    := cardPublicKeyRegex.FindStringSubmatch(line)
			publicKeyInt, _ := strconv.ParseInt(publicKeyStr[1], 10, 32)
			publicKeys = append(publicKeys, int(publicKeyInt))
		}

		if doorPublicKeyRegex.MatchString(line) {
			publicKeyStr    := doorPublicKeyRegex.FindStringSubmatch(line)
			publicKeyInt, _ := strconv.ParseInt(publicKeyStr[1], 10, 32)
			publicKeys = append(publicKeys, int(publicKeyInt))
		}
	}

	return publicKeys, nil
}


func inferLoopSize(targetNum int, subjectNum int, maxNum int) int {

	var loopSize   = 0
	var currentNum = 1

	for currentNum != targetNum {
		currentNum *= subjectNum
		currentNum %= maxNum
		loopSize += 1
	}

	return loopSize
}


func calcEncryptionKey(subjectNum int, loopSize int, maxNum int) int {

	var currentNum = 1

	for i := 0; i < loopSize; i++ {
		currentNum *= subjectNum
		currentNum %= maxNum
	}

	return currentNum
}
