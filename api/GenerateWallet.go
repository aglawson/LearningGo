package api

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Result struct {
	PrivateKey string
	Address    string
}

func CreateWallet() Result {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	result := Result{PrivateKey: hexutil.Encode(privateKeyBytes)[2:], Address: address}
	return result
}
