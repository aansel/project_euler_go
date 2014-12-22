package main;

import (
	"fmt"
)

func main() {
	first := 1
	second := 2
	sum := second
	for  i := 0; i < 4000000; i++ {
		i = first + second
		first = second
		second = i

		if i % 2 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
