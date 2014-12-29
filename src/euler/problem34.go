package main

import "fmt"

func main() {
	getCuriousNumbers()
}

func getCuriousNumbers() {
	for i := 0; i < 100000000; i++ {
		if i == getFactSum(i) {
			fmt.Println(i)
		}
	}
}

func getFactSum(nb int) int {
	if nb < 10 {

		return fact(nb)
	} else {
		lastDigit := nb - nb / 10 * 10
		return fact(lastDigit) + getFactSum(nb / 10)
	}
}


func fact(nb int) int {
	if nb == 0 {
		return 1
	} else if nb == 1 {
		return 1
	} else if nb == 2 {
		return 2
	} else if nb == 3 {
		return 6
	} else if nb == 4 {
		return 24
	} else if nb == 5 {
		return 120
	} else if nb == 6 {
		return 720
	} else if nb == 7 {
		return 5040
	} else if nb == 8 {
		return 40320
	} else {
		return 362880
	}
}
