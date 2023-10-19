package main

func main() {
	cards := newCard()
	hand, remainingCards := deal(cards, 3)
	hand.print()
	remainingCards.print()
}
