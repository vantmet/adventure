package main

import "testing"

func TestWelcome(t *testing.T) {
	got := Welcome()
	want := "Welcome to Gopher Castle! While on a walking holiday, you notice an interesting-looking castle in the distance, and decide to investigate."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFirstRoomDescription(t *testing.T) {
	room := Room{
		Title:       "Outside the Castle",
		Description: "\tYou're standing outside a tumbledown castle, in front of a massive and\n\theavily-scarred wooden front door to the north. It is standing\n\tinvitingly half-open. Do you dare to enter?",
	}

	got := EnterRoomDescription(room)
	want := `Outside the Castle

	You're standing outside a tumbledown castle, in front of a massive and
	heavily-scarred wooden front door to the north. It is standing
	invitingly half-open. Do you dare to enter?`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
