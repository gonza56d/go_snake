package gamerunner

import "github.com/gonza56d/go_snake/internal/gamecomponents"

// Match holds data from current game going on.
// Interacts with game components as snake moves (MoveSnake function is called).
type Match struct {
	GameMap gamecomponents.Map
	Score uint64
	Snake gamecomponents.Snake
}

// MakeMove for snake of the current match to the provided direction.
// Grows snake and adds score if snake ate food after moving (respawns food if so).
// Returns if snake crashed after moving (game over).
func (m *Match) MakeMove(direction gamecomponents.SnakeMovement) bool {
	gamecomponents.MoveSnake(&m.Snake, direction)
	snakeHeadAtFood := m.Snake[0].XAt == m.GameMap.FoodAt.XAt && m.Snake[0].YAt == m.GameMap.FoodAt.YAt
	if snakeHeadAtFood { 
		gamecomponents.Grow(&m.Snake)
		m.Score++
		m.GameMap.FoodAt = *gamecomponents.SpawnFood(&m.Snake, &m.GameMap)
	}
	return gamecomponents.IsSnakeCrashed(&m.Snake, &m.GameMap)
}

