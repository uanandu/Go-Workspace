package main

func main() {

	cards := newDeck()

	deal(cards, 5)
}

func newCard() string {
	return "Five of Diamonds"
}
