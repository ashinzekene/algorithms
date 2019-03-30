package algorithms

import (
	"math"
)

func CatsAndAMouse(x int32, y int32, z int32) string {
	xDiff := math.Abs(float64(z - x))
	yDiff := math.Abs(float64(z - y))
	if xDiff == yDiff {
		return "Mouse C"
	} else if xDiff > yDiff {
		return "Cat B"
	} else {
		return "Cat A"
	}
}
