package main

import "fmt"

//Create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "2"} // , "3", "4"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	//fmt.Println(cards)
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
