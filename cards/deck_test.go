package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	firstCard := card{name: "Ace", suit: "Spades"}
	lastCard := card{name: "King", suit: "Clubs"}

	if d[0] != firstCard {
		t.Errorf("Expected 'Ace of Spades' but got %v", d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("Expected 'King of Clubs' but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
