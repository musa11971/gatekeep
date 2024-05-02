package main

import (
	"fmt"
	"io"
	"log"
	"musa11971/gatekeep/policy"
	"net/http"

	"github.com/fatih/color"
)

const port = 9090

func main() {
	// Start the web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedPath := r.URL.Path

		policy, error := policy.FindWithRoutingPath(requestedPath)
		_ = policy

		// No policy found
		if error != nil {
			fmt.Fprintf(
				w,
				fmt.Sprintf("No policy found for path: %s", requestedPath),
			)
			return
		}

		// Create the request to the endpoint
		client := &http.Client{}

		method := r.Method
		url := policy.FullEndpointURL()

		req, error := http.NewRequest(method, url, nil)
		response, error := client.Do(req)
		body, error := io.ReadAll(response.Body)

		fmt.Fprintf(
			w,
			fmt.Sprintf("Data: %s", body),
		)
	})

	color.HiMagenta("Starting to listen on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
