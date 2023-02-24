package api

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(address string, network string) (*big.Int, error) {
	var client, err = ethclient.Dial(GetRPC(network))

	if err != nil {
		return nil, err
	}

	blockNumber, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Retrieve the latest block number
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), blockNumber.Number())
	if err != nil {
		return nil, err
	}

	return balance, nil
}
