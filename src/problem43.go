package main

import "fmt"
import "strconv"
import "strings"

func main() {
	fmt.Println("Total: ", getSumPandigital())
}

func getPandigitals() []string {
	return getArrangements("0123456789")
}

func getArrangements(str string) []string {
	var arrangements []string
	if len(str) == 1 {
		arrangements = append(arrangements, str)
	} else {
		for i := 0; i < len(str); i++ {
			firstCar := str[i]
			subArrangements := getArrangements(str[0:i] + str[i+1:len(str)])
			for j := 0; j < len(subArrangements); j++ {
				arrangements = append(arrangements, string(firstCar) + subArrangements[j])
			}
		}
	}
	return arrangements
}

func getSumPandigital() int64 {
	var sum int64 = 0
	var i int64
	for _, pandi := range getPandigitals() {
		i, _ = strconv.ParseInt(pandi, 10, 64)
		if !isPandigital(i) {
			continue
		}
		if getSubNumber(i, 2, 4) % 2 == 0 && getSubNumber(i, 3, 5) % 3 == 0 && getSubNumber(i, 4, 6) % 5 == 0 && getSubNumber(i, 5, 7) % 7 == 0 && getSubNumber(i, 6, 8) % 11 == 0 && getSubNumber(i, 7, 9) % 13 == 0 && getSubNumber(i, 8, 10) % 17 == 0 {
			fmt.Println(i)
			sum += i
		}
	}
	return sum
}

func getSubNumber(nb int64, idxMin int, idxMax int) int {
	nbStr := strconv.FormatInt(nb, 10)
	subNumber, _ := strconv.ParseInt(nbStr[idxMin - 1:idxMax], 10, 32)
	return int(subNumber)
}


func isPandigital(nb int64) bool {
	nbStr := strconv.FormatInt(nb, 10)
	if len(nbStr) != 10 {
		return false
	}
	return strings.Contains(nbStr, "0") && strings.Contains(nbStr, "1") && strings.Contains(nbStr, "2") && strings.Contains(nbStr, "3") && strings.Contains(nbStr, "4") && strings.Contains(nbStr, "5") && strings.Contains(nbStr, "6") && strings.Contains(nbStr, "7") && strings.Contains(nbStr, "8") && strings.Contains(nbStr, "9")
}
