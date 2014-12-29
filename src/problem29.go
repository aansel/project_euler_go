package main

import (
	"fmt"
	"math/big"
)


func main() {
	//fmt.Println(get(1000).Size())
	set := NewIntSet()
	set.Add(big.NewInt(1))
	fmt.Println(pow(2, 3))
}

//func get(limit int64) *IntSet {
//	set := NewIntSet()
//	var a, b int64 = 2, 2
//	for ; a <= limit; a++ {
//		for ; b <= limit; b++ {
//			set.Add(pow(a, b))
//		}
//	}
//	return set
//}

func pow(a int64, b int64) *big.Int {
	var p = big.NewInt(1)
	var i int64 = 0
	for ; i < b; i++ {
		p = p.Mul(p, big.NewInt(a))
	}
	return p
}



type IntSet struct {
	set map[big.Int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[big.Int]bool)}
}

func (set *IntSet) Add(i *big.Int) bool {
	fmt.Println(i)
	_, found := set.set[i]
	set.set[i] = true
	return !found	//False if it existed already
}

func (set *IntSet) Contains(i *big.Int) bool {
	_, found := set.set[i]
	return found	//true if it existed already
}

func (set *IntSet) Remove(i *big.Int) {
	delete(set.set, i)
}

func (set *IntSet) Size() big.Int {
	return len(set.set)
}
