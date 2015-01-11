package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	resp, _ := http.Get("https://projecteuler.net/project/resources/p059_cipher.txt")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	cipheredCar := strings.Split(string(body), ",")
	cipheredCar[1200] = "73"
//	for i := 97; i <= 122; i++ {
//		for j := 97; j <= 122; j++ {
//			for k := 97; k <= 122; k++ {
//				deciphered := decipher(cipheredCar, i, j, k)
//				if !strings.Contains(deciphered, "{") && !strings.Contains(deciphered, "}") && strings.Contains(deciphered, ".") {
//					fmt.Println(i, j, k, deciphered)
//				}
//			}
//		}
//	}

	deciphered, sum := decipher(cipheredCar, 103, 111, 100)
	fmt.Println(deciphered)
	fmt.Println(sum)
}


func decipher(value []string, car1Key int, car2Key int, car3Key int) (string, int) {
	res := ""
	sum := 0
	for idx, cipherStr :=  range value {
		cipher, _ := strconv.Atoi(cipherStr)
		var key int
		if idx%3 == 0 {
			key = car1Key
		} else if idx%3 == 1 {
			key = car2Key
		} else if idx%3 == 2 {
			key = car3Key
		}
		deciphered := cipher ^ key
		sum += deciphered
		res += string(deciphered)
	}
	return res, sum
}
