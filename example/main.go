package main

import (
	"fmt"
	"github.com/arifamir/goscreen"
	"time"
)

func main() {
	// Clear all the characters on the screen
	goscreen.Clear()

	for {
		// Moves the cursor to the top-left position of the screen
		goscreen.MoveTopLeft()

		// Animate the time always in the same position
		fmt.Println(time.Now())

		time.Sleep(time.Second)
	}
}