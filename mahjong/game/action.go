package game

import (
	"log"
	"mahjong/card"
	"mahjong/player"
)

type PlayerAction struct {
	Location int
	Action   int
	Cards    card.CardSlice
}

//func (g *Game) OnActionEat(player *player.Player, cards card.CardSlice) bool {
//	if len(cards) != 3 {
//		log.Fatal("OnActionEat cards count not 3", cards.Len())
//		return false
//	}
//
//	if cards[0] == cards[1] || cards[0] == cards[2] || cards[0] == cards[2] {
//		log.Fatal("OnActionEat cards not 3 diffrend cards:", cards[0], cards[1], cards[2])
//		return false
//	}
//
//	count := 0
//	for _, card := range player.Cards {
//		if g.Code == card.Code {
//			count++
//		}
//
//		if count >= 3 {
//			player.OnShowCards(cards)
//			return true
//		}
//	}
//
//	return false
//}
func (g *Game) OnActionEat(player *player.Player, cards card.CardSlice) bool {
	if len(cards) != 3 {
		log.Fatal("OnActionEat cards count not 3", cards.Len())
		return false
	}

	if cards[0] == cards[1] || cards[0] == cards[2] || cards[0] == cards[2] {
		log.Fatal("OnActionEat cards not 3 diffrend cards:", cards[0], cards[1], cards[2])
		return false
	}

	count := 0
	for _, card := range player.Cards {
		if g.Code == card.Code {
			count++
		}

		if count >= 3 {
			player.OnShowCards(cards)
			return true
		}
	}

	return false
}


//func (g *Game) OnActionThree(player *player.Player) bool {
//	cards := card.CardSlice{}
//	cards = append(cards, &card.Card{Code: g.Code})
//
//	count := 0
//	for _, card := range player.Cards {
//		if g.Code == card.Code {
//			count++
//			cards = append(cards, card)
//		}
//
//		if count >= 2 {
//			player.OnShowCards(cards)
//			return true
//		}
//	}
//
//	return false
//}

func (g *Game) OnActionFour(player *player.Player) bool {
	cards := card.CardSlice{}
	cards = append(cards, &card.Card{Code: g.Code})

	count := 0
	for _, card := range player.Cards {
		if g.Code == card.Code {
			count++
		}

		if count >= 3 {
			player.OnShowCards(cards)
			return true
		}
	}

	return false
}

