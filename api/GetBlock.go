package api

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlock(network string) *types.Block {
	client, err := ethclient.Dial(GetRPC(network))
	// Retrieve the latest block number
	blockNumber, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return blockNumber
}
