package player

import (
	"fmt"
	"log"
	"mahjong/card"
	"mahjong/game"
	"math/rand"
	"time"
)

func (p *Player) ThrowDefault() int {
	code := p.Cards[len(p.Cards)-1].Code
	p.CardsThrown = append(p.CardsThrown, p.Cards[len(p.Cards)-1])
	p.Cards = p.Cards[:len(p.Cards)-1]

	return code
}

func (p *Player) Throw(index int) {
	p.ThrowOne = true
}


func (p *Player) ThrowDice(count uint8) uint8 {
	if count < 1 || count > 2 {
		log.Fatal("ThrowDice arg overload:", count)
	}

	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(6 * int(count)))
}

func (p *Player) ShowAllCards() {
	fmt.Printf("位置%d：\n", p.Location)

	for _, c := range p.Cards {
		fmt.Println(card.CodeToDesp(c.Code))
	}
}

func (p *Player) Eat(code, index int) {

}

func (p *Player) Action(ch chan game.PlayerAction, code int) {

}