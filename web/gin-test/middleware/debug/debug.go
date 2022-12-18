package debug

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MiddlewareDebug() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		fmt.Println("[Debug Middleware] - request: \n", req)
		c.Next()
		fmt.Println("[Debug Middleware] - response: \n", c.Writer.Status())
	}
}
