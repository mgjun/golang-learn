package app

type EndingA struct {
	Player  Player
	Monster Monster
}

func NewEndingA(player Player, monster Monster) EndingA {
	return EndingA{player, monster}
}
