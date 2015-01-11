package main

import (
	"fmt"
	"math/big"
	"strings"
)

func getReverse(nb *big.Int) *big.Int {
	r := new(big.Int)
	r.SetString(getReverseStr(nb.String()), 10)
	return r
}

func getReverseStr(orig string) string {
	var c []string = strings.Split( orig, "")
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}

	return strings.Join( c, "" );
}

func isPalindromic(s string) bool {
	for idx, car := range s {
		if int(s[len(s) - idx - 1]) != int(car) {
			return false
		}
	}
	return true
}

func isLychrelNumber(number int64) bool {
	nb := big.NewInt(number)
	for i:= 0; i < 50; i++ {
		reverseNb := getReverse(nb)
		nb := nb.Add(nb, reverseNb)
		if isPalindromic(nb.String()) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isLychrelNumber(196))
	nbLychrel := 0
	var i int64
	for i = 0 ; i <= 10000; i++ {
		if isLychrelNumber(i) {
			fmt.Println(i)
			nbLychrel++
		}
	}
	fmt.Println(nbLychrel)
}



