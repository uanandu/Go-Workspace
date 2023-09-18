package main

import "fmt"

// Create a new type of 'deck' which is a slice of string
// this borrows the []string type(slice).
type deck []string

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
