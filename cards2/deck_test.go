package main

import "testing"

func TestDeckHas16Cards(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Deck has to have 16 cards, but it has: %+v", len(deck))
	}
}

func TestDeckHas4Spades(t *testing.T) {
	deck := newDeck()

	var counter = 0

	for _, card := range deck {
		if card.suit == "Spades" {
			counter++
		}
	}

	if counter != 4 {
		t.Errorf("Deck must have four Spades, but has %+v", counter)
	}
}

func TestDeckChangedAfterShuffle(t *testing.T) {
	deck := newDeck()
	deck.shuffle()

	if deck[0].suit == "Spades" && deck[0].value == "Two" {
		t.Errorf("Deck is not shuffled.")
	}
}

func TestDealMethod(t *testing.T) {
	deck := newDeck()

	oneDeck, anotherDeck := deal(deck, 5)

	if len(oneDeck) != 5 && len(anotherDeck) != 11 {
		t.Error("Deal doesnt work. Deal size: ", 5,
			"oneDeck size: ", len(oneDeck),
			"another deck size: ", len(anotherDeck))
	}
}
