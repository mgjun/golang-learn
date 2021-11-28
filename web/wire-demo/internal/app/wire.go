//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"
	"wire-demo/internal/app/api"
	"wire-demo/internal/app/dao"
	"wire-demo/internal/app/service"
)

var playerAndMonsterSet = wire.NewSet(NewPlayer, NewMonster)

func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}

var endingASet = wire.NewSet(playerAndMonsterSet, wire.Struct(new(EndingA), "*"))
var endingBSet = wire.NewSet(playerAndMonsterSet, wire.Struct(new(EndingB), "*"))

func InitEndingA(name string) EndingA {
	wire.Build(endingASet)
	return EndingA{}
}

func InitEndingB(name string) EndingB {
	wire.Build(endingBSet)
	return EndingB{}
}

// 推荐
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		api.APISet,
		service.SrvSet,
		dao.DaoSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
