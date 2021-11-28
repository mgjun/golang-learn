package app

import (
	"fmt"
)

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(player Player, monster Monster) Mission {
	return Mission{Player: player, Monster: monster}
}

func (m Mission) start() {
	fmt.Printf("%s defeats %s, world peace!", m.Player.Name, m.Monster.Name)
}
