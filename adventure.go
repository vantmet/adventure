// Start Here!
package main

import (
	"fmt"
)

type Room struct {
	Title       string
	Description string
}

func Welcome() string {
	return "Welcome to Gopher Castle! While on a walking holiday, you notice an interesting-looking castle in the distance, and decide to investigate."
}

func EnterRoomDescription(room Room) string {
	result := room.Title + "\n" + "\n" + room.Description

	return result
}

func main() {
	// Print welcome message to console
	fmt.Println(Welcome())
}
