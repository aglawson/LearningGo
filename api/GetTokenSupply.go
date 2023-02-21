package api

import (
	"fmt"
	"math/big"
	"web3/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenSupply(contract string, network string) big.Int {
	var client, err = ethclient.Dial(GetRPC(network))
	if err != nil {
		fmt.Println(err)
		return *big.NewInt(-1)
	}

	contractAddress := common.HexToAddress(contract)

	nft, err := contracts.NewIERC721ACaller(contractAddress, client)

	supply, err := nft.TotalSupply(nil)

	if err != nil {
		fmt.Println("error with tx", err)
		return *big.NewInt(-2)
	}

	return *supply
}
