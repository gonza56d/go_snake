package gameassets

import "fmt"

type Snake []Location // Index 0 is the snake's head. Last index is the tail.

type snakeMovement uint8

const (
	Up snakeMovement = iota
	Down
	Left
	Right
)

// getSnakeFacingDirection determines the direction the snake's head is facing
// based on the position of its head and neck.
//
// Parameters:
//   - snake: A pointer to a Snake instance.
//
// Returns:
//   - The direction (Up, Down, Left, Right) that the snake's head is facing.
//
// Panics:
//   - If the snake is in an illegal state (e.g., head and neck positions do not align).
func getSnakeFacingDirection(snake *Snake) snakeMovement {
	var headLocation Location = (*snake)[0]
	var neckLocation Location = (*snake)[1] 
	var facingDirection snakeMovement

	if headLocation.XAt == neckLocation.XAt && headLocation.YAt > neckLocation.YAt {
		facingDirection = Right
	} else if headLocation.XAt == neckLocation.XAt && headLocation.YAt < neckLocation.YAt {
		facingDirection = Left	
	} else if headLocation.YAt == neckLocation.YAt && headLocation.XAt > neckLocation.XAt {
		facingDirection = Up
	} else if headLocation.YAt == neckLocation.YAt && headLocation.XAt < neckLocation.XAt {
		facingDirection = Down
	} else {
		panic("Illegal snake state.")
	}
	return facingDirection
}

// isIllegalMovement checks if a proposed movement direction is illegal
// (i.e., the snake is trying to move in the opposite direction to its current facing).
//
// Parameters:
//   - facingDirection: The current direction the snake is facing.
//   - movingDirection: The proposed new direction.
//
// Returns:
//   - True if the movement is illegal; false otherwise.
func isIllegalMovement(facingDirection snakeMovement, movingDirection snakeMovement) bool {
	return (facingDirection == Up && movingDirection == Down) ||
		(facingDirection == Down && movingDirection == Up) ||
		(facingDirection == Left && movingDirection == Right) ||
		(facingDirection == Right && movingDirection == Left)
}

// MoveSnake moves the snake one square in the specified direction, updating
// the positions of the head and body.
//
// Parameters:
//   - snake: A pointer to a Snake instance to update.
//   - movingDirection: The direction to move the snake (Up, Down, Left, or Right).
//
// Behavior:
//   - If the proposed direction is illegal, the snake continues in its current direction.
//   - The head moves one square in the given direction.
//   - Each segment of the body moves to the position of the segment in front of it.
//
// Panics:
//   - If an invalid movement direction is provided.
func MoveSnake(snake *Snake, movingDirection snakeMovement) {
	facingDirection := getSnakeFacingDirection(snake)
	if isIllegalMovement(facingDirection, movingDirection) {
		movingDirection = facingDirection
	}

	for i := 1; i < len(*snake); i++ {
		(*snake)[i] = (*snake)[i-1]
	}

	switch movingDirection {
		case Up: (*snake)[0].YAt++
		case Down: (*snake)[0].YAt--
		case Left: (*snake)[0].XAt--
		case Right: (*snake)[0].XAt++
		default: panic(fmt.Errorf("Illegal snake movement direction: %d.", movingDirection))
	}
}

