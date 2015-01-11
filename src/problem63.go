package main

import (
	"fmt"
	"time"
	"math/big"
)

func main() {
	start := time.Now()
	fmt.Println(findNumberOfPowerfulDigits())
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}

func findNumberOfPowerfulDigits() int {
	res := 0
	for digits := 1; ; digits++ {
		for i := 1; i <= 9; i++ {
			n := pow(i, digits)
			if len(n.String()) == int(digits) {
				fmt.Println(i, "^", digits, "=", n)
				res++
			}
		}
		if len(pow(9, digits).String()) < int(digits) {
			fmt.Println("Stop à ", digits, " digits")
			break
		}
	}
	return res
}

func pow(n int, p int) *big.Int {
	if p == 1 {
		return big.NewInt(int64(n))
	} else {
		bn := big.NewInt(int64(n))
		return bn.Mul(bn, pow(n, p - 1))
	}
}








