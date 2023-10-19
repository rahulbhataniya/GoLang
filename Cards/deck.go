package main

import "fmt"

//define data type

type deck []string

func newCard() deck {
	cards := []string{"one", "two", "three","four","five"}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck,handsize int) (deck,deck){
	return d[:handsize],d[handsize:]
}
