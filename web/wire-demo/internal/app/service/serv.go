package service

import "github.com/google/wire"

var SrvSet = wire.NewSet(LoginSrvSet)
