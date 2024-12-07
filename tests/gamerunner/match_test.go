package gamerunner

import (
	"testing"

	"github.com/gonza56d/go_snake/internal/gamecomponents"
	"github.com/gonza56d/go_snake/internal/gamerunner"
)

func TestMakeMoveAndNothingHappens(t *testing.T) {
	tests := []struct{
		snakeMovement gamecomponents.SnakeMovement
		match *gamerunner.Match
	}{
		{
			gamecomponents.Up,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 1, YAt: 1}},
				Score: uint64(3),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 5},
					gamecomponents.Location{XAt: 5, YAt: 4},
					gamecomponents.Location{XAt: 5, YAt: 3},
					gamecomponents.Location{XAt: 5, YAt: 2},
				},
			},
		},
		{
			gamecomponents.Right,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 1, YAt: 1}},
				Score: uint64(3),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 5},
					gamecomponents.Location{XAt: 4, YAt: 5},
					gamecomponents.Location{XAt: 3, YAt: 5},
					gamecomponents.Location{XAt: 2, YAt: 5},
				},
			},
		},
		{
			gamecomponents.Left,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 1, YAt: 1}},
				Score: uint64(3),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 5},
					gamecomponents.Location{XAt: 6, YAt: 5},
					gamecomponents.Location{XAt: 7, YAt: 5},
					gamecomponents.Location{XAt: 8, YAt: 5},
				},
			},
		},
		{
			gamecomponents.Down,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 1, YAt: 1}},
				Score: uint64(3),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 5},
					gamecomponents.Location{XAt: 5, YAt: 6},
					gamecomponents.Location{XAt: 5, YAt: 7},
					gamecomponents.Location{XAt: 5, YAt: 8},
				},
			},
		},
	}

	for _, test := range tests {
		scoreBeforeMoving := test.match.Score
		var gameOver bool = test.match.MakeMove(test.snakeMovement)
		if gameOver {
			t.Error("Snake shouldn't have crashed after MakeMove.")
		}
		if scoreBeforeMoving != test.match.Score {
			t.Errorf("scoreBeforeMoving (%d) should be equal to test.match.Score (%d) after MakeMove(...) call.",
				scoreBeforeMoving, test.match.Score)
		}
	}
}

func TestMakeMoveAndEatFood(t *testing.T) {
	tests := []struct{
		snakeMovement gamecomponents.SnakeMovement
		match *gamerunner.Match
	}{
		{
			gamecomponents.Up,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 5, YAt: 6}},
				Score: uint64(6),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 5},
					gamecomponents.Location{XAt: 5, YAt: 4},
					gamecomponents.Location{XAt: 5, YAt: 3},
					gamecomponents.Location{XAt: 5, YAt: 2},
				},
			},
		},
		{
			gamecomponents.Right,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 6, YAt: 5}},
				Score: uint64(6),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 5},
					gamecomponents.Location{XAt: 4, YAt: 5},
					gamecomponents.Location{XAt: 3, YAt: 5},
					gamecomponents.Location{XAt: 2, YAt: 5},
				},
			},
		},
		{
			gamecomponents.Left,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 4, YAt: 5}},
				Score: uint64(6),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 5},
					gamecomponents.Location{XAt: 6, YAt: 5},
					gamecomponents.Location{XAt: 7, YAt: 5},
					gamecomponents.Location{XAt: 8, YAt: 5},
				},
			},
		},
		{
			gamecomponents.Down,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 5, YAt: 4}},
				Score: uint64(6),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 5},
					gamecomponents.Location{XAt: 5, YAt: 6},
					gamecomponents.Location{XAt: 5, YAt: 7},
					gamecomponents.Location{XAt: 5, YAt: 8},
				},
			},
		},
	}

	for _, test := range tests {
		scoreBeforeMoving := test.match.Score
		var gameOver bool = test.match.MakeMove(test.snakeMovement)
		if gameOver {
			t.Error("Snake shouldn't have crashed after MakeMove.")
		}
		if scoreBeforeMoving + 1 != test.match.Score {
			t.Errorf("test.match.Score (%d) should be +1 greater than scoreBeforeMoving (%d) " + 
				"after MakeMove(...) call.", test.match.Score, scoreBeforeMoving)
		}
	}
}

func TestMakeMoveAndGameOver(t *testing.T) {
	tests := []struct{
		snakeMovement gamecomponents.SnakeMovement
		match *gamerunner.Match
	}{
		{
			gamecomponents.Up,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 5, YAt: 6}},
				Score: uint64(6),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 30},
					gamecomponents.Location{XAt: 5, YAt: 29},
					gamecomponents.Location{XAt: 5, YAt: 28},
					gamecomponents.Location{XAt: 5, YAt: 27},
				},
			},
		},
		{
			gamecomponents.Right,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 6, YAt: 5}},
				Score: uint64(6),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 30, YAt: 5},
					gamecomponents.Location{XAt: 29, YAt: 5},
					gamecomponents.Location{XAt: 28, YAt: 5},
					gamecomponents.Location{XAt: 27, YAt: 5},
				},
			},
		},
		{
			gamecomponents.Left,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 4, YAt: 5}},
				Score: uint64(6),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 0, YAt: 5},
					gamecomponents.Location{XAt: 1, YAt: 5},
					gamecomponents.Location{XAt: 2, YAt: 5},
					gamecomponents.Location{XAt: 3, YAt: 5},
				},
			},
		},
		{
			gamecomponents.Down,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: gamecomponents.Food{XAt: 5, YAt: 4}},
				Score: uint64(6),
				Snake: &gamecomponents.Snake{
					gamecomponents.Location{XAt: 5, YAt: 0},
					gamecomponents.Location{XAt: 5, YAt: 1},
					gamecomponents.Location{XAt: 5, YAt: 2},
					gamecomponents.Location{XAt: 5, YAt: 3},
				},
			},
		},
	}

	for _, test := range tests {
		scoreBeforeMoving := test.match.Score
		var gameOver bool = test.match.MakeMove(test.snakeMovement)
		if !gameOver {
			t.Error("Snake should have crashed after MakeMove and Game Over should happen.")
		}
		if scoreBeforeMoving != test.match.Score {
			t.Errorf("test.match.Score (%d) should be equal than scoreBeforeMoving (%d) " + 
				"after MakeMove(...) call.", test.match.Score, scoreBeforeMoving)
		}
	}
}

