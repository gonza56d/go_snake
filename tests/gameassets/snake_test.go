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
	}{
		{
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
		{
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
		{
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
	}	

	var gameMap gameassets.Map = gameassets.Map{XSize: 30, YSize: 30}
	for _, test := range tests {
		gameassets.MoveSnake(&test.snake, test.movingDirection)
		if !reflect.DeepEqual(test.snake, test.expected) { 
			t.Errorf("Unexpected snake state after MoveSnake function call.\n" +
				"Snake: %d\nExpected: %d", test.snake, test.expected)
		}
		if gameassets.IsSnakeCrashed(&test.snake, &gameMap) {
			t.Error("Snake crashed after legal movement.")
		}
	}
}

