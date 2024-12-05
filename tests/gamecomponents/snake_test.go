package gamecomponents

import (
	"reflect"
	"testing"

	"github.com/gonza56d/go_snake/internal/gamecomponents"
)

func TestMoveSnakeSuccess(t *testing.T) {
	tests := []struct {
		snake gamecomponents.Snake
		movingDirection gamecomponents.SnakeMovement
		expected gamecomponents.Snake
	}{    ///////////////////////////// FACING UP ///////////////////////////
		{ // snake is facing up and moves up (continues moving up).
			gamecomponents.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 7},
			},
			gamecomponents.Up,
			gamecomponents.Snake{
				{XAt: 10, YAt: 11},
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
			},
		},
		{ // snake is facing up and moves right (then moves right).
			gamecomponents.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 7},
			},
			gamecomponents.Right,
			gamecomponents.Snake{
				{XAt: 11, YAt: 10},
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
			},
		},
		{ // snake is facing up and moves down (illegal movement, continues moving up).
			gamecomponents.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 7},
			},
			gamecomponents.Down,
			gamecomponents.Snake{
				{XAt: 10, YAt: 11},
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
			},
		},
		{ // snake is facing up and moves left (then moves left).
			gamecomponents.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 7},
			},
			gamecomponents.Left,
			gamecomponents.Snake{
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
			},
		},
		////////////////////// FACING RIGHT ///////////////////////////////
		{ // snake is facing right and moves up (then moves up).
			gamecomponents.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 7, YAt: 10},
			},
			gamecomponents.Up,
			gamecomponents.Snake{
				{XAt: 10, YAt: 11},
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
			},
		},
		{ // snake is facing right and moves right (continues moving right).
			gamecomponents.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 7, YAt: 10},
			},
			gamecomponents.Right,
			gamecomponents.Snake{
				{XAt: 11, YAt: 10},
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
			},
		},
		{ // snake is facing right and moves down (then moves down).
			gamecomponents.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 7, YAt: 10},
			},
			gamecomponents.Down,
			gamecomponents.Snake{
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
			},
		},
		{ // snake is facing right and moves left (illegal movement, continues moving right).
			gamecomponents.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 7, YAt: 10},
			},
			gamecomponents.Left,
			gamecomponents.Snake{
				{XAt: 11, YAt: 10},
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
			},
		},
		///////////////////////////// FACING DOWN /////////////////////////////////
		{ // snake is facing down and moves up (illegal movement, continues moving down).
			gamecomponents.Snake{
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
			},
			gamecomponents.Up,
			gamecomponents.Snake{
				{XAt: 10, YAt: 6},
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
			},
		},
		{ // snake is facing down and moves right (then moves right).
			gamecomponents.Snake{
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
			},
			gamecomponents.Right,
			gamecomponents.Snake{
				{XAt: 11, YAt: 7},
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
			},
		},
		{ // snake is facing down and moves down (then continues moving down).
			gamecomponents.Snake{
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
			},
			gamecomponents.Down,
			gamecomponents.Snake{
				{XAt: 10, YAt: 6},
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
			},
		},
		{ // snake is facing down and moves left (then moves left).
			gamecomponents.Snake{
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
			},
			gamecomponents.Left,
			gamecomponents.Snake{
				{XAt: 9, YAt: 7},
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
			},
		},
		//////////////////////////// FACING LEFT ///////////////////////////////
		{ // snake is facing left and moves up (then moves up).
			gamecomponents.Snake{
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
			},
			gamecomponents.Up,
			gamecomponents.Snake{
				{XAt: 7, YAt: 11},
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
			},
		},
		{ // snake is facing left and moves right (illegal movement, continues moving left).
			gamecomponents.Snake{
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
			},
			gamecomponents.Right,
			gamecomponents.Snake{
				{XAt: 6, YAt: 10},
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
			},
		},
		{ // snake is facing left and moves down (then moves down).
			gamecomponents.Snake{
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
			},
			gamecomponents.Down,
			gamecomponents.Snake{
				{XAt: 7, YAt: 9},
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
			},
		},
		{ // snake is facing left and moves left (then continues moving left).
			gamecomponents.Snake{
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
			},
			gamecomponents.Left,
			gamecomponents.Snake{
				{XAt: 6, YAt: 10},
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
			},
		},
	}	

	var gameMap gamecomponents.Map = gamecomponents.Map{XSize: 30, YSize: 30}
	for _, test := range tests {
		gamecomponents.MoveSnake(&test.snake, test.movingDirection)
		if !reflect.DeepEqual(test.snake, test.expected) { 
			t.Errorf("Unexpected snake state after MoveSnake function call. (Moving %d).\n" +
				"Snake: %d\nExpected: %d", test.movingDirection, test.snake, test.expected)
		}
		if gamecomponents.IsSnakeCrashed(&test.snake, &gameMap) {
			t.Error("Snake crashed after legal movement.")
		}
	}
}

