package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/okxBRC20/app/controller/default/controllerDefault"
)

func setRouterDefault(r *gin.Engine) {
	r.GET(
		"/",
		controllerDefault.Home)
	r.GET(
		"/ping",
		controllerDefault.Ping)
	r.GET(
		"/favicon.ico",
		controllerDefault.Favicon)
}
