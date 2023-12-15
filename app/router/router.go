package router

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	setRouterDefault(r)
	setRouterV1(r)
	return
}
