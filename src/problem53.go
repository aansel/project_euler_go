package main

import (
	"fmt"
	"math/big"
)


func main() {
	fmt.Println(getNumberValuesBelow(1000000))
}

func getNumberValuesBelow(limit int64) int {
	result := 0
	var n int64
	var r int64
	l := big.NewInt(limit)
	for n = 1 ; n <= 100; n++ {
		for r = 1 ; r <= n; r++ {
			comb := comb(n, r)
			if comb.Cmp(l) == 1 {
				result++
			}
		}
	}
	return result
}

func comb(n int64, r int64) *big.Int {
	factN := factorial(n)
	factR := factorial(r)
	factNmR := factorial(n - r)
	return factN.Div(factN, (factR.Mul(factR, factNmR)))
}

func factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	} else {
		nb := big.NewInt(n)
		return nb.Mul(nb, factorial(n - 1))
	}
}
