package serviceInscribe

import (
	"encoding/json"
	"github.com/kanelinweihao/BRC20-goSDK/coins/brc20"
	"github.com/kanelinweihao/okxBRC20/app/service/serviceBase"
	"github.com/kanelinweihao/okxBRC20/app/utils/err"
	"log"
)

func Inscribe(entityInscribeRequest *brc20.InscriptionRequest) (jsonFromOKX string) {
	request := getRequest(entityInscribeRequest)
	network := serviceBase.GetNetwork()
	txs, errInscribe := brc20.Inscribe(network, request)
	err.ErrCheck(errInscribe)
	jsonFromOKX = getDataFromOKX(txs)
	return jsonFromOKX
}

func getRequest(entityRequestInscribe *brc20.InscriptionRequest) (request *brc20.InscriptionRequest) {
	request = entityRequestInscribe
	bytesRequest, errMarshal := json.Marshal(request)
	err.ErrCheck(errMarshal)
	strRequest := string(bytesRequest)
	log.Println(strRequest)
	return request
}

func getDataFromOKX(txs *brc20.InscribeTxs) (jsonFromOKX string) {
	bytesTxs, errMarshal := json.Marshal(txs)
	err.ErrCheck(errMarshal)
	strTxs := string(bytesTxs)
	log.Println(strTxs)
	jsonFromOKX = strTxs
	return jsonFromOKX
}
