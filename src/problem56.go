package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(getMaxSumDigit())
}

func getMaxSumDigit() int {
	maxSum := 0
	var a int64
	var b int
	for a = 0; a < 100; a++ {
		for b = 0; b < 100; b++ {
			sum := getSumDigits(pow(a, b))
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func getSumDigits(nb *big.Int) int {
	sum := 0
	for _, digit := range nb.String() {
			sum += int(digit - 48)
	}
	return sum
}

func pow(nb int64, power int) *big.Int {
	result := big.NewInt(1)
	bigNb := big.NewInt(nb)
	for i := 0; i < power; i++ {
		result = result.Mul(result, bigNb)
	}
	return result
}
