package controllers

import (
	"io"
	"musa11971/gatekeep/policy"
	"net/http"
	"time"
)

func ForwardToPolicyHandler(policy policy.Policy, w http.ResponseWriter, r *http.Request) {
	// Create the request to the endpoint
	client := &http.Client{
		Timeout: time.Second * time.Duration(policy.EndpointReadTimeout),
	}

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