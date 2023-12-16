package serviceBase

import "github.com/btcsuite/btcd/chaincfg"

func GetNetwork() (network *chaincfg.Params) {
	network = &chaincfg.TestNet3Params
	return network
}
