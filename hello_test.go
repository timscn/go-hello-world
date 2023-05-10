package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got, err := Hello1()
	if err != nil {
		want := "same digits - please try again"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
