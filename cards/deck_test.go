package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckLength := len(d)

	if deckLength != 52 {
		t.Errorf("Expected deck length of 52, but got %v", deckLength)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first position to be 'Ace of Spades', but go %v", d[0])
	}

	if d[deckLength-1] != "King of Clubs" {
		t.Errorf("Expected last position to be 'King of Clubs', but go %v", d[deckLength-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
