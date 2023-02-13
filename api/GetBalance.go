package api

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(address string, network string) *big.Int {
	var client, err = ethclient.Dial(GetRPC(network))

	if err != nil {
		fmt.Println("Error: ", err)
	}

	blockNumber, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Retrieve the latest block number
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), blockNumber.Number())
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return balance
}
