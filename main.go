package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//hand, remainingCards := deal(cards, 3)
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	//cards.print()

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
