package card

type Card struct {
	Index int
	Code  int
	State int
}

func (c *Card) Type() int {
	return CardType(c.Code)
}
