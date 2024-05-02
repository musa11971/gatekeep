package main

import (
	"fmt"
	"io"
	"log"
	"musa11971/gatekeep/policy"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

const port = 9090

func main() {
	// Register the policy routes
	r := mux.NewRouter()

	go registerPolicyRoutes(r)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "No policy found.")
	})

	http.Handle("/", r)

	// Start the web server
	fmt.Println()
	color.HiMagenta("🔒 GateKeep is ready.")
	color.HiMagenta("Starting to listen on %d", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func RouteToEndpointHandler(policy policy.Policy, w http.ResponseWriter, r *http.Request) {
	// Create the request to the endpoint
	client := &http.Client{}

	method := r.Method
	url := policy.FullEndpointURL()

	req, _ := http.NewRequest(method, url, nil)
	response, _ := client.Do(req)

	// Copy headers
	for key, values := range response.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Copy status code
	w.WriteHeader(response.StatusCode)

	// Copy body
	_, error := io.Copy(w, response.Body)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
}

func registerPolicyRoutes(router *mux.Router) {
	// To do: it should register each policy as a go-routine.
	for _, p := range policy.Policies {
		fmt.Println("Registering RoutingPath", p.RoutingPath)

		router.PathPrefix(p.RoutingPath).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			RouteToEndpointHandler(p, w, r)
		})
	}
}
