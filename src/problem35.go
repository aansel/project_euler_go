package main

import "fmt"
import "strconv"

func main() {
	getCircularPrimes()
}

func checkRotNumber(number int) {
	ok := true
	for _, rot := range getRotations(number) {
		if !isPrime(rot) {
			ok = false
			break
		}
	}
	if ok {
		fmt.Println(number)
	}
}

func getCircularPrimes() {
	possibleDigits := []int{1, 3, 5, 7, 9}
	for _, digit1 := range possibleDigits {
		checkRotNumber(digit1)
		for _, digit2 := range possibleDigits {
			number2 := digit2 + 10 * digit1
			checkRotNumber(number2)
			for _, digit3 := range possibleDigits {
				number3 := digit3 + 10 * digit2 + 100 * digit1
				checkRotNumber(number3)
				for _, digit4 := range possibleDigits {
					number4 := digit4 + 10 * digit3 + 100 * digit2 + 1000 * digit1
					checkRotNumber(number4)
					for _, digit5 := range possibleDigits {
						number5 := digit5 + 10 * digit4 + 100 * digit3 + 1000 * digit2 + 10000 * digit1
						checkRotNumber(number5)
						for _, digit6 := range possibleDigits {
							number6 := number5 * 10 + digit6
							checkRotNumber(number6)
						}

					}
				}
			}
		}
	}
}

func getRotations(nb int) []int {
	str := strconv.Itoa(nb)
	rotations := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		var i64 int64
		str = str[1:] + str[:1]
		i64, _ = strconv.ParseInt(str, 10, 32)
		rotations[i] = int(i64)
	}
	return rotations
}

func isPrime(nb int) bool {
	for i := 2; i < nb / 2; i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
