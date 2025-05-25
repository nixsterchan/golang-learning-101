package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// The following uses http to call a website, and copy its body contents to say a log writer for example.
// Note that in this case, we cannot simply create the log writer variable and immediately do a copy of the resp.Body
// to the logwriter.

// The logWriter must implement a Write method in order for io.Copy to work on it. Try commenting out the Write method below and see what happens

func main() {
	website := "http://google.com"
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println("Ran into an error when attempting to ping the website", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// Try commenting this function out. It is required in order for io.Copy to work with the logWriter type we created.
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote num bytes: ", len(bs))

	return len(bs), nil
}
