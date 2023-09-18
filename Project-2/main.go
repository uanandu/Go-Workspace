package main

func main() {

	// we declare multiple cards
	cards := deck{"Ace of DIamonds", newCard()}

	// append or add
	cards = append(cards, "Six of spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
