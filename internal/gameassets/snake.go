package gameassets

import "fmt"

type Snake []Location // Index 0 is the snake's head. Last index is the tail.

type SnakeMovement uint8

const (
	Up SnakeMovement = iota
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
func getSnakeFacingDirection(snake *Snake) SnakeMovement {
	var headLocation Location = (*snake)[0]
	var neckLocation Location = (*snake)[1] 
	var facingDirection SnakeMovement

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
func isIllegalMovement(facingDirection SnakeMovement, movingDirection SnakeMovement) bool {
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
func MoveSnake(snake *Snake, movingDirection SnakeMovement) {
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

// IsSnakeCrashed checks if the snake has crashed into the map boundaries.
// 
// Parameters:
// - snake (*Snake): A pointer to the Snake, where the head is at index 0.
// - gameMap (*Map): A pointer to the Map representing the game area dimensions.
//
// Returns:
// - bool: True if the snake's head is outside the map boundaries, false otherwise.
//
// Behavior:
// The function compares the coordinates of the snake's head (XAt, YAt) to the map's size
// (XSize, YSize). If the head's position is less than 0 or greater than the map's size,
// the function returns true, indicating a crash. Otherwise, it returns false.
func IsSnakeCrashed(snake *Snake, gameMap *Map) bool {
	return (*snake)[0].YAt < 0 || (*snake)[0].XAt < 0 ||
		(*snake)[0].YAt > int16(gameMap.YSize) || (*snake)[0].XAt > int16(gameMap.XSize)
}
