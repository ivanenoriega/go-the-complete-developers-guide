package main

import "testing"

func TestNewDeck(t *testing.T) {
	e := 16
	d := newDeck()

	if len(d) != e {
		t.Errorf("Expected deck of length %d, but got %v", e, len(d))
	}
}
