package gameassets

import (
	"math"
	"math/rand"
)

type Map struct {
	XSize uint8
	YSize uint8
}

type Location struct {
	XAt int16 // 0 -> most left
	YAt int16 // 0 -> most bottom
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
	return &Map{XSize: newMapXSize, YSize: newMapYSize}
}

