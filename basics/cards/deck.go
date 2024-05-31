package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// * Create a new type of 'deck'
// * This borrows all the behavior of a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"} // , "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// _ We don't need the index, so we use the blank identifier _
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// & This is a receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	s := strings.Split(string(bs), ",")

	return deck(s), nil
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		randomIndex := r.Intn(len(d) - 1)

		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}
