package service

import (
	"encoding/json"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/kanelinweihao/BRC20-goSDK/coins/brc20"
	"github.com/kanelinweihao/okxBRC20/app/utils/err"
	"log"
)

func Inscribe(entityRequestInscribe *brc20.InscriptionRequest) (jsonFromOKX string) {
	network := getNetwork()
	// requestExample := getRequestExample()
	request := getRequest(entityRequestInscribe)
	txs, errInscribe := brc20.Inscribe(network, request)
	err.ErrCheck(errInscribe)
	jsonFromOKX = getDataFromOKX(txs)
	return jsonFromOKX
}

func getNetwork() (network *chaincfg.Params) {
	network = &chaincfg.TestNet3Params
	return network
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

/*
Test
*/

func getRequestExample() (requestExample *brc20.InscriptionRequest) {
	commitTxPrevOutputList := getCommitTxPrevOutputList()
	inscriptionDataList := getinscriptionDataList()
	requestExample = &brc20.InscriptionRequest{
		CommitTxPrevOutputList: commitTxPrevOutputList,
		CommitFeeRate:          2,
		RevealFeeRate:          2,
		RevealOutValue:         546,
		InscriptionDataList:    inscriptionDataList,
		ChangeAddress:          "2NF33rckfiQTiE5Guk5ufUdwms8PgmtnEdc",
	}
	return requestExample
}

func getCommitTxPrevOutputList() (commitTxPrevOutputList []*brc20.PrevOutput) {
	commitTxPrevOutputList = make([]*brc20.PrevOutput, 0)
	commitTxPrevOutputList = append(commitTxPrevOutputList, &brc20.PrevOutput{
		TxId:       "fcd1a1c33df653427e20159a799e6c1ba28421fd168fe353a54508c956fb382e",
		VOut:       1,
		Amount:     546,
		Address:    "tb1qtsq9c4fje6qsmheql8gajwtrrdrs38kdzeersc",
		PrivateKey: "cPnvkvUYyHcSSS26iD1dkrJdV7k1RoUqJLhn3CYxpo398PdLVE22",
	})
	commitTxPrevOutputList = append(commitTxPrevOutputList, &brc20.PrevOutput{
		TxId:       "fcd1a1c33df653427e20159a799e6c1ba28421fd168fe353a54508c956fb382e",
		VOut:       0,
		Amount:     252198,
		Address:    "2NF33rckfiQTiE5Guk5ufUdwms8PgmtnEdc",
		PrivateKey: "cPnvkvUYyHcSSS26iD1dkrJdV7k1RoUqJLhn3CYxpo398PdLVE22",
	})
	commitTxPrevOutputList = append(commitTxPrevOutputList, &brc20.PrevOutput{
		TxId:       "fcd1a1c33df653427e20159a799e6c1ba28421fd168fe353a54508c956fb382e",
		VOut:       2,
		Amount:     100000,
		Address:    "mouQtmBWDS7JnT65Grj2tPzdSmGKJgRMhE",
		PrivateKey: "cPnvkvUYyHcSSS26iD1dkrJdV7k1RoUqJLhn3CYxpo398PdLVE22",
	})
	commitTxPrevOutputList = append(commitTxPrevOutputList, &brc20.PrevOutput{
		TxId:       "fcd1a1c33df653427e20159a799e6c1ba28421fd168fe353a54508c956fb382e",
		VOut:       3,
		Amount:     796800,
		Address:    "tb1pklh8lqax5l7m2ycypptv2emc4gata2dy28svnwcp9u32wlkenvsspcvhsr",
		PrivateKey: "cPnvkvUYyHcSSS26iD1dkrJdV7k1RoUqJLhn3CYxpo398PdLVE22",
	})
	return commitTxPrevOutputList
}

func getinscriptionDataList() (inscriptionDataList []brc20.InscriptionData) {
	inscriptionDataList = make([]brc20.InscriptionData, 0)
	inscriptionDataList = append(inscriptionDataList, brc20.InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"xcvb","amt":"1000"}`),
		RevealAddr:  "tb1qtsq9c4fje6qsmheql8gajwtrrdrs38kdzeersc",
	})
	inscriptionDataList = append(inscriptionDataList, brc20.InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"xcvb","amt":"1000"}`),
		RevealAddr:  "2NF33rckfiQTiE5Guk5ufUdwms8PgmtnEdc",
	})
	inscriptionDataList = append(inscriptionDataList, brc20.InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"xcvb","amt":"1000"}`),
		RevealAddr:  "mouQtmBWDS7JnT65Grj2tPzdSmGKJgRMhE",
	})
	inscriptionDataList = append(inscriptionDataList, brc20.InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"xcvb","amt":"1000"}`),
		RevealAddr:  "tb1pklh8lqax5l7m2ycypptv2emc4gata2dy28svnwcp9u32wlkenvsspcvhsr",
	})
	return inscriptionDataList
}
