package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func FuncMiddleWareTimeCost() (funcMiddle gin.HandlerFunc) {
	funcMiddle = func(c *gin.Context) {
		timeBegin := time.Now()
		fmt.Println("timeBegin:", timeBegin)
		c.Next()
		code := c.Writer.Status()
		fmt.Println("code:", code)
		timeCost := time.Since(timeBegin)
		fmt.Println("timeCost:", timeCost)
		timeEnd := time.Now()
		fmt.Println("timeEnd:", timeEnd)
	}
	return funcMiddle
}
