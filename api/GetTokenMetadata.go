package api

import (
	"math/big"
	"web3/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Metadata struct {
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	TotalSupply *big.Int `json:"totalSupply"`
	Uri         string   `json:"uri"`
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

	uri, err := token.TokenURI(nil, big.NewInt(1))
	if err != nil {
		uri = "no uri link available"
	}

	metadata = Metadata{
		Name:        name,
		Symbol:      symbol,
		TotalSupply: supply,
		Uri:         uri,
	}

	return metadata, nil
}
