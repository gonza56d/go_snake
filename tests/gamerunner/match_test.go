package gamerunner

import (
	"reflect"
	"testing"

	"github.com/gonza56d/go_snake/internal/gamecomponents"
	"github.com/gonza56d/go_snake/internal/gamerunner"
)

func TestNewMatch(t *testing.T) {
	tests := make(map[uint8][2]uint8)
	tests[0] = [2]uint8{25, 35}
	tests[1] = [2]uint8{40, 50}
	tests[2] = [2]uint8{55, 65}
	tests[3] = [2]uint8{75, 85}

	for key, value := range tests {
		match := gamerunner.NewMatch(key) 
		if len(*match.Snake) != 4 {
			t.Error("Initial snake length must be 4.")
		}
		if match.Score != 0 {
			t.Error("Initial score must be 0.")
		}
		if !(match.GameMap.XSize == value[0] && match.GameMap.YSize == value[1] || 
			match.GameMap.XSize == value[1] && match.GameMap.YSize == value[0]) {
			t.Errorf("Initial map size for option %d should be %dx%d or vice versa. \n" +
				"Found %dx%d.", key, value[0], value[1], match.GameMap.XSize, match.GameMap.YSize)
		}
		if !reflect.DeepEqual(
			*match.Snake,
			gamecomponents.Snake{
				{XAt: int16(match.GameMap.XSize / 2), YAt: int16(match.GameMap.YSize / 2)},
				{XAt: int16(match.GameMap.XSize / 2), YAt: int16(match.GameMap.YSize / 2) - 1},
				{XAt: int16(match.GameMap.XSize / 2), YAt: int16(match.GameMap.YSize / 2) - 2},
				{XAt: int16(match.GameMap.XSize / 2), YAt: int16(match.GameMap.YSize / 2) - 3},
			},
		) {
			t.Error("Snake is not positioned at the center of map.")
		}
	}
}

func TestMakeMoveAndNothingHappens(t *testing.T) {
	tests := []struct{
		snakeMovement gamecomponents.SnakeMovement
		match *gamerunner.Match
	}{
		{
			gamecomponents.Up,
			&gamerunner.Match{
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 1, YAt: 1}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 1, YAt: 1}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 1, YAt: 1}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 1, YAt: 1}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 5, YAt: 6}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 6, YAt: 5}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 4, YAt: 5}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 5, YAt: 4}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 5, YAt: 6}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 6, YAt: 5}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 4, YAt: 5}},
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
				GameMap: &gamecomponents.Map{XSize: 30, YSize: 30, FoodAt: &gamecomponents.Food{XAt: 5, YAt: 4}},
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

