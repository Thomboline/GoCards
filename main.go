package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	cards := []string{"Five of Diamonds", "Ace of Spades"}

	//When appending, Go does not just add another variable to the slice.
	//It instead creates a new slice from the specified data
	cards = append(cards, "Two of Hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
