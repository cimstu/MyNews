package player

import (
	"fmt"
	"mahjong/card"
	"mahjong/game"
)

func (p *Player) OnGotCards(cards card.CardSlice) {
	for _, c := range cards {
		p.OnGotCard(c)
	}
}

func (p *Player) OnGotCard(c *card.Card) *card.Card {
	switch p.Location {
	case game.PLAY_LOCATION_EAST:
		c.State = card.CARD_STATE_IN_HAND_EAST
	case game.PLAY_LOCATION_SORTH:
		c.State = card.CARD_STATE_IN_HAND_SORTH
	case game.PLAY_LOCATION_WEST:
		c.State = card.CARD_STATE_IN_HAND_WEST
	case game.PLAY_LOCATION_NORTH:
		c.State = card.CARD_STATE_IN_HAND_NORTH
	}
	p.Cards = append(p.Cards, c)

	p.sortCards()
	return c
}

func (p *Player) OnShowCards(cards card.CardSlice, tp int) {
	for _, toshow := range cards {
		for j, card := range p.Cards {
			if card.Code == toshow.Code {
				p.Cards = append(p.Cards[:j-1], p.Cards[j:]...)
				break
			}
		}
	}

	switch tp {
	case game.SHOWN_CARDS_TYPE_EAT:
		p.CardsEat = append(p.CardsEat, cards)
	case game.SHOWN_CARDS_TYPE_PENG:
		p.CardsPeng = append(p.CardsPeng, cards)
	case game.SHOWN_CARDS_TYPE_GANG:
		p.CardsGang = append(p.CardsGang, cards)
	case game.SHOWN_CARDS_TYPE_EGG:
		p.CardsGang = append(p.CardsGang, cards)
	}

	p.sortCards()
}

func (p *Player) OnGameThrownCard(code, location int) {
	if p.Location == location {
		return
	}

	if p.CanEat(code, location) != nil {
		fmt.Println("Can eat:", card.CodeToDesp(code))
	}

	if p.CanThree(code, location) != nil {
		fmt.Println("Can three:", card.CodeToDesp(code))
	}

	if p.CanGangInHand(code, location) != nil {
		fmt.Println("Can gang:", card.CodeToDesp(code))
	}
}

func (p *Player) OnCheck() {
	if !p.ThrowOne {
		if(p.CanEggInHand()) {
			fmt.Println("")
		}
	}

	if p.CanThree(code, p.Location) != nil {
		fmt.Println("Can three:", card.CodeToDesp(code))
	}

	if p.CanGangInHand(code, p.Location) != nil {
		fmt.Println("Can gang:", card.CodeToDesp(code))
	}
}
