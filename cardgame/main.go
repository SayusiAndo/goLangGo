package main

func main() {
	cards := newDeckFromFile("deck.txt")
	cards.shuffle()
	cards.print()
}
