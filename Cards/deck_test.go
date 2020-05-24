package main

import ("testing"
"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) !=16 {
		t.Errorf("expected 20 but recieved %v", len(d))
	}
	// if d[0] != "Ace of Spades" {
	// 	t.Errorf("Expected first card to be Ace of Spades,but got %v", d[0])
	// }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) !=16 {
		t.Errorf("Expected 16 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}