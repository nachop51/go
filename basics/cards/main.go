package main

import (
	"log"
)

// This works because we are declaring a new variable
// ^ var card string = "Ace of Spades"

// This doesn't work, is a non-declaration statement outside function body
// ^ card := "Ace of Spades"

func main() {
	// & Equivalent expressions:
	// ^ var card string = "Ace of Spades"
	// ^ card := "Ace of Spades"

	// reassigning the value of card
	// ^ card = "Five of Diamonds"

	// This won't work because we have already declared card as a string, so we can't redeclare it
	// card := "Five of Diamonds"

	// card := getCard()

	// fmt.Println(card)

	// Multiple cards

	// var cards deck = deck{"Ace of spaces", getCard()}
	// cards = append(cards, "Two of hearts")

	// var cards deck = newDeck()
	// cards.print()

	// hand, remainingCards := deal(cards, 8)

	// hand.print()
	// fmt.Println("-----------------------")
	// remainingCards.print()

	// * Type conversion
	// fmt.Println([]byte("Hi there!"))

	var filename string = "my_cards"

	// Save to file
	// cards.saveToFile(filename)

	newDeck, err := loadDeckFromFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	// newDeck.print()

	newDeck.shuffle()

	newDeck.print()
}

// & Function declaration
// func getCard() string {
// 	return "Eight of pikes"
// }
