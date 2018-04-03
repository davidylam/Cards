package main

func main() {
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spacdes")

	cards := newDeck()

	//cards.print()
	//cards.saveToFile("my_card")
	cards.shuffle()
	cards.print()
}
