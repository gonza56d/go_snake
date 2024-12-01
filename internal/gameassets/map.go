package gameassets

import (
	"math"
	"math/rand"
)

type Map struct {
	x_size uint8
	y_size uint8
}

var mapTotalSquares [4]uint16 = [4]uint16{900, 2025, 3600, 6400}

func GenerateNewMap(newMapSizeOption uint8) *Map {
	 if newMapSizeOption > 3 {
		panic("mapSize options are between 0 (smaller) and 3 (biggest).")
	}
	var newMapSize uint16 = mapTotalSquares[newMapSizeOption]
	var newMapSizeSqrtValue uint8 = uint8(math.Sqrt(float64(newMapSize)))

	var randomOffsetX int8 = 5
	if rand.Intn(2) == 0 {
		randomOffsetX = -5
	}
	var randomOffsetY int8 = -5
	if randomOffsetX == -5 {
		randomOffsetY = 5
	}

	var newMapXSize uint8 = newMapSizeSqrtValue + uint8(randomOffsetX)
	var newMapYSize uint8 = newMapSizeSqrtValue + uint8(randomOffsetY)
	return &Map{x_size: newMapXSize, y_size: newMapYSize}
}

