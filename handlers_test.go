package main

import "testing"

func TestCreateURL(t *testing.T) {
	id := "wcz20k"
	action := "v"
	want := "/v/wcz20k"
	got := createURL(id, action)
	if got != want {
		t.Errorf("createURL(%s, %s) should produce %s, but was %s", id, action, want, got)
	}
}
