package dao

import (
	"fmt"
	"github.com/google/wire"
)

var LoginDaoSet = wire.NewSet(wire.Struct(new(LoginDao), "*"))

type LoginDao struct {
}

func (d *LoginDao) FindByUserName() {
	fmt.Println("login is required")
}
