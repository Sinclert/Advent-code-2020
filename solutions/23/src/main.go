package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {

	cards, err := parseCupsFile("solutions/23/files/cups_order.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var circle = NewCircularArray(cards)

	for iter := 0; iter < 100; iter++ {
		circle.cards = computeNextOrder(circle)
		circle.index = circle.getIncrementalIndex(circle.index)

		fmt.Printf("-- Move %d --\n", iter+1)
		fmt.Printf("Circle of cards: %v\n", circle.cards)
		fmt.Printf("Current card: %d\n", circle.cards[circle.index])
	}

	fmt.Printf("--------------\n")
	fmt.Printf("The cards values after label 1 are: %v\n", getLabelsAfterValue(circle.cards, 1))
}


func parseCupsFile(filePath string) ([]int, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	var cardsStr = scanner.Text()
	var cardsStrArr = strings.Split(cardsStr, "")
	var cardsIntArr []int

	for _, cardStr   := range cardsStrArr {
		cardInt, err := strconv.ParseInt(cardStr, 10, 32)
		if err != nil {
			return nil, err
		}

		cardsIntArr = append(cardsIntArr, int(cardInt))
	}

	return cardsIntArr, nil
}


func computeNextOrder(circle CircularArray) []int {

	newCardsList  := make([]int, len(circle.cards))
	pickedUpCards := circle.getPickUpCards(3)
	destinateCard := circle.getDestinationCard(pickedUpCards)

	var insertions = 0
	var readIndex  = circle.index
	var writeIndex = circle.index

	for insertions < len(newCardsList) {

		// If the original circle holds a picked up card: skip
		if checkIndexInArrayElements(pickedUpCards, readIndex) {
			readIndex = circle.getIncrementalIndex(readIndex)
			continue
		}

		// Always add the corresponding card
		newCardsList[writeIndex] = circle.cards[readIndex]

		// If the original circle holds the destinate card: insert the picked up cards
		if readIndex == destinateCard.index {
			for _, card := range pickedUpCards {
				writeIndex = circle.getIncrementalIndex(writeIndex)
				newCardsList[writeIndex] = card.value
				insertions += 1
			}
		}

		readIndex = circle.getIncrementalIndex(readIndex)
		writeIndex = circle.getIncrementalIndex(writeIndex)
		insertions += 1
	}

	return newCardsList
}


func checkIndexInArrayElements(elements []ArrayElement, index int) bool {

	for _, elem := range elements {
		if elem.index == index {
			return true
		}
	}

	return false
}


func checkValueInArrayElements(elements []ArrayElement, value int) bool {

	for _, elem := range elements {
		if elem.value == value {
			return true
		}
	}

	return false
}


func getLabelsAfterValue(cards []int, targetCard int) []int {

	var values []int
	var index    int

	for i, card := range cards {
		if card == targetCard {
			index = i
		}
	}

	for i := index; i < len(cards); i++ {
		newIndex := (i + 1) % len(cards)
		values = append(values, cards[newIndex])
	}

	return values
}
