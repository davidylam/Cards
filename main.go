package main

import "fmt"

func main() {
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spacdes")

	cards := newDeck()

	//cards.print()
	hand, remaining := deal(cards, 2)
	hand.print()
	fmt.Println("=========")
	remaining.print()

}

func newCard() string {
	return "Five of Diamonds"
}
