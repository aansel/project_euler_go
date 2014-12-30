package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getFirstDecomposition4())
}

func getFirstDecomposition3() int64{
	var i int64 = 1
	dec1 := getDecomposition(i)
	dec2 := getDecomposition(i + 1)
	dec3 := getDecomposition(i + 2)

	for ; ; i++ {
		if are3DecompositionDistinct(dec1, dec2, dec3) {
			return i
		}

		dec1 = dec2
		dec2 = dec3
		dec3 = getDecomposition(i + 3)
	}
}

func getFirstDecomposition4() int64{
	var i int64 = 2
	dec1 := getDecomposition(i)
	dec2 := getDecomposition(i + 1)
	dec3 := getDecomposition(i + 2)
	dec4 := getDecomposition(i + 3)

	for ; ; i++ {
		if are4DecompositionDistinct(dec1, dec2, dec3, dec4) {
			for _, p1 := range dec1 {
				fmt.Println(i, p1)
			}
			for _, p2 := range dec2 {
				fmt.Println(i + 1, p2)
			}
			for _, p3 := range dec3 {
				fmt.Println(i + 2, p3)
			}
			for _, p4 := range dec4 {
				fmt.Println(i + 3, p4)
			}
			return i
		}

		dec1 = dec2
		dec2 = dec3
		dec3 = dec4
		dec4 = getDecomposition(i + 4)
	}
}

func are3DecompositionDistinct(dec1 []int64, dec2 []int64, dec3 []int64) bool {
	if len(dec1) < 3 || len(dec2) < 3 || len(dec3) < 3 {
		return false
	}
	return are2DecompositionDistinct(dec1, dec2) && are2DecompositionDistinct(dec1, dec3) && are2DecompositionDistinct(dec2, dec3)
}

func are4DecompositionDistinct(dec1 []int64, dec2 []int64, dec3 []int64, dec4 []int64) bool {
	if len(dec1) < 4 || len(dec2) < 4 || len(dec3) < 4 || len(dec4) < 4 {
		return false
	}
	return are2DecompositionDistinct(dec1, dec2) && are2DecompositionDistinct(dec1, dec3) && are2DecompositionDistinct(dec1, dec4) && are2DecompositionDistinct(dec2, dec3) && are2DecompositionDistinct(dec2, dec4) && are2DecompositionDistinct(dec3, dec4)
}

func are2DecompositionDistinct(dec1 []int64, dec2 []int64) bool {
	for _, i := range dec1 {
		for _, j := range dec2 {
			if i == j {
				return false
			}
		}
	}
	return true
}

func getDecomposition(nb int64) []int64 {
	var dec []int64
	var i int64

	var divisor2 int64 = 1
	for ; nb % 2 == 0; {
		nb = nb / 2
		divisor2 *= 2
	}
	if divisor2 > 1 {
		dec = append(dec, divisor2)
	}

	for i = 3; nb > 1; i+=2 {
		var divisor int64 = 1
		for ; nb % i == 0 && isPrime(i); {
			nb = nb / i
			divisor *= i
		}
		if divisor > 1 {
			dec = append(dec, divisor)
		}
	}
	return dec
}

func isPrime(nb int64) bool {
	if nb < 2 {
		return false
	}
	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(nb))); i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
