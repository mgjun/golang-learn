package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"wire-demo/internal/app/service"
)

var LoginApiSet = wire.NewSet(wire.Struct(new(Login), "*"))

type Login struct {
	LoginSrv service.LoginService
}

func (login *Login) Login(c *gin.Context) {
	login.LoginSrv.Login()
}
