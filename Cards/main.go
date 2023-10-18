package main

import "fmt"

func main() {
	cards := deck{newCard(), "Diamond"}
	cards = append(cards, "new value")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "value returned from function"
}
