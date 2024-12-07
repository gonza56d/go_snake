package gamecomponents

import "math/rand"

type Food Location

func MoveFood(food *Food, snake *Snake, gameMap *Map) {
	var xRand, yRand int16	

	for {
		xRand = int16(rand.Intn(int(gameMap.XSize)))
		yRand = int16(rand.Intn(int(gameMap.YSize)))
		if !overlapsSnake(xRand, yRand, snake) {
			break
		}
	}
	food.XAt = xRand
	food.YAt = xRand
}

func overlapsSnake(xRand int16, yRand int16, snake *Snake) bool {
	for _, segment := range *snake {
		if xRand == segment.XAt || yRand == segment.YAt {
			return true
		}
	}
	return false
}

