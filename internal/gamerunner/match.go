package gamerunner

import (
	"github.com/gonza56d/go_snake/internal/gamecomponents"
)

type Match struct {
	gameMap gamecomponents.Map
	score uint64
	snake gamecomponents.Snake
}

// MoveSnake of the current match to the provided direction.
// Grows snake and adds score if snake ate food after moving (respawns food if so).
// Returns if snake crashed after moving (game over).
func (m Match) MoveSnake(direction gamecomponents.SnakeMovement) bool {
	gamecomponents.MoveSnake(&m.snake, direction)
	snakeHeadAtFood := m.snake[0].XAt == m.gameMap.FoodAt.XAt && m.snake[0].YAt == m.gameMap.FoodAt.YAt
	if snakeHeadAtFood { 
		gamecomponents.Grow(&m.snake)
		m.score++
		m.gameMap.FoodAt = *gamecomponents.SpawnFood(&m.snake, &m.gameMap)
	}
	return gamecomponents.IsSnakeCrashed(&m.snake, &m.gameMap)
}

