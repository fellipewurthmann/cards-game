package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
