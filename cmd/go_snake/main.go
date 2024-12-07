package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/gonza56d/go_snake/internal/gamecomponents"
	"github.com/gonza56d/go_snake/internal/gamerunner"
)

func main() {
	match := gamerunner.NewMatch(3)
	nextMovement := gamecomponents.Up
	movementsMap := map[rune]gamecomponents.SnakeMovement{
		'w': gamecomponents.Up,
		'W': gamecomponents.Up,
		'a': gamecomponents.Left,
		'A': gamecomponents.Left,
		's': gamecomponents.Down,
		'S': gamecomponents.Down,
		'd': gamecomponents.Right,
		'D': gamecomponents.Right,
	}
	// Initialize the keyboard listener
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	movementChannel := make(chan gamecomponents.SnakeMovement, 1)

	go func() {
		ticket := time.NewTicker(100 * time.Millisecond)
		defer ticket.Stop()

		movement := nextMovement
		for {
			select {
			case <-ticker.C:
				match.MakeMove(movement)
				var screenBuffer string
				for x := range match.GameMap.XSize {
					for y := range match.GameMap.YSize {
						if x == 0 || x == len(match.GameMap.XSize) || y == 0 || y == len(match.GameMap.YSize) {
							
						}
					}
				}
			case newMovement := <-movementChannel:
				movement = newMovement
			}
		}
	}

	for {
		char, key, err := keyboard.GetSingleKey()
		// Exit the program when ESC is pressed
		if err {
			fmt.Printf("Error happened while pressing key %v", key)
		}
		if char == '\x00' {
			fmt.Println("Exiting...")
			break
		}
		if movement, exists := movementsMap[char]; exists {
			select {
			case movementChannel <- movement:
			default:
			}
		}
	}
}
