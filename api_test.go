package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Mock handler for testing purposes
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
}

func TestHelloEndpoint(t *testing.T) {
    // Create a new HTTP request
    req, err := http.NewRequest("GET", "/api/hello", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Record the HTTP response
    rr := httptest.NewRecorder()

    // Create a new HTTP handler
    handler := http.HandlerFunc(helloHandler)

    // Serve the HTTP request
    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body
    var actual map[string]string
    if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
        t.Fatalf("could not decode response: %v", err)
    }

    expected := map[string]string{"message": "Hello, World!"}
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("handler returned unexpected body: got %v want %v",
            actual, expected)
    }
}
