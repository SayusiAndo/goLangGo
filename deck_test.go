package main

import (
	"os"
	"testing"
)

func TestNewDeckReturns16Cards(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of deck is 16, but we got %v", len(d))
	}
}

func TestFirstItem(t *testing.T) {
	d := newDeck()

	if d[0] != "SpadesofTwo" {
		t.Errorf("First item is not Spades of Two, but %v", d[0])
	}
}

func TestLastItem(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "ClubsofAce" {
		t.Errorf("Last item is not Clubs of Ace, but %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "test.txt"

	os.Remove(testFileName)

	d := newDeck()
	d.saveToFile(testFileName)

	loadedDeck := newDeckFromFile(testFileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Length must be 16, but got %v", len(loadedDeck))
	}

	os.Remove(testFileName)
}
