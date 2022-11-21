package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for key, val := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", key, val)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

		// You can write a declaration before the condition of the if statement
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		for key, val := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", key, val)
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
