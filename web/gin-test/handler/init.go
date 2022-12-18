package handler

import (
	"gin-start/handler/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(r *gin.Engine) {
	userApis(r)
}

func userApis(r *gin.Engine) {
	r.Handle(http.MethodGet, "/users", user.GetUsers())
	r.Handle(http.MethodPost, "/users", user.CreateUser())
	r.Handle(http.MethodPut, "/users/:id", user.ModifyUserId())
	r.Handle(http.MethodDelete, "/users", user.DeleteUserById())
}
