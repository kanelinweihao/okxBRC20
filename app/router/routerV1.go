package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/okxBRC20/app/controller/v1/controllerV1"
)

func setRouterV1(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST(
			"/inscribe",
			controllerV1.Inscribe)
		v1.POST(
			"/transaction",
			controllerV1.Transaction)
	}
}
