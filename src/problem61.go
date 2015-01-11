package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	sum := 0
	for _, v := range findSet() {
		fmt.Println(v)
		sum += v
	}
	elapsed := time.Since(start)
	fmt.Println("Somme", sum)
	fmt.Println("Took", elapsed)
}

func findSet() []int {
	for i := 10; i <= 99; i++ {
		set := []int{}
		functions := []func(int) bool{isTriangle, isSquare, isPentagonal, isHexagonal, isHeptagonal, isOctogonal}
		result := findRec(set, i, functions)
		if result != nil {
			return result
		}
	}
	return nil
}

func findRec(set []int, hundreds int, functions []func(int) bool) []int{
	if len(functions) == 0 {
		if hundreds == set[0] / 100 {
			return set
		} else {
			return nil
		}
	}

	for i := 10; i <= 99; i++ {
		nb := 100 * hundreds + i
		for idx, f := range functions {
			if f(nb) {
				subFunctions := []func(int) bool{}
				for _, sf := range functions[:idx] {
					subFunctions = append(subFunctions, sf)
				}
				for _, sf := range functions[idx + 1:] {
					subFunctions = append(subFunctions, sf)
				}

				newSet := []int{}
				for _, s := range set {
					newSet = append(newSet, s)
				}
				newSet = append(newSet, nb)

				result := findRec(newSet, i, subFunctions)
				if result != nil {
					return result
				}
			}
		}
	}
	return nil
}

func isTriangle(nb int) bool {
	n := (-1 + math.Sqrt(float64(1 + 8 * nb))) / 2
	return n == math.Floor(n)
}

func isSquare(nb int) bool {
	n := math.Sqrt(float64(nb))
	return n == math.Floor(n)
}

func isPentagonal(nb int) bool {
	n := (1 + math.Sqrt(float64(1 + 24 * nb))) / 6
	return n == math.Floor(n)
}

func isHexagonal(nb int) bool {
	n := (1 + math.Sqrt(float64(1 + 8 * nb))) / 4
	return n == math.Floor(n)
}

func isHeptagonal(nb int) bool {
	n := (3 + math.Sqrt(float64(9 + 40 * nb))) / 10
	return n == math.Floor(n)
}

func isOctogonal(nb int) bool {
	n := (2 + math.Sqrt(float64(4 + 12 * nb))) / 6
	return n == math.Floor(n)
}


