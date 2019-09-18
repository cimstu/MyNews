package logic

import "wolf/player"

func GetPlayerInfo(p* player.Player, role string) player.Player {
	player := *p
	if role != ROLE_JUDGER && ( role == p.Role && p.Role == ROLE_WOLF) {
		player.Role = ROLE_NONE
	}

	return player
}