func TestMoveSnakeCrashAgainstOwnBody(t *testing.T) {
	tests := []struct {
		snake gamecomponents.Snake
		movingDirection gamecomponents.SnakeMovement
	}{
		{ // crashing with themselves facing UP
			gamecomponents.Snake{
				{XAt: 20, YAt: 20}, //  XXX
				{XAt: 20, YAt: 19}, //   OX
				{XAt: 21, YAt: 19}, //   XX
				{XAt: 21, YAt: 20}, //
				{XAt: 21, YAt: 21}, //
				{XAt: 20, YAt: 21}, //
				{XAt: 19, YAt: 21}, //
			},
			gamecomponents.Up,
		},
		{ // crashing with themselves facing RIGHT
			gamecomponents.Snake{
				{XAt: 15, YAt: 15}, //
				{XAt: 14, YAt: 15}, //  XXX
				{XAt: 14, YAt: 16}, //  XOX
				{XAt: 15, YAt: 16}, //    X
				{XAt: 16, YAt: 16}, //
				{XAt: 16, YAt: 15}, //
				{XAt: 16, YAt: 14}, //
			},
			gamecomponents.Right,
		},
		{ // crashing with themselves facing DOWN
			gamecomponents.Snake{
				{XAt: 15, YAt: 15}, //
				{XAt: 15, YAt: 16}, //  XX
				{XAt: 16, YAt: 16}, //  OX
				{XAt: 16, YAt: 15}, // XXX
				{XAt: 16, YAt: 14}, // 
				{XAt: 15, YAt: 14}, // 
				{XAt: 14, YAt: 14}, //
			},
			gamecomponents.Down,
		},
		{ // crashing with themselves facing LEFT
			gamecomponents.Snake{
				{XAt: 15, YAt: 15}, //  
				{XAt: 16, YAt: 15}, // X
				{XAt: 16, YAt: 14}, // XOX
				{XAt: 15, YAt: 14}, // XXX
				{XAt: 14, YAt: 14}, //  
				{XAt: 14, YAt: 15}, //  
				{XAt: 14, YAt: 16}, //  
			},
			gamecomponents.Left,
		},
	}

	var gameMap gamecomponents.Map = gamecomponents.Map{XSize: 30, YSize: 30}
	for _, test := range tests{
		if gamecomponents.IsSnakeCrashed(&test.snake, &gameMap) {
			t.Error("Snake should NOT have crashed before moving.")
		}
		gamecomponents.MoveSnake(&test.snake, test.movingDirection)
		if !gamecomponents.IsSnakeCrashed(&test.snake, &gameMap) {
			t.Error("Snake should have crashed after moving.")
		}
	}
}

func TestMoveSnakeCrashAgainstMapWall(t *testing.T) {
	tests := []struct {
		snake gamecomponents.Snake
		movingDirection gamecomponents.SnakeMovement
	}{
		{ // moving UP when it's already at Y=30
			gamecomponents.Snake{
				{XAt: 26, YAt:30},
				{XAt: 26, YAt:29},
				{XAt: 26, YAt:28},
				{XAt: 26, YAt:27},
			},
			gamecomponents.Up,
		},
		{ // moving RIGHT when it's already at X=30
			gamecomponents.Snake{
				{XAt: 30, YAt:29},
				{XAt: 29, YAt:29},
				{XAt: 28, YAt:29},
				{XAt: 27, YAt:29},
			},
			gamecomponents.Right,
		},
		{ // moving DOWN when it's already at Y=0
			gamecomponents.Snake{
				{XAt: 3, YAt:0},
				{XAt: 3, YAt:1},
				{XAt: 3, YAt:2},
				{XAt: 3, YAt:3},
			},
			gamecomponents.Down,
		},
		{ // moving LEFT when it's already at X=0
			gamecomponents.Snake{
				{XAt: 0, YAt:4},
				{XAt: 1, YAt:4},
				{XAt: 2, YAt:4},
				{XAt: 3, YAt:4},
			},
			gamecomponents.Left,
		},
	}

	var gameMap gamecomponents.Map = gamecomponents.Map{XSize: 30, YSize: 30}
	for _, test := range tests{
		if gamecomponents.IsSnakeCrashed(&test.snake, &gameMap) {
			t.Error("Snake should NOT have crashed before moving.")
		}
		gamecomponents.MoveSnake(&test.snake, test.movingDirection)
		if !gamecomponents.IsSnakeCrashed(&test.snake, &gameMap) {
			t.Error("Snake should have crashed after moving.")
		}
	}
}

