package main


type ArrayElement struct {
	index         int
	value         int
}

type CircularArray struct {
	index          int
	cards        []int
}


func NewCircularArray(cards []int) CircularArray {

	return CircularArray{
		cards: cards,
		index: 0,
	}
}


func (a CircularArray) getIncrementalIndex(index int) int {
	return (index + 1) % len(a.cards)
}


func (a CircularArray) getPickUpCards(numCards int) []ArrayElement {

	var pickedCards []ArrayElement

	for i := 1; i <= numCards; i++ {
		index := (a.index + i) % len(a.cards)
		card  := a.cards[index]
		elem  := ArrayElement{index: index, value: card}
		pickedCards = append(pickedCards, elem)
	}

	return pickedCards
}


func (a CircularArray) getDestinationCard(avoidedCards []ArrayElement) ArrayElement {

	var currentCard = a.cards[a.index]
	var targetCard = currentCard

	for true {
		targetCard -= 1
		if targetCard == 0 {
			targetCard = len(a.cards)
		}

		for i, card := range a.cards {
			if card != targetCard {
				continue
			}

			// Check if it is not on the already picked up cards
			var found = checkValueInArrayElements(avoidedCards, card)
			if found == false {
				return ArrayElement{index: i, value: card}
			}
		}
	}

	return ArrayElement{}
}
