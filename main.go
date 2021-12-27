package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 3)

	hand.print()
	fmt.Println("------")
	remainingCards.print()

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
