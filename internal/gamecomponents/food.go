package gamecomponents

import "math/rand"

type Food Location

func MoveFood(food *Food, snake *Snake, gameMap *Map) {
	var xRand, yRand int16	
	occupied := make(map[Location]struct{})
	for _, segment := range *snake {
		occupied[segment] = struct{}{}
	}
	for {
		xRand = int16(rand.Intn(int(gameMap.XSize)))
		yRand = int16(rand.Intn(int(gameMap.YSize)))
		if _, exists := occupied[Location{YAt: yRand, XAt: xRand}]; !exists { 
			break
		}
	}
	food.XAt = xRand
	food.YAt = yRand
}

