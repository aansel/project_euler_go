package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	start := time.Now()
	fmt.Println(getNbOddPeriod(10000))
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}

func getNbOddPeriod(limit int) int {
	nbOdd := 0
	for i := 2; i <= limit; i++ {
		sqrt := math.Sqrt(float64(i))
		if math.Floor(sqrt) != sqrt {
			_, period := getSquareRootExpansion(i)
			if len(period) % 2 != 0 {
				nbOdd++
			}
		}
	}
	return nbOdd
}

func getSquareRootExpansion(nb int) (int, []int) {
	a0 := int(math.Sqrt(float64(nb)))

	m := 0
	d := 1
	a := a0

	m1 := d * a - m
	d1 := (nb - m1 * m1) / d
	a1 := (a0 + m1) / d1

	m = m1
	d = d1
	a = a1
	period := []int{a1}

	for ;; {
		m = d * a - m
		d = (nb - m * m) / d
		a = (a0 + m) / d

		if m == m1 && d == d1 && a == a1 {
			return a0, period
		} else {
			period = append(period, a)
		}
	}

	return a0, period
}








