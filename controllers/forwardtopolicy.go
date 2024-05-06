package controllers

import (
	"io"
	"musa11971/gatekeep/policy"
	"net/http"
	"time"
)

func ForwardToPolicyHandler(policy policy.Policy, w http.ResponseWriter, r *http.Request) {
	response, _ := performRequestToEndpoint(policy, r)

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

func performRequestToEndpoint(policy policy.Policy, incomingRequest *http.Request) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(policy.EndpointReadTimeout),
	}

	req, _ := http.NewRequest(
		incomingRequest.Method,
		policy.FullEndpointURL(),
		nil,
	)
	req.Header = incomingRequest.Header

	return client.Do(req)
}