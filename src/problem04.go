package main;

import (
	"fmt"
	"strconv"
)


func main() {
	fmt.Println(getLargestPalindrome(999))
}


func getLargestPalindrome(max int) int {
	largestPalindrome := 0
	for i := max; i > 0; i-- {
		if i * i <= largestPalindrome {
			break
		}
		for j := i; j > 0; j-- {
			produit := i * j
			if produit <= largestPalindrome {
				break
			} else if isPalindrome(produit) {
				largestPalindrome = produit
				break
			}
		}
	}
	return largestPalindrome
}


func isPalindrome(nb int) bool {
	nbStr := strconv.Itoa(nb)
	for i := 0; i < len(nbStr) / 2; i++ {
		if string(nbStr[i]) != string(nbStr[len(nbStr) - i - 1]) {
			return false
		}
	}
	return true
}
