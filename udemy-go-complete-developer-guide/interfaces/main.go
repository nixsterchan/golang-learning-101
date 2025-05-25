package main

import "fmt"

// Interfaces can be used to sort of group up types that have the same kind of function calls
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type japaneseBot struct{}

func main() {
	eb := englishBot{}
	jb := japaneseBot{}

	// This works
	// fmt.Println("Greetings via normal call of getGreeting function.")
	// eb.getGreeting()
	// jb.getGreeting()

	// This works too via the interface!
	fmt.Println("Greetings via use of the bot interface's printGreeting function.")
	printGreeting(eb)
	printGreeting(jb)
}

// We see here that both the english and japanese bots have the same underlying getGreeting functions
func (englishBot) getGreeting() string {
	return "Hello there, I am english."
}

func (japaneseBot) getGreeting() string {
	return "Konbanwa, I am japanese."
}

// Using that fact, we create the bot interface by grouping the two bots based on the fact that they both have the same function.
// Think of it as a membership kinda system where having a particular set of functions allows you to `join the club` where said club is the interface.
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
