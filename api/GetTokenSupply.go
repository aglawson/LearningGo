package api

import (
	"math/big"
	"web3/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenSupply(contract string, network string) (big.Int, error) {
	var client, err = ethclient.Dial(GetRPC(network))
	if err != nil {
		return *big.NewInt(-1), err
	}

	contractAddress := common.HexToAddress(contract)

	nft, err := contracts.NewIERC721ACaller(contractAddress, client)

	supply, err := nft.TotalSupply(nil)

	if err != nil {
		return *big.NewInt(-2), err
	}

	return *supply, nil
}
