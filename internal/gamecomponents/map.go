package gamecomponents

type Map struct {
	XSize   uint8
	YSize   uint8
	FoodAt *Food;
}

type Location struct {
	XAt int16 // 0 -> most left
	YAt int16 // 0 -> most bottom
}

func GenerateNewMap(newMapSizeOption uint8) *Map {
	if newMapSizeOption > 3 {
		panic("mapSize options are between 0 (smaller) and 3 (biggest).")
	}

	dimentions := map[uint8][2]uint8{
		0: {30, 15},
		1: {35, 17},
		2: {40, 20},
		3: {50, 22},
	}

	var newMapXSize uint8 =	dimentions[newMapSizeOption][0]
	var newMapYSize uint8 = dimentions[newMapSizeOption][1] 
	return &Map{YSize: newMapXSize, XSize: newMapYSize}
}

