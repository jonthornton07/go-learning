package main

import "fmt"

func main() {
	cards := newDeck()
	err := cards.saveToFile("my_cards.txt")

	if err != nil {
		fmt.Println(err)
	}

	loadedCards := newDeckFromFile("my_cards.txt")
	loadedCards.shuffle()
	fmt.Println(loadedCards.toString())
}
