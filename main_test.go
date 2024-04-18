package main

import (
	"fmt"
	"testing"
)

func Test_welcome(t *testing.T) {
	text := "Jack"
	result := fmt.Sprintf("Welcome, %s", text)

	if got := welcome(text); result != got {
		t.Errorf("Expected %q got %q", result, got)
	}
}
