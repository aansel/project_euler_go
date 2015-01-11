package main

import "fmt"
import "math"

func main() {
	primes := getPrimesBelow(20000)
	fmt.Println(findLongestPrimeChain(primes, 1000000))
}

func findLongestPrimeChain(primes []int, limit int) int {
	prime := 0
	for l := 21; l < len(primes) && prime <= limit; l++ {
		res := getSumPrimeChain(primes, l, limit)
		if res != -1 && res <= limit {
			prime = res
		}
	}
	return prime
}

func getSumPrimeChain(primes []int, length int, limit int) int {
	for start := 0; start < len(primes) - length; start++ {
		sum := 0
		for i := start; i < start + length; i++ {
			sum += primes[i]
		}
		if isPrime(sum) {
			return sum
		}
	}
	return -1
}


func getPrimesBelow(limit int) []int {
	var primes []int
	primes = append(primes, 2)
	for i := 3; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
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
