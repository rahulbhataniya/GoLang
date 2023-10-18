package main

import "fmt"

func main() {
	//var card string="Ace of cards"
	card := newCard()
	fmt.Println(card)
	printState();
}

func newCard() string {
	return "value returned from function"
}
