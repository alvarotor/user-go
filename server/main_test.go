package main

// create a simple test
import (
	"testing"
)

func TestAdd(t *testing.T) {
	if got, want := Add(2, 2), 4; got != want {
		t.Errorf("Add(2,2) = %d, want %d", got, want)
	}
}

func Add(a, b int) int {
	return a + b
}
