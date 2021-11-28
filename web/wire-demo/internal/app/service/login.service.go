package service

import (
	"github.com/google/wire"
	"wire-demo/internal/app/dao"
)

var LoginSrvSet = wire.NewSet(wire.Struct(new(LoginService), "*"))

type LoginService struct {
	LoginDao dao.LoginDao
}

func (s *LoginService) Login() {
	s.LoginDao.FindByUserName()
}
