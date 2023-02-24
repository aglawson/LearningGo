package api

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetGasPrice(network string) (*big.Int, error) {
	var client, err = ethclient.Dial(GetRPC(network))
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//fmt.Println("Gas price:", gasPrice)
	return gasPrice, nil
}
