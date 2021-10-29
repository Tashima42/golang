package main

//import "fmt"

func main() {
	/*
	cards := newDeck()
	cards.saveToFile("./deck.txt")
	*/
	//cards := newDeckFromFile("./decks.txt")
	//cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
	/*
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	*/
}
