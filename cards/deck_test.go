package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) < 20 {
		t.Errorf("Expected deck of at least 20 cards, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card to be Five of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndLoadDeckFromFile(t *testing.T) {
	var filename = "_decktesting"

	os.Remove(filename)

	deck := newDeck()

	deck.saveToFile(filename)

	loadedDeck, err := loadDeckFromFile(filename)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(filename)

	_, err = loadDeckFromFile(filename)

	if err == nil {
		t.Errorf("Expected an error, but got none.")
	}
}

func TestDeal(t *testing.T) {
	deck := newDeck()

	hand, remainingCards := deal(deck, 5)

	if len(hand) != 5 {
		t.Errorf("Expected 5 cards in hand, got %v", len(hand))
	}

	if len(remainingCards) != 15 {
		t.Errorf("Expected 15 cards in remaining deck, got %v", len(remainingCards))
	}
}

func TestPrint(t *testing.T) {
	deck := newDeck()

	deck.print()
}

func TestShuffle(t *testing.T) {
	deck := newDeck()

	deck.shuffle()
}
