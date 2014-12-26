package main

import (
"fmt"
)


func main() {
	fmt.Println(getIthPermutation(999999, 10))
}

func getIthPermutation(nb int64, nbDigits int) int64 {
	if nbDigits == 0 {
		return 0
	}
	n := factorial(nbDigits - 1)
	digit := nb / n
	reste := nb % n
	return digit * int64(pow10(nbDigits - 1)) + getIthPermutation(reste, nbDigits - 1)
}


func factorial(x int) int64 {
	if x == 0 {
		return 1
	}
	return int64(x) * factorial(x-1)
}

func pow10(e int) int64 {
	var v int64 = 1
	for i := 0 ; i < e ; i++ {
		v *= 10
	}
	return v
}
