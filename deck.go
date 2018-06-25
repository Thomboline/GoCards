package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/***
	Create a new type of 'deck' which is a slice of strings
**/
type deck []string

//Receiver function. Receiver is of type deck
//By convention, use a one-letter abreviation of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//Using _ as an index variable tells Go that you don't want to actually use the index variable
	//Go will display an error if the index variable is not used (unless it's an _)
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
