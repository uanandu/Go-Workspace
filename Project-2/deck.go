package main

import "fmt"

// Create a new type of 'deck' which is a slice of string
// this borrows the []string type(slice).
type deck []string

func newDeck() deck {
	cards := deck{}

	// list of suites
	cardSuites := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}

	// list of values
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King"}

	// we put _ because we dont need to use the index for anything.
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			// here we append it to the cards slice
			cards = append(cards, value+" of "+suit)
		}
	}

	// we return what we created
	return cards
}

// Anything with the type 'deck' will have access to the method print()
// d-> It is the actual copy of the deck that we are working with is available
// deck -> It is the type declared above.

// Process => When Cards.print() is executed, the 'd' here gets replaced by the cards array with type 'deck'
// d = receiver argument
func (d deck) print() {
	// Iterate through the cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// here we set the deal function with 3 cards on hand
// here:
// d = replaces the cards (receiver)
// deck = []string for the deck
// handSize = the number of cards in hand

// SInce we are returning two slices from this, we return: deck (hand and reserve)
// It is important to annotate the return types
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
