package main

import (
	"fmt"
)


func main() {
	fmt.Println(getMaxProduct(1000))
}


func getMaxProduct(limit int) int {
	var maxProduct int
	maxNumberOfSuccessivePrimes := 0
	for _, b := range getPrimes(limit) {
		for a:= 1 - b; a <= limit; a++ {
			numberOfSuccessivePrimes := getNumberOfSucessivePrimes(a, b)
			if numberOfSuccessivePrimes > maxNumberOfSuccessivePrimes {
				maxNumberOfSuccessivePrimes = numberOfSuccessivePrimes
				maxProduct = a * b
			}
		}
	}
	return maxProduct
}

func getNumberOfSucessivePrimes(a int, b int) int {
	nbPrimes := 0
	for n := 0; ; n++ {
		if isPrime(n * n + a * n + b) {
			nbPrimes++
		} else {
			break
		}
	}
	return nbPrimes
}


func getPrimes(max int) []int {
	var primes []int
	for i := 2 ; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}


func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2 ; i < nb / 2 ; i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
