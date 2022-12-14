package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.path = %q \n", r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
