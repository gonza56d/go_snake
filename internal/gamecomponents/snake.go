package gamecomponents

import "fmt"

type Snake []Location // Index 0 is the snake's head. Last index is the tail.

type SnakeMovement uint8

const (
	Up SnakeMovement = iota
	Down
	Left
	Right
)

func GenerateNewSnake(gameMap *Map) *Snake {
	halfX, halfY := gameMap.XSize / 2, gameMap.YSize / 2
	return &Snake{
		Location{XAt: int16(halfX), YAt: int16(halfY)},
		Location{XAt: int16(halfX), YAt: int16(halfY) - 1},
		Location{XAt: int16(halfX), YAt: int16(halfY) - 2},
		Location{XAt: int16(halfX), YAt: int16(halfY) - 3},
	}
}

// getFacingDirection determines the direction of movement (facing) 
// between two coordinates in a 2D space. It is used to calculate the 
// facing direction of the snake's head or tail, which is essential for 
// growing or moving the snake in the game.
//
// Parameters:
// - pointingCoord (Location): The coordinate whose direction is to be determined.
// - previousCoord (Location): The coordinate being compared to.
//
// Returns:
// - SnakeMovement: The direction (`Right`, `Left`, `Up`, or `Down`) 
//   that `pointingCoord` is facing relative to `previousCoord`.
//
// Panics:
// - If the two coordinates are not aligned horizontally or vertically 
//   (invalid snake state), the function will panic.
//
// Example:
//   pointingCoord := Location{XAt: 3, YAt: 2}
//   previousCoord := Location{XAt: 2, YAt: 2}
//   direction := getFacingDirection(pointingCoord, previousCoord)
//   fmt.Println(direction) // Output: Right
func getFacingDirection(pointingCoord Location, previousCoord Location) SnakeMovement {
	if pointingCoord.XAt > previousCoord.XAt && pointingCoord.YAt == previousCoord.YAt {
		return Right
	} else if pointingCoord.XAt < previousCoord.XAt && pointingCoord.YAt == previousCoord.YAt {
		return Left	
	} else if pointingCoord.YAt > previousCoord.YAt && pointingCoord.XAt == previousCoord.XAt {
		return Up
	} else if pointingCoord.YAt < previousCoord.YAt && pointingCoord.XAt == previousCoord.XAt {
		return Down
	} else {
		panic("Illegal snake state.")
	}
}

// Grow the snake by reference.
// Appends one segment after snake's tail.
func Grow(snake *Snake) {
	snakeLength := len(*snake)
	tail := (*snake)[snakeLength - 1] 
	preTail := (*snake)[snakeLength - 2]

	tailFacingDirection := getFacingDirection(tail, preTail)
	newSegment := Location{XAt: tail.XAt, YAt: tail.YAt}

	switch tailFacingDirection {
	case Up:
		newSegment.YAt++
	case Down:
		newSegment.YAt--
	case Left:
		newSegment.XAt--
	case Right:
		newSegment.XAt++
	}
	*snake = append(*snake, newSegment)
}

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
	return getFacingDirection(headLocation, neckLocation)
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

	for i := len(*snake) - 1; i > 0; i-- { 
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
// - bool: True if the snake's head is outside the map boundaries or overlapping their own body, false otherwise.
func IsSnakeCrashed(snake *Snake, gameMap *Map) bool {
	crashedAgainstWall := (*snake)[0].YAt < 0 || (*snake)[0].XAt < 0 ||
		(*snake)[0].YAt > int16(gameMap.YSize) || (*snake)[0].XAt > int16(gameMap.XSize)

	if crashedAgainstWall {
		return true
	}

	snakeLocations := make(map[Location]struct{})
	for _, segment := range (*snake)[1:] {
		snakeLocations[segment] = struct{}{}
	}
	snakeHead := (*snake)[0]
	_, isSnakeHeadOverlappingBody := snakeLocations[snakeHead]
	return isSnakeHeadOverlappingBody  
}
