package main

import (
	"gin-start/handler"
	"gin-start/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	middleware.RegisterMiddleware(r)
	handler.RegisterHandler(r)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
