package main

import (
	"fmt"
	"time"
	"sort"
	"strconv"
)

func main() {
	start := time.Now()
	sc := findSmallestCube()
	fmt.Println(sc * sc * sc)
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}

func getCanonicalForm(nb int64) string {
	return SortString(strconv.FormatInt(nb, 10))
}

var cubeMap map[string][]int64 = make(map[string][]int64)

func findSmallestCube() int64 {
	var i int64
	currentLength := 1
	for i = 0 ;; i++ {
		cube := i * i * i
		l := len(strconv.FormatInt(cube, 10))
		if l > currentLength {
			for _, cubes := range cubeMap {
				if len(cubes) == 5 {
					return cubes[0]
				}
			}
			cubeMap = make(map[string][]int64)
			currentLength = l
		}
		cubeMap[getCanonicalForm(cube)] = append(cubeMap[getCanonicalForm(cube)], i)
	}
	return 0
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}








