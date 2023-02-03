package main

import (
	"github.com/cards_project/deck"
)

func main() {
	cards := deck.NewDeck
	cards().Shuffle()
	cards().Print()
}