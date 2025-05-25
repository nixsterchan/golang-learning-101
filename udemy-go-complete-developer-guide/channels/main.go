package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Make a channel which allows communications between child and main routines
	c := make(chan string)
	for _, l := range links {
		// Using go in front of function calls will start utilizing the go scheduler to run things concurrently [not true parallelism]
		go checkLink(l, c)
	}

	// Must catch messages from the channel len(links) number of times
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// For repeating routines, will set a continous loop. Since checkLink is essentially a blocking call, the routine is only
	// called when the channel retrieves something
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Can also set up a loop using the channel. It waits for a value[link] to be put into the channel before
	// the go routine is run in the loop
	for l := range c {
		go checkLink(l, c)
	}

	// Now to try the lambda method where you will wrap the functions to call in a function literal (think of lambda in python)
	// when calling a go routine.
	// NOTE: when using this, must ensure to pass in the l from the loop as a variable into this function literal
	// failing to do so, the main routine and child routine will be pointing to the same address in memory, to which if the main
	// routine were to change the current value of l, the child routine might get funky and start utilizing an unexpected
	// new value instead of what is previously would have expected.
	//
	// Hence in this case, using GO's pass by style, will pass in this variable l into the func so that it would have
	// a copy in memory to work with instead of the same variable that the main routine is looking at

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

// This example is for usage with the repeating routines paradigm where we pass the link back into the channel
// for a new checkLink child routine to be started again via the forEVER loop
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")

		// To send messages into the channel
		c <- link
		return
	}
	fmt.Println(link, " is up!!")
	c <- link
}

// Normal example of passing some value into the channel to mark the end of the go routine
// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, " might be down!")

// 		// To send messages into the channel
// 		c <- "Might be down I think"
// 		return
// 	}
// 	c <- "Yup its up"
// }
