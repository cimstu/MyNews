package game

import (
	"mahjong/player"
	"mahjong/table"
	"time"
)

const (
	MAX_CARDS_INDEX = 135
	MAX_CARDS_CODE  = 34

	TIMER_ESCAPE = 10 * time.Second
)

type Game struct {
	Table   *table.Table
	Players [4]*player.Player
	Dealer  int //庄家index

	Timer    *time.Timer
	Location int
	Code     int
}

func (g *Game) init() {
	g.Table = &table.Table{}
	g.Table.Create()

	g.Players[0] = &player.Player{Location: PLAY_LOCATION_EAST}
	g.Players[1] = &player.Player{Location: PLAY_LOCATION_SORTH}
	g.Players[2] = &player.Player{Location: PLAY_LOCATION_WEST}
	g.Players[3] = &player.Player{Location: PLAY_LOCATION_NORTH}

	g.Dealer = PLAY_LOCATION_EAST
	g.Location = PLAY_LOCATION_EAST

	g.Timer = time.AfterFunc(TIMER_ESCAPE, g.onTimer)
	g.Timer.Stop()
}

func (g *Game) Run() {
	g.init()

	//1
	dealer := g.Players[g.Dealer]

	dice_points := dealer.ThrowDice(2)
	if dice_points%4 == 0 {
		dice_points += g.Players[3].ThrowDice(2)
	} else {
		dice_points += g.Players[dice_points%4-1].ThrowDice(2)
	}

	//2
	g.Table.StartPos = int(dice_points-1) * 2

	for i := 0; i < 3; i++ {
		g.Players[g.Dealer].OnGotCards(g.Table.CatchSome(4))
		g.Players[(g.Dealer+1)%4].OnGotCards(g.Table.CatchSome(4))
		g.Players[(g.Dealer+2)%4].OnGotCards(g.Table.CatchSome(4))
		g.Players[(g.Dealer+3)%4].OnGotCards(g.Table.CatchSome(4))
	}

	g.Players[g.Dealer].OnGotCard(g.Table.CatchOne())
	g.Code = g.Players[g.Dealer].OnGotCard(g.Table.SelectOne(g.Table.CatchPos + 3)).Code
	g.Players[(g.Dealer+1)%4].OnGotCard(g.Table.CatchOne())
	g.Players[(g.Dealer+2)%4].OnGotCard(g.Table.CatchOne())
	g.Players[(g.Dealer+3)%4].OnGotCard(g.Table.CatchOne())

	//3
	g.Players[g.Dealer].State = player.PLAYER_STATE_THROWING
	g.waitThrower()
}

func (g *Game) onTimer() {
	g.Code = g.Players[g.Location].ThrowDefault()
}

func (g *Game) waitThrower(players ...player.Player) {
	g.Timer.Reset(TIMER_ESCAPE)

	ch := make(chan PlayerAction, 1)
	g.Players[g.Dealer].Action(ch, g.Code)

	action :<- ch
}

func (g *Game) NotifyThrow(code int, location int) {
	for i := PLAY_LOCATION_EAST; i <= PLAY_LOCATION_NORTH; i++ {
		g.Players[i].OnGameThrownCard(code, location)
	}
}

func (g *Game) NotifyWait(code int, location int) {
	for i := PLAY_LOCATION_EAST; i <= PLAY_LOCATION_NORTH; i++ {
		if g.Players[i].State == player.PLAYER_STATE_THROWING {
			g.Players[i].OnWait()
		} else {
			g.Players[i].OnCheck()
		}
	}
}
