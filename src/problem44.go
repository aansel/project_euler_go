package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getSmallestPentagonalDifference(100000000))
}

func getSmallestPentagonalDifference(limit int64) int64 {
	var bestDifference int64 = limit
	var i, j int64
	for i = 1; i < limit; i++ {
		var pj int64 = i * (3 * i - 1) / 2

		var pk int64 = bestDifference
		for j = i + 1; pk - pj < bestDifference; j++ {
			pk = j * (3 * j - 1) / 2
			if isPentagonal(pj + pk) && isPentagonal(pk - pj){
				fmt.Println("Trouvé", pj, pk)
				bestDifference = pk - pj
			}
		}
	}
	return bestDifference
}

func isPentagonal(nb int64) bool {
	if nb <= 0 {
		return false
	}
	n := (1 + math.Sqrt(float64(1 + 24 * nb))) / 6
	return float64(int64(n)) == n
}
