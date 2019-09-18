package player

import (
	"mahjong/card"
	"sort"
)

type IndexSlice []int

type Player struct {
	Cards       card.CardSlice //手中未展示牌
	CardsThrown card.CardSlice

	CardsEat  []card.CardSlice
	CardsPeng []card.CardSlice
	CardsGang []card.CardSlice
	CardsEgg  []card.CardSlice

	State    int

	Location int

	EatPairs   []IndexSlice
	ThreePair  IndexSlice
	GangSlices []IndexSlice
	ToEgg      []IndexSlice

	ThrowOne bool
}

func (p *Player) sortCards() {
	sort.Sort(card.CardSlice(p.Cards))
	for i, card := range p.Cards {
		card.Index = i
	}
}


func (p *Player) findCard(cards card.CardSlice, code, index_beg int) *card.Card {
	for i := index_beg; i < cards.Len(); i++ {
		if code == cards[i].Code {
			return cards[i]
		}
	}

	return nil
}

func (p *Player) GetCardsByType(slice card.CardSlice, t int) card.CardSlice {
	cards := card.CardSlice{}
	for _, card := range slice {
		if card.Type() == t {
			cards = append(cards, card)
		}
	}

	return cards
}

