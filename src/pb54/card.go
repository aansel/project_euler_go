package pb54

import (
	"strconv"
)

type Card struct {
	Color string
	Value int
}

func NewCard(card string) *Card {
	c := new(Card)
	val := string(card[0])
	if val == "A" {
		c.Value = 14
	} else if val == "K" {
		c.Value = 13
	} else if val == "Q" {
		c.Value = 12
	} else if val == "J" {
		c.Value = 11
	}  else if val == "T" {
		c.Value = 10
	} else {
		v, _ := strconv.Atoi(val)
		c.Value = v
	}

	c.Color = string(card[1])

	return c
}
