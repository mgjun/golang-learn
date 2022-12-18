package middleware

import (
	"gin-start/middleware/debug"
	"github.com/gin-gonic/gin"
)

func RegisterMiddleware(r *gin.Engine) {
	r.Use(debug.MiddlewareDebug())
}
