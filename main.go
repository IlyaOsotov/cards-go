package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 4)
	remainingCards.print()
	hand.saveToFile("cards")
}
