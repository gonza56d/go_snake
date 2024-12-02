package gameassets

import (
	"reflect"
	"testing"

	"github.com/gonza56d/go_snake/internal/gameassets"
)

func TestMoveSnakeSuccess(t *testing.T) {
	tests := []struct {
		snake gameassets.Snake
		movingDirection gameassets.SnakeMovement
		expected gameassets.Snake
	}{    ///////////////////////////// FACING UP ///////////////////////////
		{ // snake is facing up and moves up (continues moving up).
			gameassets.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 7},
			},
			gameassets.Up,
			gameassets.Snake{
				{XAt: 10, YAt: 11},
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
			},
		},
		{ // snake is facing up and moves right (then moves right).
			gameassets.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 7},
			},
			gameassets.Right,
			gameassets.Snake{
				{XAt: 11, YAt: 10},
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
			},
		},
		{ // snake is facing up and moves down (illegal movement, continues moving up).
			gameassets.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 7},
			},
			gameassets.Down,
			gameassets.Snake{
				{XAt: 10, YAt: 11},
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
			},
		},
		{ // snake is facing up and moves left (then moves left).
			gameassets.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 7},
			},
			gameassets.Left,
			gameassets.Snake{
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 8},
			},
		},
		////////////////////// FACING RIGHT ///////////////////////////////
		{ // snake is facing right and moves up (then moves up).
			gameassets.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 7, YAt: 10},
			},
			gameassets.Up,
			gameassets.Snake{
				{XAt: 10, YAt: 11},
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
			},
		},
		{ // snake is facing right and moves right (continues moving right).
			gameassets.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 7, YAt: 10},
			},
			gameassets.Right,
			gameassets.Snake{
				{XAt: 11, YAt: 10},
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
			},
		},
		{ // snake is facing right and moves down (then moves down).
			gameassets.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 7, YAt: 10},
			},
			gameassets.Down,
			gameassets.Snake{
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
			},
		},
		{ // snake is facing right and moves left (illegal movement, continues moving right).
			gameassets.Snake{
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 7, YAt: 10},
			},
			gameassets.Left,
			gameassets.Snake{
				{XAt: 11, YAt: 10},
				{XAt: 10, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 8, YAt: 10},
			},
		},
		///////////////////////////// FACING DOWN /////////////////////////////////
		{ // snake is facing down and moves up (illegal movement, continues moving down).
			gameassets.Snake{
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
			},
			gameassets.Up,
			gameassets.Snake{
				{XAt: 10, YAt: 6},
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
			},
		},
		{ // snake is facing down and moves right (then moves right).
			gameassets.Snake{
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
			},
			gameassets.Right,
			gameassets.Snake{
				{XAt: 11, YAt: 7},
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
			},
		},
		{ // snake is facing down and moves down (then continues moving down).
			gameassets.Snake{
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
			},
			gameassets.Down,
			gameassets.Snake{
				{XAt: 10, YAt: 6},
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
			},
		},
		{ // snake is facing down and moves left (then moves left).
			gameassets.Snake{
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
				{XAt: 10, YAt: 10},
			},
			gameassets.Left,
			gameassets.Snake{
				{XAt: 9, YAt: 7},
				{XAt: 10, YAt: 7},
				{XAt: 10, YAt: 8},
				{XAt: 10, YAt: 9},
			},
		},
		//////////////////////////// FACING LEFT ///////////////////////////////
		{ // snake is facing left and moves up (then moves up).
			gameassets.Snake{
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
			},
			gameassets.Up,
			gameassets.Snake{
				{XAt: 7, YAt: 11},
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
			},
		},
		{ // snake is facing left and moves right (illegal movement, continues moving left).
			gameassets.Snake{
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
			},
			gameassets.Right,
			gameassets.Snake{
				{XAt: 6, YAt: 10},
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
			},
		},
		{ // snake is facing left and moves down (then moves down).
			gameassets.Snake{
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
			},
			gameassets.Down,
			gameassets.Snake{
				{XAt: 7, YAt: 9},
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
			},
		},
		{ // snake is facing left and moves left (then continues moving left).
			gameassets.Snake{
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
				{XAt: 10, YAt: 10},
			},
			gameassets.Left,
			gameassets.Snake{
				{XAt: 6, YAt: 10},
				{XAt: 7, YAt: 10},
				{XAt: 8, YAt: 10},
				{XAt: 9, YAt: 10},
			},
		},
	}	

	var gameMap gameassets.Map = gameassets.Map{XSize: 30, YSize: 30}
	for _, test := range tests {
		gameassets.MoveSnake(&test.snake, test.movingDirection)
		if !reflect.DeepEqual(test.snake, test.expected) { 
			t.Errorf("Unexpected snake state after MoveSnake function call. (Moving %d).\n" +
				"Snake: %d\nExpected: %d", test.movingDirection, test.snake, test.expected)
		}
		if gameassets.IsSnakeCrashed(&test.snake, &gameMap) {
			t.Error("Snake crashed after legal movement.")
		}
	}
}

