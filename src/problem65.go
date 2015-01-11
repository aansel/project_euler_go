package main

import (
	"fmt"
	"time"
	"math/big"
)

func main() {
	start := time.Now()
	fmt.Println(getResult(100))
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}

func getResult(ith int64) int {
	return getSumNumerator(getIthValue(ith))
}

func getSumNumerator(n *big.Rat) int {
	s := n.Num().String()
	sum := 0
	for _, c := range s {
		sum += int(c) - 48
	}
	return sum
}

func getA(pos int64) int64 {
	if (pos - 1) % 3 == 0 {
		return ((pos - 1) / 3 + 1) * 2
	} else {
		return 1
	}
}


func getIthValue(ith int64) *big.Rat {
	r := getIthIterValue(ith - 1)
	return r.Add(r, big.NewRat(2, 1))
}

func getIthIterValue(ith int64) *big.Rat {
	if ith == 0 {
		return big.NewRat(0, 1)
	}

	r := big.NewRat(getA(ith - 1), 1)
	var i int64
	for i = ith - 2; i >= 0; i-- {
		r = r.Inv(r)
		r = r.Add(r, big.NewRat(getA(i), 1))
	}
	r = r.Inv(r)
	return r
}






