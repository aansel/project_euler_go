package main

import (
	"fmt"
)


func main() {
	r := 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19
	for i := r; ; i += r {
		if isEvenlyDivisible(i, 20) {
			fmt.Println(i)
			break
		}
	}
}


func isEvenlyDivisible(nb int, maxDivisor int) bool {
	for i := maxDivisor ; i >= 2 ; i-- {
		if nb % i != 0 {
			return false
		}
	}
	return true
}
