package main

func main() {
	cards := deck{newCard(), "Diamond"}
	cards = append(cards, "new value")
	cards.print()
}

func newCard() string {
	return "value returned from function"
}
