package main

func main() {
	cards := deck{"Ace of Diamonds", "Five of Diamonds"}
	cards.saveToFile("my_cards")

	newCards := newDeckFromFile("new_cards")
	newCards.shuffle()
	newCards.print()
}
