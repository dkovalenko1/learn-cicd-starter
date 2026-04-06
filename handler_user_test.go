package main

import (
	"regexp"
	"testing"
)

func TestGenerateRandomSHA256Hash(t *testing.T) {
	first, err := generateRandomSHA256Hash()
	if err != nil {
		t.Fatalf("expected no error generating first hash, got %v", err)
	}

	second, err := generateRandomSHA256Hash()
	if err != nil {
		t.Fatalf("expected no error generating second hash, got %v", err)
	}

	if len(first) != 64 {
		t.Fatalf("expected first hash length 64, got %d", len(first))
	}

	if len(second) != 64 {
		t.Fatalf("expected second hash length 64, got %d", len(second))
	}

	hexPattern := regexp.MustCompile("^[0-9a-f]{64}$")
	if !hexPattern.MatchString(first) {
		t.Fatalf("expected first hash to be lowercase hex, got %q", first)
	}

	if !hexPattern.MatchString(second) {
		t.Fatalf("expected second hash to be lowercase hex, got %q", second)
	}

	if first == second {
		t.Fatal("expected generated hashes to differ")
	}
}
