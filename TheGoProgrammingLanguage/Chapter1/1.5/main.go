package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Go have a library speciffically to fetching urls, this is net.
// You can do GET and POST requests.

// This code use http.Get to do a get request with that you put in command-line
// every http request give us a response. We can read this response as string with ioutil
// and the body of the response

// try;
//   go run main.go http://gopl.io
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
