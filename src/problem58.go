package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getFirstSpiralWithPrimeBelow(0.1))
}

func getFirstSpiralWithPrimeBelow(limit float32) int {
	nbOnDiagonal := 0
	nbPrimesOnDiagonal := 0
	nb := 1
	for layer := 3; ; layer += 2 {
		lastOfPreviousLayer := nb
		for idx := 0 ; idx < 4 * (layer - 1); idx++ {
			nb += 1
			if (nb - lastOfPreviousLayer) % (layer - 1) == 0 {
				nbOnDiagonal++
				if isPrime(nb) {
					nbPrimesOnDiagonal++
				}
			}

		}
		ratio := float32(nbPrimesOnDiagonal) / float32(nbOnDiagonal + 1)
		if ratio < 0.1 {
			return layer
		}
	}
}

func isPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	var i int
	for i = 2; i <= int(math.Sqrt(float64(nb))); i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
