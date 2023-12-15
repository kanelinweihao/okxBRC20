package middleware

import (
	"github.com/gin-gonic/gin"
)

func FuncMiddleWareEmpty() (funcMiddle gin.HandlerFunc) {
	funcMiddle = func(c *gin.Context) {
		c.Next()
	}
	return funcMiddle
}
