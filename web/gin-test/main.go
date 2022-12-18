package main

import (
	"gin-start/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userApis(r)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func userApis(r *gin.Engine) {
	controller.GetUsers(r, "/users")
	controller.CreateUser(r, "/users")
	controller.ModifyUserId(r, "/users/:id")
	controller.DeleteUserById(r, "/users/:id")
}
