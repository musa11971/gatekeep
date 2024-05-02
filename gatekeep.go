package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

const port = 9090

func main() {
	// Start the web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("You called path: %s", r.URL.Path)

		fmt.Fprintf(w, response)
	})

	color.HiMagenta("Starting to listen on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
