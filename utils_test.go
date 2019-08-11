package main

import "testing"

func TestGenerateIDLength(t *testing.T) {
	want := 6
	got := len(generateID(6))
	if got != want {
		t.Errorf("generateID(%d) should produce an ID of length %d, but was %d", want, want, got)
	}
}
