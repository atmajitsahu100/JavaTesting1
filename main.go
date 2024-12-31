package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIWithHardcodedKey(t *testing.T) {
	// Vulnerability: Hardcoded sensitive information
	apiKey := "1234567890abcdef" // Should not be hardcoded in the code

	url := "https://example.com/api?api_key=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Error while making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Checking if the status code is 200 (OK)
	assert.Equal(t, 200, resp.StatusCode, "Expected status code 200")

	// Reading the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// Checking if the response body contains expected data (optional)
	assert.Contains(t, string(body), "success", "Expected response body to contain 'success'")
	fmt.Println("Response body:", string(body))
}
