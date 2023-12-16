package controllerV1

import (
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/BRC20-goSDK/coins/brc20"
	"github.com/kanelinweihao/okxBRC20/app/controller/base/controllerBase"
	"github.com/kanelinweihao/okxBRC20/app/service/serviceInscribe"
	"github.com/kanelinweihao/okxBRC20/app/service/serviceTransaction"
)

type EntityTransaction struct {
	Inputs  []*brc20.TxInput  `json:"inputs"`
	Outputs []*brc20.TxOutput `json:"outputs"`
}

func Inscribe(c *gin.Context) {
	entityInscribe := new(brc20.InscriptionRequest)
	errBindJson := c.ShouldBindJSON(entityInscribe)
	if errBindJson != nil {
		controllerBase.Fail(c, errBindJson)
	}
	jsonFromOKX := serviceInscribe.Inscribe(entityInscribe)
	controllerBase.Back(c, jsonFromOKX)
	return
}

func Transaction(c *gin.Context) {
	entityTransaction := new(EntityTransaction)
	errBindJson := c.ShouldBindJSON(entityTransaction)
	if errBindJson != nil {
		controllerBase.Fail(c, errBindJson)
	}
	jsonFromOKX := serviceTransaction.Transaction(entityTransaction.Inputs, entityTransaction.Outputs)
	controllerBase.Back(c, jsonFromOKX)
	return
}
