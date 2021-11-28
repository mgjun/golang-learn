package app

type EndingB struct {
	Player  Player
	Monster Monster
}

func NewEndingB(player Player, monster Monster) EndingB {
	return EndingB{player, monster}
}
