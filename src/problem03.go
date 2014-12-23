package main;

import (
	"fmt"
)


func main() {
	//fmt.Println(getLargestPrimeFactor(28))
	fmt.Println(getLargestPrimeFactor(600851475143))
}


func getLargestPrimeFactor(nb int) int {
	var i = 2
	for ;nb != 1; {
		if nb % i == 0 {
			nb = nb / i
		} else {
			i += 1
		}
	}
	return i
}
