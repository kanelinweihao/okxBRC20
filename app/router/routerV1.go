package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/okxBRC20/app/controller"
)

func setRouterV1(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST(
			"/inscribe",
			controller.Inscribe)
		v1.POST(
			"/transaction",
			controller.Transaction)
	}
}
