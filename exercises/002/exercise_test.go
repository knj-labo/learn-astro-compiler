package main

import (
	"testing"
)

func TestEx002(t *testing.T) {
	// check for error
	want := int(0)
	got, err := Exercise002(-10)
	if err == nil {
		t.Errorf("Exercise002() = %v, want error", got)
	}

	want = 1
	got, err = Exercise002(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("Expected '%v' but got '%v'", want, got)
	}

	want = 40320
	got, err = Exercise002(8)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("Exercise002() = %v, want %v", got, want)
	}
}
