package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " %s from Malay", quote.Hello())
	})

	http.ListenAndServe(":8081", nil)
}
