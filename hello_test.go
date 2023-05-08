package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello1()
	want := "Hello, world :D"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
