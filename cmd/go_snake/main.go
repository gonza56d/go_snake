package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	// Initialize the keyboard listener
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Press any key to see its value. Press ESC to quit.")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			// Handle the error but don't exit the program
			fmt.Println("Error reading key:", err)
			continue // Skip this iteration and wait for the next key press
		}

		// Display the key pressed
		fmt.Printf("You pressed: char=%q, key=%v\n", char, key)

		// Exit the program when ESC is pressed
		if key == keyboard.KeyEsc {
			fmt.Println("Exiting...")
			break
		}
	}
}
