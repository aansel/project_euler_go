package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"strings"
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


type Hand struct {
	Cards []Card
	cardsOcc [15][]Card
}

func NewHand(cardsTab []string) *Hand {
	h := new(Hand)
	for _, cardStr := range cardsTab {
		h.Cards = append(h.Cards, *NewCard(cardStr))
	}
	for _, card := range h.Cards {
		h.cardsOcc[card.Value] = append(h.cardsOcc[card.Value], card)
	}
	return h
}

func (h Hand) hasNothing() int {
	card1 := 0
	card2 := 0
	card3 := 0
	card4 := 0
	card5 := 0
	for value, cardOcc := range h.cardsOcc {
		if len(cardOcc) == 1 {
			if value > card1 {
				card5 = card4
				card4 = card3
				card3 = card2
				card2 = card1
				card1 = value
			} else if value > card2 {
				card5 = card4
				card4 = card3
				card3 = card2
				card2 = value
			} else if value > card3 {
				card5 = card4
				card4 = card3
				card3 = value
			} else if value > card4 {
				card5 = card4
				card4 = value
			} else if value > card5 {
				card5 = value
			}
		}
	}
	return 100000000 * card1 + 1000000 * card2 + 10000 * card3 + 100 * card4 + card5
}

func (h Hand) hasPair() int {
	cardNotPair1 := 0
	cardNotPair2 := 0
	cardNotPair3 := 0
	bestPair := 0
	for value, cardOcc := range h.cardsOcc {
		if len(cardOcc) == 2 {
			bestPair = value
		} else if len(cardOcc) == 1 {
			if value > cardNotPair1 {
				cardNotPair3 = cardNotPair2
				cardNotPair2 = cardNotPair1
				cardNotPair1 = value
			} else if value > cardNotPair2 {
				cardNotPair3 = cardNotPair2
				cardNotPair2 = value
			} else if value > cardNotPair3 {
				cardNotPair3 = value
			}
		}
	}
	if bestPair > 0 {
		return 10000000000 + 100000000 * bestPair + 1000000 * cardNotPair1 + 10000 * cardNotPair2 + 100 * cardNotPair3
	} else {
		return 0
	}
}

func (h Hand) hasTwoPairs() int {
	bestPair := 0
	bestSecondPair := 0
	lastCard := 0
	for value, cardOcc := range h.cardsOcc {
		if len(cardOcc) == 2 {
			bestSecondPair = bestPair
			bestPair = value
		} else if len(cardOcc) == 1 {
			lastCard = value
		}
	}
	if bestPair > 0 && bestSecondPair > 0 {
		return 20000000000 + bestPair * 100000000 + bestSecondPair * 1000000 + lastCard
	} else {
		return 0
	}
}

func (h Hand) hasThreeOfAKind() int {
	for value, cardOcc := range h.cardsOcc {
		if len(cardOcc) == 3 {
			return 30000000000 + value
		}
	}
	return 0
}

func (h Hand) hasStraight() int {
	lastOfStraight := 0
	nbStraight := 0
	for value, cardOcc := range h.cardsOcc {
		if len(cardOcc) >= 1 {
			if lastOfStraight == 0 || lastOfStraight + 1 == value {
				nbStraight++
				lastOfStraight = value
			}
		}
	}
	if nbStraight == 5 {
		return 40000000000 +  lastOfStraight
	} else {
		return 0
	}
}

func (h Hand) hasFlush() int {
	color := ""
	for _, card := range h.Cards {
		if color == "" {
			color = card.Color
		} else if color != card.Color {
			return 0
		}
	}
	return 50000000000 + h.hasNothing()
}

func (h Hand) hasFullHouse() int {
	scoreThreeKind := h.hasThreeOfAKind()
	scorePair := h.hasPair()
	if scoreThreeKind > 0 && scorePair > 0 {
		return 60000000000 + (scoreThreeKind - 30000000000) * 100000000 + (scorePair - 10000000000) / 100000000
	} else {
		return 0
	}
}

func (h Hand) hasFourOfAKind() int {
	for value, cardOcc := range h.cardsOcc {
		if len(cardOcc) == 4 {
			return 70000000000 + value
		}
	}
	return 0
}

func (h Hand) hasStraightFlush() int {
	straight := h.hasStraight()
	flush := h.hasFlush()
	if straight > 0 && flush > 0 {
		return straight + 40000000000
	} else {
		return 0
	}
}


func (h Hand) Score() int {
	score := 0
	if lScore := h.hasStraightFlush(); lScore > 0 {
		score = lScore
	} else if lScore := h.hasFourOfAKind(); lScore > 0 {
		score = lScore
	} else if lScore := h.hasFullHouse(); lScore > 0 {
		score = lScore
	} else if lScore := h.hasFlush(); lScore > 0 {
		score = lScore
	} else if lScore := h.hasStraight(); lScore > 0 {
		score = lScore
	} else if lScore := h.hasThreeOfAKind(); lScore > 0 {
		score = lScore
	} else if lScore := h.hasTwoPairs(); lScore > 0 {
		score = lScore
	} else if lScore := h.hasPair(); lScore > 0 {
		score = lScore
	} else if lScore := h.hasNothing(); lScore > 0 {
		score = lScore
	}
	return score
}



func main() {
	resp, _ := http.Get("https://projecteuler.net/project/resources/p054_poker.txt")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	lines := strings.Split(string(body), "\n")

	nbVictoriesPlayer1 := 0
	for _, line := range lines {
		cards := strings.Split(line, " ")
		if len(cards) == 10 {
			hand1 := NewHand(cards[:5])
			hand2 := NewHand(cards[5:])
			if hand1.Score() > hand2.Score() {
				nbVictoriesPlayer1++
			}
		}
	}
	fmt.Println(nbVictoriesPlayer1)
}


