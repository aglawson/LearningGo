package api

import (
	"math/big"
	"web3/contracts"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Metadata struct {
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	TotalSupply *big.Int `json:"totalSupply"`
}

func GetTokenMetadata(contract string, network string) (Metadata, error) {
	metadata := Metadata{}
	var client, err = ethclient.Dial(GetRPC(network))
	if err != nil {
		return metadata, err
	}

	contractAddress := common.HexToAddress(contract)

	token, err := contracts.NewIERC721ACaller(contractAddress, client)
	if err != nil {
		return metadata, err
	}

	supply, err := token.TotalSupply(nil)
	fmt.Println(supply)
	if err != nil {
		return metadata, err
	}

	name, err := token.Name(nil)
	if err != nil {
		return metadata, err
	}

	symbol, err := token.Symbol(nil)
	if err != nil {
		return metadata, err
	}

	metadata = Metadata{
		Name:        name,
		Symbol:      symbol,
		TotalSupply: supply,
	}

	if err != nil {
		return metadata, err
	}

	return metadata, nil
}
