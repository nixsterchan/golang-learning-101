package main

import (
	"os"
	"testing"
)

// Following file helps run some tests on the deck.go mod created. Simply run `go test deck_test.go deck.go` in the terminal to start the test.
// The testing package is utilized to help with this testing endeavour.

// Helps test the creation of a new deck based on our newDeck function
func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 52
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "King of Clubs"
	if len(d) != expectedLength {
		t.Errorf("Expected the deck to be of length %v but instead got %v.\n", expectedLength, len(d))
	}

	if d[0] != expectedFirstCard {
		t.Errorf("Expected %v to be the very first card but instead got %v.\n", expectedFirstCard, d[0])
	}

	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected %v to be the very last card but instead got %v.\n", expectedLastCard, d[len(d)-1])
	}
}

// Perform some testing on the writing and reading functionalities in the deck.go module
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Initialize the expected test file name
	testFileName := "_testingDeckFile"

	// Check to ensure and remove the test file if it still exists
	os.Remove(testFileName)

	// First create the new deck, save it out, and load it back in
	originalDeck := newDeck()
	originalDeck.saveToFile(testFileName)
	loadedDeck := newDeckFromFile(testFileName)

	// Check if their lengths align
	if len(loadedDeck) != len(originalDeck) {
		t.Errorf("Expected length of deck to be %v but got %v instead.\n", len(originalDeck), len(loadedDeck))
	}

	// Then check the positioning of the cards. By right they should match.
	for index := range originalDeck {
		if loadedDeck[index] != originalDeck[index] {
			t.Errorf("A card was found wrongly aligned in the loaded deck: %v.\n", loadedDeck[index])
		}
	}

	// Cleanup the test file
	os.Remove(testFileName)
}
