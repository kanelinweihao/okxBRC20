package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/OKXBRC20/app/controller"
)

func setRouterV1(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET(
			"/inscribe",
			controller.Inscribe)
		v1.GET(
			"/transaction",
			controller.Transaction)
	}
}
