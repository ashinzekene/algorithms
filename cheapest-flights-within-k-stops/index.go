package algorithms

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	flightPrices := make([]int, n)
	for i := range flightPrices {
		flightPrices[i] = math.MaxInt
	}

	flightPrices[src] = 0
	for i := k+1; i >= 0; i-- {
		tmpFlightPrices := make([]int, n)
		copy(tmpFlightPrices, flightPrices)
		for _, flight := range flights {
			s := flight[0]
			d := flight[1]
			price := flight[2]
		
			if flightPrices[s] == math.MaxInt {
				continue
			}
			currPrice := flightPrices[s] + price
			tmpFlightPrices[d] = min(currPrice, tmpFlightPrices[d])
		}

		flightPrices = tmpFlightPrices
	}

	if flightPrices[dst] == math.MaxInt {
		return -1
	}
	return flightPrices[dst]
}