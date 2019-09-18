package table

import (
	"mahjong/card"
	"mahjong/game"
	"math/rand"
	"time"
)

type Table struct {
	Cards    CardSlice
	StartPos int
	CatchPos int
}

func (t *Table) Init() {

}

func (t *Table) Create() {
	for n := 0; n < 4; n++ {
		rand.Seed(time.Now().UnixNano())
		for i := n * MAX_CARDS_CODE; i < MAX_CARDS_CODE*(n+1); i++ {
			t.Cards = append(t.Cards, &Card{Code: rand.Intn(MAX_CARDS_CODE - 1), State: CARD_STATE_ON_TABLE_UNSHOWN})
		}
	}
}

func (t *Table) MoveCatchPos(step int) {
	if t.CatchPos+step <= MAX_CARDS_INDEX {
		t.CatchPos += step
	} else {
		t.CatchPos = t.CatchPos + step - MAX_CARDS_INDEX - 1
	}
}

func (t *Table) SelectOne(pos int) *card.Card {
	if (pos >= t.CatchPos && pos <= game.MAX_CARDS_INDEX) || pos < t.StartPos {
		card := t.Cards[pos]

		if card.State != CARD_STATE_ON_TABLE_UNSHOWN {
			return nil
		}

		card.State = CARD_STATE_ON_TABLE_SENT
		return card
	}

	return nil
}

func (t *Table) CatchOne() *Card {
	card := t.SelectOne(t.CatchPos)
	t.MoveCatchPos(1)

	for card == nil {
		t.MoveCatchPos(1)
		card = t.SelectOne(t.CatchPos)
	}

	return card
}

func (t *Table) CatchSome(count int) CardSlice {
	cards := CardSlice{}
	for i := 0; i < count; i++ {
		cards = append(cards, t.CatchOne())
	}

	return cards
}
