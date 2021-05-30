package main

import "testing"

func TestWelcome(t *testing.T) {
	got := Welcome()
	want := "Welcome to Gopher Castle! While on a walking holiday, you notice an interesting-looking castle in the distance, and decide to investigate."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
