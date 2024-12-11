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
		'w': gamecomponents.Left,
		'W': gamecomponents.Left,
		'a': gamecomponents.Down,
		'A': gamecomponents.Down,
		's': gamecomponents.Right,
		'S': gamecomponents.Right,
		'd': gamecomponents.Up,
		'D': gamecomponents.Up,
	}
	// initial draw for screen (walls)
	maxX := int(match.GameMap.XSize) + 2
	maxY := int(match.GameMap.YSize) + 2
	screen := make([][]rune, maxX+1)
	for i := range screen {
		screen[i] = make([]rune, maxY+1)
	}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if (x==0 && y==0) || (x==0 && y==maxY) || (x==maxX && y==0) || (x==maxX && y==maxY) {
				screen[x][y] = 'X' // corners
			} else if x == 0 || x == maxX { 
				screen[x][y] = '-' // bottom/top walls
			} else if y == 0 || y == maxY {
				screen[x][y] = '|' // side walls
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
		ticker := time.NewTicker(60 * time.Millisecond)
		defer ticker.Stop()

		movement := nextMovement
		gameLoop:
		for {
			select {
			case <-ticker.C:
				snakeTail := (*match.Snake)[len(*match.Snake)-1]
				screen[snakeTail.XAt+1][snakeTail.YAt+1] = ' '
				crashed := match.MakeMove(movement)
				if crashed {
					break gameLoop
				}
				fmt.Println("\033[H\033[2J")
				snake := *match.Snake
				screen[snake[0].XAt+1][snake[0].YAt+1] = 'O'
				for _, segment := range snake[1:] {
					screen[segment.XAt+1][segment.YAt+1] = 'X'
				}
				screen[match.GameMap.FoodAt.XAt+1][match.GameMap.FoodAt.YAt+1] = '*'
				for x := range screen {
					fmt.Println(string(screen[x]))
				}
			case newMovement := <-movementChannel:
				movement = newMovement
			}
		}
		fmt.Printf("X.X Snake crashed X.X\nGame over!\nYour score: %d\n", match.Score)
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
