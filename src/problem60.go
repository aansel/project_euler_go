package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	getSmallestSumGroup(5, 10000)
}

type PrimeGroup struct {
	Primes []int
	Sum int
}

func NewPrimeGroup(primes []int) *PrimeGroup {
	p := new(PrimeGroup)
	for _, prime := range primes {
		p.Primes = append(p.Primes, prime)
		p.Sum += prime
	}
	return p
}

func (pg PrimeGroup) Size() int {
	return len(pg.Primes)
}

func (pg PrimeGroup) CanJoin(p int) bool {
	for _, primeInGroup := range pg.Primes {
		if !isPrimeConcat(p, primeInGroup) {
			return false
		}
	}
	return true
}

func getSmallestSumGroup(size int, limit int) {
	smallestSum := 0
	var smallestGroup *PrimeGroup
	for _, groups := range getPrimeGroups(limit) {
		for _, group := range groups {
//			fmt.Println(p, "a le groupe", group.Primes)
			if group.Size() == size {
				if smallestSum == 0 || group.Sum < smallestSum {
					smallestSum = group.Sum
					smallestGroup = group
				}
			}
		}
	}

	if smallestGroup != nil {
		for _, p := range smallestGroup.Primes {
			fmt.Println(p)
		}
		fmt.Println("Sum", smallestSum)
	} else {
		fmt.Println("Not found")
	}
}


func getPrimeGroups(limit int) map[int][]*PrimeGroup {
	m := make(map[int][]*PrimeGroup)
	for i := 3; i < limit; i += 2 {
		if isPrime(i) {
			for prime, groups := range m {
				for _, group := range groups {
					if group.CanJoin(i) {
						pg := NewPrimeGroup(append(group.Primes, i))
						m[prime] = append(m[prime], pg)
					}
				}
			}
			m[i] = []*PrimeGroup{NewPrimeGroup([]int{i})}
		}
	}
	return m
}


func isPrimeConcat(p1 int, p2 int) bool {
	pstr1 := strconv.Itoa(p1) + strconv.Itoa(p2)
	if pp1, _ := strconv.Atoi(pstr1); !isPrime(pp1) {
		return false
	}

	pstr2 := strconv.Itoa(p2) + strconv.Itoa(p1)
	if pp2, _ := strconv.Atoi(pstr2); !isPrime(pp2) {
		return false
	}

	return true
}

func isHPrime(nb int) bool {
	if nb == 13 || nb == 5197 || nb == 5701 || nb == 6733  || nb == 8389 {
		return true
	} else {
		return false
	}
}

var primesMap map[int]int = make(map[int]int)
func isPrime(nb int) bool {
	r :=  primesMap[nb]
	if r == 1 {
		return true
	} else if r == -1 {
		return false
	} else {
		if nb < 2 {
			return false
		}
		var i int
		for i = 2; i <= int(math.Sqrt(float64(nb))); i++ {
			if nb % i == 0 {
				primesMap[nb] = -1
				return false
			}
		}
		primesMap[nb] = 1
		return true
	}
}
