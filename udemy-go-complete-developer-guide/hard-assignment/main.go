package main

import (
	"fmt"
	"os"
)

// This was just an assignment given by the udemy tutor. Task was to have a file passed in via command line, where said file would be
// read and its contents read out to the terminal.

func main() {
	// Get the file via usage of the argument from the command line
	if len(os.Args) < 2 {
		fmt.Println("Please provide a text file name as input in the command line.")
		os.Exit(1)
	}

	// Retrieve the argument as a variable
	filepath := os.Args[1]

	// Read out the data from the provided filepath
	data, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println("An error occured while trying to read the file:", err)
		os.Exit(1)
	}
	fmt.Println("File successfully read! It's contents are:")
	fmt.Println(string(data))

	// // Course version of the answer
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// io.Copy(os.Stdout, f)

}
