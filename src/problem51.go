package main

import "fmt"
import "strconv"
import "math"

func main() {
	fmt.Println(findFirstInPrimeFamily(8))
//	var nb int64 = 56003
//	fmt.Println(isInPrimeFamily(nb, 8))
}

func findFirstInPrimeFamily(family int) int64 {
	var nb int64
	for nb = 1; nb <= 1000000 ; nb+=2 {
		if isPrime(nb) && isInPrimeFamily(nb, family) {
			return nb
		}
	}
	return -1
}

func isInPrimeFamily(nb int64, family int) bool {
	nbStr := strconv.FormatInt(nb, 10)
	for rep1 := 0; rep1 < len(nbStr) - 3; rep1++ {
		for rep2 := rep1 + 1; rep2 < len(nbStr) - 2; rep2++ {
			for rep3 := rep2 + 1; rep3 < len(nbStr) - 1; rep3++ {
//				for rep4 := rep3 + 1; rep4 < len(nbStr) - 1; rep4++ {
//					for rep5 := rep4 + 1; rep5 < len(nbStr) - 1; rep5++ {
						fmt.Println("*** Replace", rep1, rep2, rep3, "in", nb, "***")
						positions := []int{rep1, rep2, rep3}
						if isInPrimeFamilyPos(nb, positions, family) {
							return true
						}
//					}
//				}
			}
		}
	}
	return false
}

func isInPrimeFamilyPos(nb int64, positions []int, family int) bool {
	replacements :=  getReplacements(nb, positions)
	nbPrimes := 0
	for _, rep := range replacements {
		if isPrime(rep) {
			fmt.Println(rep)
			nbPrimes++
		}
	}
	fmt.Println(nbPrimes, "primes")
	return nbPrimes >= family
}

func getReplacements(nb int64, positions []int) []int64 {
	nbStr := strconv.FormatInt(nb, 10)
	var reps []int64
	for i := 0; i <= 9; i++ {
		repStr := nbStr
		toAdd := true
		for _, pos := range positions {
			if pos == 0 && i == 0 {
				toAdd = false
				break
			}
			repStr = repStr[:pos] + strconv.Itoa(i) +  repStr[pos + 1:]
		}
		if toAdd {
			rep, _ := strconv.ParseInt(repStr, 10, 32)
			reps = append(reps, rep)
		}
	}
	return reps
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
