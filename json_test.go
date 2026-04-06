package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	recorder := httptest.NewRecorder()

	respondWithJSON(recorder, http.StatusCreated, map[string]string{"status": "ok"})

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, recorder.Code)
	}

	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected content type application/json, got %q", got)
	}

	var payload map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid json body, got %v", err)
	}

	if payload["status"] != "ok" {
		t.Fatalf("expected status payload %q, got %q", "ok", payload["status"])
	}
}

func TestRespondWithError(t *testing.T) {
	recorder := httptest.NewRecorder()

	respondWithError(recorder, http.StatusBadRequest, "bad input", errors.New("boom"))

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	var payload map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid json body, got %v", err)
	}

	if payload["error"] != "bad input" {
		t.Fatalf("expected error payload %q, got %q", "bad input", payload["error"])
	}
}
