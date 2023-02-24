package api

import (
	"fmt"
	"math/big"
	"web3/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetOwnedIds(wallet string, contract string, network string) ([]*big.Int, error) {
	var client, err = ethclient.Dial(GetRPC(network))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	walletAddress := common.HexToAddress(wallet)
	contractAddress := common.HexToAddress(contract)

	nft, err := contracts.NewIERC721ACaller(contractAddress, client)

	amount, err := nft.TotalSupply(nil)
	balance := GetTokenBalance(wallet, contract, network)

	if err != nil {
		fmt.Println("error with tx", err)
		return nil, err
	}

	var nftIds []*big.Int
	for i := big.NewInt(0); i.Cmp(amount) < 0; i.Add(i, big.NewInt(1)) {
		owner, err := nft.OwnerOf(nil, i)
		if err != nil {
			//return nil, err
		}
		if walletAddress == owner {
			newValue := big.NewInt(i.Int64())
			fmt.Println(i)
			nftIds = append(nftIds, newValue)
		}
		if balance.Cmp(big.NewInt(int64(len(nftIds)))) == 0 {
			fmt.Println(i.String() + " iterations through loop")
			break
		}
	}

	return nftIds, nil
}
