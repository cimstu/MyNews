package player

import (
	"mahjong/card"
	"sort"
)

func (p *Player) CanEat(code, location int) []IndexSlice {
	p.EatPairs = p.EatPairs[0:0]

	if location+1 != p.Location {
		return nil
	}

	card := card.Card{Code: code}
	//card on left
	{
		card1 := p.findCard(p.Cards, code+1, 0)
		if card1 != nil && card1.Type() == card.Type() {
			card2 := p.findCard(p.Cards, code+2, card1.Index+1)
			if card2 != nil && card2.Type() == card.Type() {
				cards := IndexSlice{card1.Index, card2.Index}
				p.EatPairs = append(p.EatPairs, cards)
			}
		}

	}

	//card on right
	{
		card1 := p.findCard(p.Cards, code-1, 0)
		if card1 != nil && card1.Type() == card.Type() {
			card2 := p.findCard(p.Cards, code-2, card1.Index+1)
			if card2 != nil && card2.Type() == card.Type() {
				cards := IndexSlice{card1.Index, card2.Index}
				p.EatPairs = append(p.EatPairs, cards)
			}
		}
	}

	//card in middle
	{
		card1 := p.findCard(p.Cards, code-1, 0)
		if card1 != nil && card1.Type() == card.Type() {
			card2 := p.findCard(p.Cards, code+1, card1.Index+1)
			if card2 != nil && card2.Type() == card.Type() {
				cards := IndexSlice{card1.Index, card2.Index}
				p.EatPairs = append(p.EatPairs, cards)
			}
		}
	}

	return p.EatPairs
}

func (p *Player) CanThree(code, location int) IndexSlice {
	p.ThreePair = p.ThreePair[0:0]

	card := card.Card{Code: code}

	card1 := p.findCard(p.Cards, code, 0)
	if card1 != nil && card1.Type() == card.Type() {
		card2 := p.findCard(p.Cards, code, card1.Index+1)
		if card2 != nil && card2.Type() == card.Type() {
			p.ThreePair = IndexSlice{card1.Index, card2.Index}
		}
	}

	return p.ThreePair
}

func (p *Player) CanGangInHand(code, location int) []IndexSlice {
	p.GangSlices = p.GangSlices[0:0]

	card := card.Card{Code: code}

	card1 := p.findCard(p.Cards, code, 0)
	if card1 != nil && card1.Type() == card.Type() {
		card2 := p.findCard(p.Cards, code, card1.Index+1)
		if card2 != nil && card2.Type() == card.Type() {
			card3 := p.findCard(p.Cards, code, card2.Index+1)
			if card3 != nil && card3.Type() == card.Type() {
				p.ThreePair = IndexSlice{card1.Index, card2.Index, card3.Index}
				p.GangSlices = append(p.GangSlices, p.ThreePair)
			}
		}
	}

	return p.GangSlices
}

func (p *Player) CanGangInShown(code int) int {
	for i, pair := range p.CardsPeng {
		if pair[0].Code == code {
			return i
		}
	}
	return -1
}

func (p *Player) CanEggInShown(code, location int) int {
	t := card.GetEggTypeByCode(code)
	if t != card.EGG_TYPE_NONE {
		return -1
	}

	for i, egg := range p.CardsEgg {
		egg_type := card.GetEggType(egg)
		if t&egg_type > 0 {
			return i
		}
	}

	return -1
}

func (p *Player) CanEggInHand(code, location int) []IndexSlice {
	t := card.GetEggTypeByCode(code)
	if t != card.EGG_TYPE_NONE {
		return nil
	}

	eggs := []IndexSlice{}

	for i, first_card := range p.Cards {
		if card.GetEggTypeByCode(first_card.Code)|t > 0 {
			for j := i + 1; j < p.Cards.Len(); j++ {
				second_card := p.Cards[j]
				if (card.GetEggTypeByCode(second_card.Code)|t > 0) && (first_card.Code != second_card.Code) {
					eggs = append(eggs, IndexSlice{first_card.Index, second_card.Index})
				}
			}
		}
	}

	return eggs
}

func (p *Player) CheckEggInHand() (eggs []card.CardSlice) {
	isRepeat := func(slice card.CardSlice, code int) bool {
		for _, c := range slice {
			if c.Code == code {
				return true
			}
		}

		return false
	}

	norepeat_slice := card.CardSlice{}
	for _, c := range p.Cards {
		if c.Code == card.CARD_CODE_TIAO_BEGIN || !isRepeat(norepeat_slice, c.Code) {
			norepeat_slice = append(norepeat_slice, c)
		}
	}

	tmp_eggs := []card.CardSlice{}
	combin_slice := card.Combine3Cards(norepeat_slice)
	for _, cards := range combin_slice {
		if card.IsEgg(cards) {
			tmp_eggs = append(tmp_eggs, cards)
		}
	}

	isRepeatEgg := func(slice []card.CardSlice, cards card.CardSlice) bool {
		sort.Sort(card.CardSlice(cards))
		for _, egg := range slice {
			sort.Sort(card.CardSlice(egg))
			if egg[0]==cards[0] && egg[1]==cards[1] && egg[2]==cards[2] {
				return true
			}
		}

		return false
	}

	for _, egg := range tmp_eggs {
		if !isRepeatEgg(eggs, egg) {
			eggs = append(eggs, egg)
		}
	}

	return eggs
}

func (p *Player) CanWin(code int) bool {
	return false
}