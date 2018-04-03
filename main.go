package main

import "fmt"

func main() {
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spacdes")

	cards := newDeck()

	//cards.print()
	//cards.saveToFile("my_card")
	cards = newDeckFromFile("my_card")
	fmt.Println("--------------")
	for _, v := range cards {
		fmt.Println(v)
	}
}
