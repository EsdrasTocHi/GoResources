package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Run with: go run main.go https://golang.org http://gopl.io https://godoc.org

// Golang is a programming language that have native concurrency, to use the concurrency we have to use the go routines.

// What is a go routine?
// A go routine is similar to a threads, but with concurrency.
// To use a go routine, use the word "go" before of the calling to a function (Example:  go f(s))

// What is a channel?
// Is the way to communicate the go routines. The channels have a specified data type.
// To send information with the channel use:       channelName <- data
// To access to the information in the channel use:      data <- channelName

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go func(url string, ch chan<- string) {
			start := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				ch <- fmt.Sprint(err)
				return
			}

			nbytes, err := io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			if err != nil {
				ch <- fmt.Sprintf("while reading %s: %v", url, err)
				return
			}

			secs := time.Since(start).Seconds()
			ch <- fmt.Sprintf("%.2fs  %7d %s", secs, nbytes, url)
		}(url, ch)
	}

	for range os.Args[1:] {
		// Everytime that you try to access to the information in the channel, the program will stop
		// waiting to receive data.
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
