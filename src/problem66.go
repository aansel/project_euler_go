package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	start := time.Now()
	var a float64  =
	fmt.Println()
	//fmt.Println(findMinimalXSolution(367))
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}

func findDForLargestX(limit int) int {
	dForMaxX := 0
	var maxX uint64 = 0
	for i := 2; i <= limit; i++ {
		x, y := findMinimalXSolution(uint64(i))
		if x > maxX {
			maxX = x
			dForMaxX = i
		}
		fmt.Println(i, x, y)
	}
	return dForMaxX
}

func findMinimalXSolution(d uint64) (uint64, uint64) {
	if isSquare(d) {
		return 0, 0
	}
	for x := uint64(math.Sqrt(float64(d)));; x++ {
		//y := math.Sqrt(float64((x * x - 1)) / float64(d))
		if x % 100000 == 0 {
			fmt.Println(x)
		}
		y := math.Sqrt(float64(x - 1)) * math.Sqrt(float64(x + 1)) / math.Sqrt(float64(d))
		if isInteger(y) && x * x - d * uint64(y) * uint64(y) == 1{
			return x, uint64(y)
		}
	}
}

//func isSquare(x *big.Rat) bool {
//
//	if x.Denom().Int64() != 1 {
//		return false
//	}
//
//	var seed float64 = 100
//	f, _ := x.Float64()
//	if  f < 100 {
//		seed, _ = x.Mul(x, big.NewRat(1, 2)).Float64()
//	} else {
//		seed = 100
//	}
//
//	var pSqrt float64 = 0
//	for sqrt := seed;; {
//		fmt.Println("sqrt", sqrt)
//		if sqrt == 1 {
//			return false
//		}
//		sqrt, _ = x.Mul(x, big.NewRat(1, int64(sqrt))).Float64()
//		if sqrt == pSqrt {
//			fmt.Println(sqrt, pSqrt)
//			return big.NewRat(int64(sqrt) * int64(sqrt), 1) == x
//		}
//	}
//}

func isSquare(nb uint64) bool {
	sqrt := math.Sqrt(float64(nb))
	return sqrt == math.Floor(sqrt)
}

func isInteger(nb float64) bool {
	return math.Floor(nb) == nb
}






