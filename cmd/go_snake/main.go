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
	// initial draw for screen (walls)
	maxX := int(match.GameMap.XSize) + 2
	maxY := int(match.GameMap.YSize) + 2
	screen := make([][]byte, maxX+1)
	for i := range screen {
		screen[i] = make([]byte, maxY+1)
	}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if (x==0 && y==0) || (x==0 && y==maxY) || (x==maxX && y==0) || (x==maxX && y==maxY) {
				screen[x][y] = 'X' // corners
			} else if x == 0 || x == maxX { 
				screen[x][y] = '-' // bottom/top walls
			} else if y == 0 || y == maxY {
				screen[x][y] = 'l' // side walls
			} else {
				screen[x][y] = ' ' // blank spaces for snake & food
			}
		}
	}
	for x := range screen {
		fmt.Println(string(screen[x]))
	}

	// Initialize the keyboard listener
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	movementChannel := make(chan gamecomponents.SnakeMovement, 1)

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		movement := nextMovement
		for {
			select {
			case <-ticker.C:
				match.MakeMove(movement)
			case newMovement := <-movementChannel:
				movement = newMovement
			}
		}
	}()

	for {
		char, key, err := keyboard.GetSingleKey()
		// Exit the program when ESC is pressed
		if err != nil {
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
