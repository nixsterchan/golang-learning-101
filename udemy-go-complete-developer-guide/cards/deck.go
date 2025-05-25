package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type called 'deck' which represents a slice of string type values
type deck []string

// Build a typical poker cards deck
func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Utilize _ since we do not need the index, else go will go and complain
	for _, suit := range cardSuit {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Function with deck receiver to simply print cards within a deck type variable
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Converts the deck from string slice to a single string concated with "," and using the strings package
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Writes out the deck as a string to a file with the given filename. Basically uses the os package, where the
// deck type is converted to a string, then to a slice of bytes and then passed into the WriteFile function. Lastly the permission type is stated
func (d deck) saveToFile(filename string) error {
	fmt.Printf("Saving file to path:%v.\n", filename)
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Shuffles the deck with help from the rand package
func (d deck) shuffle() {
	// In order to ensure true randomness, a source seed a first created. The current time is used since it is an
	// easily available number that changes all the "time", get it? ok
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// After setting the random seed, then just iterate through the indexes in the deck and do a bunch of position swaps
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	fmt.Println("Deck has been shuffled.")
}

// Function that utilizes slices to return part of the deck depending on the provided hand size and the rest of the remaining cards
func deal(d deck, handSize int) (deck, deck) {
	fmt.Printf("Dealing a hand of size %v.\n", handSize)
	return d[:handSize], d[handSize:]
}

// This function will help with reading out a string from a provided file, and convert it into a deck if properly formatted
func newDeckFromFile(filename string) deck {
	fmt.Printf("Attempting to read and retrieve a deck from the file: %v.\n", filename)
	// First read out the byte slice via os package's readfile
	bs, err := os.ReadFile(filename)

	// Handle the error depending on how critical a failure in reading is
	if err != nil {
		// Option 1: Log error and return some default deck
		// Option 2: If this is mission critical, log error and terminate the program
		fmt.Println("Error occured while attempting to generate a deck from the specified filepath. Please check your input.")
		fmt.Println(err)
		os.Exit(1)
	}

	// Retrieve a string slice by first converting the byteslice to a string, followed by a split on the "," to get the string slice
	s := strings.Split(string(bs), ",")

	// Then we create the deck variable and return
	fmt.Println("Successfully read file and generated deck.")
	return deck(s)
}
