package main

import "fmt"

func main() {
	// Both these lines are valid ways to create the variable card. Note that The second line can only be called within a function scope while
	// the first can be called outside
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"

	handSize := 5
	myHandFilePath := "maCards"

	// Start by creating a new deck
	cards := newDeck()
	// We shuffle it a lil
	cards.shuffle()
	// Then we show its contents
	cards.print()

	// Next we can do a dealing of the cards based on the hand size we want
	hand, remainingCards := deal(cards, handSize)

	// Show the current hand and the remaining cards
	fmt.Println("My current hand:")
	hand.print()
	fmt.Println("The remaining cards:")
	remainingCards.print()

	// Save my hand
	hand.saveToFile(myHandFilePath)

	// Now read my hand
	loadedHand := newDeckFromFile(myHandFilePath)
	loadedHand.print()

}
