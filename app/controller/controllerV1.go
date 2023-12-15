package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/BRC20-goSDK/coins/brc20"
	"github.com/kanelinweihao/okxBRC20/app/service"
)

func Inscribe(c *gin.Context) {
	entityRequestInscribe := new(brc20.InscriptionRequest)
	errBindJson := c.ShouldBindJSON(entityRequestInscribe)
	if errBindJson != nil {
		fail(c, errBindJson)
	}
	jsonFromOKX := service.Inscribe(entityRequestInscribe)
	back(c, jsonFromOKX)
	return
}

func Transaction(c *gin.Context) {
	name := c.DefaultQuery("name", "Mr.Jin")
	data := fmt.Sprintf("hello %s", name)
	success(c, data)
	return
}
