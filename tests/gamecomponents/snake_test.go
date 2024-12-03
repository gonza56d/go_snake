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

