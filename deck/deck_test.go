package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of deck which is a type of string

type Deck []string


func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.ToString()), 0666)
}