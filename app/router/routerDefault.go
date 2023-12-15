package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/okxBRC20/app/controller"
)

func setRouterDefault(r *gin.Engine) {
	r.GET(
		"/",
		controller.Home)
	r.GET(
		"/ping",
		controller.Ping)
	r.GET(
		"/favicon.ico",
		controller.Favicon)
}
