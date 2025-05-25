package main

import "fmt"

func main() {
	// You can use this method of specifying the map
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#g20sfs",
	// 	"white": "#ffffff",
	// }

	// You could also utilize the make function for doing so
	fmt.Println("Building map.....")
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#g20sfs"
	colors["white"] = "#ffffff"

	// fmt.Println(colors)
	printMap(colors)

	// you can delete directly from a map too via
	fmt.Println("Removing red from the colors map....")
	delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v.\n", color, hex)
	}
}
