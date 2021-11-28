package api

import "github.com/google/wire"

// APISet repo injection
var APISet = wire.NewSet(
	LoginApiSet,
)
