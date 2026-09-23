package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatal(err)
	}

	if key != "test-key" {
		t.Fatalf("Got: %s. Expected 'test-key'", key)
	}
}

func TestGetApiKeyMissingHeader(t *testing.T) {
	headers := http.Header{}

	key, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("Got: %v. Expected ErrNoAuthHeaderIncluded", err)
	}

	if key != "" {
		t.Fatalf("Got: %s. Expected empty key", key)
	}
}

func TestGetApiKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer test-key")

	key, err := GetAPIKey(headers)

	if err == nil || err.Error() != "malformed authorization header" {
		t.Fatalf("Got: %v. Expected malformed authorization header", err)
	}

	if key != "" {
		t.Fatalf("Got: %s. Expected empty key", key)
	}
}
