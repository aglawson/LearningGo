package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"web3/api"
)

type Response struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

var res = Response{Success: true, Data: "wut"}

// Connect to an Ethereum node
type Result struct {
	PrivateKey string
	Address    string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/getBlock", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		queryParams := r.URL.Query()

		network := queryParams.Get("network")

		blockNumber := api.GetBlock(network)

		res.Data = strconv.Itoa(int(blockNumber.NumberU64()))
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/getBalance", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		queryParams := r.URL.Query()

		address := queryParams.Get("wallet")
		network := queryParams.Get("network")

		balance := api.GetBalance(address, network)

		json.NewEncoder(w).Encode(balance)
	})

	http.HandleFunc("/create_wallet", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result := api.CreateWallet()

		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/get_gas_price", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		queryParams := r.URL.Query()
		network := queryParams.Get("network")

		gasPrice := api.GetGasPrice(network)

		json.NewEncoder(w).Encode(gasPrice)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		res.Success = false
		res.Data = err.Error()

		panic(res)
	}

}
