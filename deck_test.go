package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 8
	if len(d) != expectedLength {
		t.Errorf("Expected length: %d; Actual: %d", expectedLength, len(d))
	}
}

func TestFirstCard(t *testing.T) {
	d := newDeck()
	expected := "Ace of Spades"
	if d[0] != expected {
		t.Errorf("Expected length: %s; Actual: %s", expected, d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	removeUnwantedFiles(filename)
	deck := newDeck()
	err := deck.saveToFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	loadedDeck := newDeckFromFile(filename)
	expected := 8

	if len(loadedDeck) != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, len(loadedDeck))
	}
	removeUnwantedFiles(filename)
}

func removeUnwantedFiles(filename string) {
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
}
