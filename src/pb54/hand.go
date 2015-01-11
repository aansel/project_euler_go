package pb54


type Hand struct {
	Cards []Card
}

func NewHand(cardsTab []string) *Hand {
	h := new(Hand)
	for _, cardStr := range cardsTab {
		h.Cards = append(h.Cards, *NewCard(cardStr))
	}
	return h
}
