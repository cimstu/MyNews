package lobby

import (
	"sync"
	"wolf/room"
)

type Lobby struct {
	Mux*		sync.Mutex
	RoomList 	[]*room.Room
}

func (l *Lobby) FindRoomByPlayer(id string) *room.Room {
	l.Mux.Lock()
	defer l.Mux.Unlock()

	for _, room := range l.RoomList {
		if room.IsPlayerIn(id) {
			return room
		}
	}

	return nil
}