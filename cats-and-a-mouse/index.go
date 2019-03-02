package main

import (
	"math"
)

func main() {

}

func catAndMouse(x int32, y int32, z int32) string {
	xDiff := math.Abs(float64(z - x))
	yDiff := math.Abs(float64(z - y))
	if xDiff == yDiff {
		return "Mouse Z"
	} else if xDiff > yDiff {
		return "Cat A"
	} else {
		return "Cat B"
	}
}
