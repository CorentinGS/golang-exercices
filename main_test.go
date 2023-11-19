package main

import (
	"io"
	"net/http"
	"sync"
	"testing"
)

// Replace this with the actual endpoint URL
var endpoint = "http://localhost:8080/user"

// Number of concurrent requests to send
var concurrentRequests = 500

// Test function to benchmark the endpoint
func TestEndpoint(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeRequest(t)
		}()
	}

	wg.Wait()
}

func makeRequest(t *testing.T) {
	// Make the request
	resp, err := http.Get(endpoint)
	if err != nil {
		t.Fatalf("could not send request: %v", err)
	}

	// Read the response body
	defer resp.Body.Close()
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		t.Fatalf("could not read response: %v", err)
	}
}
