package app

import "github.com/google/wire"

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InjectorSet,
	)
	return new(Injector), nil, nil
}
