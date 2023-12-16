package serviceTransaction

import (
	"github.com/kanelinweihao/BRC20-goSDK/coins/brc20"
	"github.com/kanelinweihao/okxBRC20/app/service/serviceBase"
	"github.com/kanelinweihao/okxBRC20/app/utils/err"
	"log"
)

func Transaction(inputs []*brc20.TxInput, outputs []*brc20.TxOutput) (jsonFromOKX string) {
	network := serviceBase.GetNetwork()
	signedTx, errTransaction := brc20.Transfer(inputs, outputs, network)
	err.ErrCheck(errTransaction)
	jsonFromOKX = getDataFromOKX(signedTx)
	return jsonFromOKX
}

func getDataFromOKX(signedTx string) (jsonFromOKX string) {
	log.Println(signedTx)
	jsonFromOKX = signedTx
	return jsonFromOKX
}
