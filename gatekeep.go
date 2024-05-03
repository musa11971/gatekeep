package main

import (
	"fmt"
	"log"
	"musa11971/gatekeep/routing"
	"net/http"
	"github.com/fatih/color"
)

const port = 9090

func main() {
	routing.Initialize()

	// Start the web server
	fmt.Println()
	color.HiMagenta("🔒 GateKeep is ready.")
	color.HiMagenta("Starting to listen on %d", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}