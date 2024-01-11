package main

var a = 1

func main() {
	cards := deck{"wildcard", newCard()}
	cards = append(cards, newCard())

	// card := newCard()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
