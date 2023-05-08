package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello1()
	want := "same digits - please try again"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
