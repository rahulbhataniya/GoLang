package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16 but we got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spade but got %s", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf(" Expected length 16 but got %d", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
