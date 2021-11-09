package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards.print()
}

// Learned about package main, imports, data types, functions, receivers
