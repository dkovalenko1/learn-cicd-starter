package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerReadiness(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/healthz", nil)
	recorder := httptest.NewRecorder()

	handlerReadiness(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var payload map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid json body, got %v", err)
	}

	if payload["status"] != "ok" {
		t.Fatalf("expected health status %q, got %q", "ok", payload["status"])
	}
}
