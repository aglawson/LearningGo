package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Age  int    `json:"age"`
}

type Response struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

var res = Response{Success: true, Data: "wut"}

const KEY = "<YOUR-KEY-HERE>"

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
}

// Connect to an Ethereum node
var client, err = ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/" + KEY)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		queryParams := r.URL.Query()
		name := queryParams.Get("name")

		i := 0
		for i < len(users) {
			if name == users[i].Name {
				res.Success = true
				res.Data = name + " " + strconv.Itoa(users[i].ID)
				json.NewEncoder(w).Encode(res)
				return
			}
			i += 1
		}
		res.Success = false
		res.Data = "Not Found"
		json.NewEncoder(w).Encode(res)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/getBlock", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Retrieve the latest block number
		blockNumber, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		// fmt.Println(blockNumber.NumberU64())
		res.Data = strconv.Itoa(int(blockNumber.NumberU64()))
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/getBalance", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		queryParams := r.URL.Query()
		address := queryParams.Get("wallet")
		blockNumber, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		// Retrieve the latest block number
		balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), blockNumber.Number())
		if err != nil {
			fmt.Println("Error: ", err)
		}
		// fmt.Println(blockNumber.NumberU64())
		//res.Data = strconv.Itoa(balance)
		json.NewEncoder(w).Encode(balance)
	})

	http.HandleFunc("/createWallet", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		privateKeyBytes := crypto.FromECDSA(privateKey)
		//fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		type Result struct {
			PrivateKey string
			Address    string
		}
		result := Result{PrivateKey: hexutil.Encode(privateKeyBytes)[2:], Address: address}

		// hash := sha3.NewKeccak256()
		// hash.Write(publicKeyBytes[1:])
		// fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))

		json.NewEncoder(w).Encode(result)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		res.Success = false
		res.Data = err.Error()

		panic(res)
	}

}
