package middleware

import (
	"github.com/gin-gonic/gin"
)

func FuncMiddleWareBase() (funcMiddle gin.HandlerFunc) {
	funcMiddle = FuncMiddleWareEmpty()
	return funcMiddle
}
