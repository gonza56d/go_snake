package gamecomponents

import (
	"testing"

	"github.com/gonza56d/go_snake/internal/gamecomponents"
)

func TestSpawnFood(t *testing.T) {
	snake := &gamecomponents.Snake{
		{XAt: 1, YAt: 1},
		{XAt: 2, YAt: 1},
		{XAt: 3, YAt: 1},
		{XAt: 4, YAt: 1},
	}
	gameMap := &gamecomponents.Map{XSize: 10, YSize: 10}
	for i := 0; i < 10000; i++ {
		food := gamecomponents.SpawnFood(snake, gameMap)
		for _, segment := range *snake {
			if food.XAt == segment.XAt && food.YAt == segment.YAt {
				t.Errorf("Spawned food overlapped snake.")
			}
		}
	}
}
