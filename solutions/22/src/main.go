package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)


var playerRegex = regexp.MustCompile(`Player \d+:`)


func main() {

	playerDecks, err := parsePlayerDecksFile("solutions/22/files/player_decks.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for checkDecksNotEmpty(playerDecks) {

		deckA := &playerDecks[0]
		deckB := &playerDecks[1]

		cardA := takeTopCard(deckA)
		cardB := takeTopCard(deckB)

		if cardA > cardB {
			saveBottomCards(deckA, []int{cardA, cardB})
			continue
		}

		if cardA < cardB {
			saveBottomCards(deckB, []int{cardB, cardA})
			continue
		}
	}

	winnerIndex, err := getWinnerIndex(playerDecks)
	if err != nil {
		fmt.Println(err)
		return
	}

	winnerDeck  := playerDecks[winnerIndex]
	winnerScore := countDeckScore(winnerDeck)

	fmt.Printf("The final deck of player %d is: %v\n", winnerIndex+1, winnerDeck)
	fmt.Printf("The final score of player %d is: %d\n", winnerIndex+1, winnerScore)
}


func parsePlayerDecksFile(filePath string) ([][]int, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var cardDecks [][]int
	var cardDeck  []int

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if playerRegex.MatchString(line) {
			continue
		}

		if line == "" {
			cardDecks = append(cardDecks, cardDeck)
			cardDeck = nil
			continue
		}

		n, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			return nil, err
		}

		cardDeck = append(cardDeck, int(n))
	}

	cardDecks = append(cardDecks, cardDeck)
	return cardDecks, nil
}


func takeTopCard(cardDeck *[]int) int {

	topCard  := (*cardDeck)[0]
	newDeck  := (*cardDeck)[1:]
	*cardDeck = newDeck

	return topCard
}


func saveBottomCards(cardDeck *[]int, cards []int) {

	var newDeck  = *cardDeck
	for _, card := range cards {
		newDeck = append(newDeck, card)
	}

	*cardDeck = newDeck
}


func checkDecksNotEmpty(cardDecks [][]int) bool {

	for _, deck := range cardDecks {
		if len(deck) == 0 {
			return false
		}
	}

	return true
}


func getWinnerIndex(cardDecks [][]int) (int, error) {

	for i, deck := range cardDecks {
		if len(deck) > 0 {
			return i, nil
		}
	}

	return 0, errors.New("no player has won")
}


func countDeckScore(cardDeck []int) int {

	var sumScore = 0
	var lenDeck  = len(cardDeck)

	for i, card := range cardDeck {
		sumScore += card * (lenDeck-i)
	}

	return sumScore
}
