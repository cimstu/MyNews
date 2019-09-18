package room

import (
	"sync"
	"wolf/logic"
	"wolf/player"
)

type Room struct {
	Mux*	sync.Mutex

	PlayerList []player.Player
}

func(r* Room)IsPlayerIn(id string) bool {
	r.Mux.Lock()
	defer r.Mux.Unlock()

	for _, player := range r.PlayerList  {
		if player.ID == id {
			return true
		}
	}

	return false
}

func(r *Room)GetRoomInfo(role string) {
	players := []player.Player{}

	r.Mux.Lock()
	defer r.Mux.Unlock()

	for _, p := range r.PlayerList  {
		players = append(players, logic.GetPlayerInfo(&p, role))
	}
}