package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//hand, remainingCards := deal(cards, 3)
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	// Http requests
	getPersonInfo()
}
