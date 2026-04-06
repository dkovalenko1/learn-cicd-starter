package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey secret-key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if apiKey != "secret-key" {
		t.Fatalf("expected api key %q, got %q", "secret-key", apiKey)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer secret-key")

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected malformed header error, got nil")
	}

	if err.Error() != "malformed authorization header" {
		t.Fatalf("expected malformed authorization header error, got %v", err)
	}
}
