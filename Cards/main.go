package main

func main() {
	// Writing to file
	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")

	// Reading file
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
