package api

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetGasPrice(network string) *big.Int {
	var client, err = ethclient.Dial(GetRPC(network))
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return big.NewInt(-1)
	}

	//fmt.Println("Gas price:", gasPrice)
	return gasPrice
}
