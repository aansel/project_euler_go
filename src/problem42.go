package main

import "fmt"
import "math"
import "net/http"
import (
	"io/ioutil"
	"strings"
)

func main() {
	//fmt.Println(isTriangleNumber(getWordValue("SKY")))
	resp, _ := http.Get("https://projecteuler.net/project/resources/p042_words.txt")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	nb := 0
	for _, word := range strings.Split(string(body), ",") {
		w := strings.Replace(word, "\"", "", -1)
		if isTriangleWord(w) {
			fmt.Println(w)
			nb++
		}
	}
	fmt.Println(nb, "triangle words")
}

func isTriangleWord(word string) bool {
	return isTriangleNumber(getWordValue(word))
}

func getWordValue(word string) int {
	sum := 0
	for i := 0; i < len(word); i++ {
		sum += int(word[i] - "A"[0] + 1)
	}
	return sum
}

func isTriangleNumber(N int) bool {
	f := (-1 + math.Sqrt(float64(1 + 8 * N))) / 2
	n := int(f)
	return n * (n + 1) / 2 == N
}
