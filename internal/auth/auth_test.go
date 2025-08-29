package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testHeader := http.Header{}

	testHeader.Set("Authorization", "ApiKey 12345")

	_, err := GetAPIKey(testHeader)
	if err != nil {
		t.Errorf("Failed to extract API key with correct format: %s", err)
	}

	testHeader.Set("Authorization", "12345")
	_, err = GetAPIKey(testHeader)
	if err == nil {
		t.Errorf("Error checking did not catch malformed authorization")
	}

	testHeader.Set("Authorization", "")
	_, err = GetAPIKey(testHeader)
	if err == nil {
		t.Errorf("Error checking did not stop empty authorization")
	}
}
