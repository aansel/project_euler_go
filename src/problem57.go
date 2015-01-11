package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(getNbWithBigerNumerator(1000))
}

func getNbWithBigerNumerator(limit int) int {
	res := 0
	r := big.NewRat(1, 1)
	for i := 0; i < limit; i++ {
		r = r.Add(r, big.NewRat(1, 1))
		r = r.Inv(r)
		r = r.Add(r, big.NewRat(1, 1))
		if len(r.Num().String()) > len(r.Denom().String()) {
			res += 1
		}
	}
	return res
}
