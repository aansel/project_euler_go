package main

import "fmt"
import "strconv"
import "strings"
import "math"

func main() {
	for i := 1; i <= 987654321; i+=2 {
		if isPandigital(i) && isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPandigital(nb int) bool {
	nbStr := strconv.Itoa(nb)
	if len(nbStr) > 9 {
		return false
	}

	for i := 1; i <= len(nbStr); i++ {
		if !strings.Contains(nbStr, strconv.Itoa(i)) {
			return false
		}
	}

	return true
}

func isPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(nb))); i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
