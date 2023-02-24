package api

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlock(network string) (*types.Block, error) {
	client, err := ethclient.Dial(GetRPC(network))
	// Retrieve the latest block number
	blockNumber, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return blockNumber, nil
}
