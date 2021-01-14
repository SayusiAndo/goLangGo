package main

import (
	"math/rand"
	"time"
)

type deck []card

type card struct {
	suit  string
	value string
}

func newDeck() deck {
	deck := []card{}

	suits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}

	values := []string{
		"Two",
		"Three",
		"Four",
		"Ace",
	}

	for _, suit := range suits {
		for _, value := range values {
			singleCard := card{
				suit:  suit,
				value: value,
			}
			deck = append(deck, singleCard)
		}
	}

	return deck
}

func (d deck) shuffle() {
	var nanoUnixTime = time.Now().UnixNano()
	var randomSource = rand.NewSource(nanoUnixTime)
	var random = rand.New(randomSource)

	for i, _ := range d {
		var newKey = random.Intn(len(d) - 1)
		d[i], d[newKey] = d[newKey], d[i]
	}
}

func (d deck) print() {
	for _, card := range d {
		println(card.value, " of ", card.suit)
	}
}

func (c card) print() {
	println("Suit:", c.suit, " value:", c.value)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
