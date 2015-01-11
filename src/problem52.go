package main

import (
	"fmt"
	"strconv"
)


func main() {
	fmt.Println(findFirstIsoDigits(7))
}

func findFirstIsoDigits(mult int64) int64 {
	var nb int64 = 125874
	for ;; nb++ {
		var i int64
		for i = 2; i <= mult; i++ {
			if !containsSameDigits(nb, nb * i) {
				break
			} else if i == mult {
				return nb
			}
		}
	}
}

func containsSameDigits(nb1 int64, nb2 int64) bool {
	digits1 := getDigits(nb1)
	digits2 := getDigits(nb2)
	for idx, nDigit := range digits1 {
		if digits2[idx] != nDigit {
			return false
		}
	}
	return true;
}

func getDigits(nb int64) [10]int {
	var digits [10]int
	nbStr := strconv.FormatInt(nb, 10)
	for _, car := range nbStr {
		digits[car - 48] += 1
	}
	return digits
}
