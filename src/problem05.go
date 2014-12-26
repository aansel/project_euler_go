package main

import (
	"fmt"
)


func main() {
	getSmallestMultiple(20)
}


func getSmallestMultiple(nb int) {
	r, nonPrimes := getPrimesFactor(nb)
	for i := r; ; i += r {
		ok := true
		for _, nonPrime := range nonPrimes {
			if (i % nonPrime != 0) {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(i)
			break
		}
	}
}

/**
 Le plus petit nombre divisible par tous les entiers de 1 à N a nécessaireement dans sa décomposition en facteurs premiers
 au moins une fois tous les nombres premiers de 1 à N
 */
func getPrimesFactor(maxValue int) (int, []int) {
	r := 1
	var nonPrimes []int
	for i := 2 ; i <= maxValue; i++ {
		if isPrime(i) {
			r *= i
		} else {
			nonPrimes = append(nonPrimes, i)
		}
	}
	return r, nonPrimes
}


func isPrime(nb int) bool {
	for i := 2 ; i < nb / 2 ; i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
