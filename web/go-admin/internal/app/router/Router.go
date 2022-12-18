package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-admin/pkg/auth"
)

var _ IRouter = (*Router)(nil)

var IRouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	Auth auth.Auth
}

func (r *Router) Register(app *gin.Engine) error {
	return nil
}

func (r *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

func (r *Router) RegisterAPI(app *gin.Engine) {
}
