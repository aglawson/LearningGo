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

var res = Response{Success: true, Data: "Possible Endpoints: get_block, get_balance, create_wallet, get_token_balance, get_gas_price, get_token_supply, get_owned_ids, is_token_holder, get_coin_price"}

func get(param string, r *http.Request) (string, string) {
	query := r.URL.Query()
	value := query.Get(param)

	if value == "" {
		return "", "expected parameter " + param + " was undefined"
	}

	return value, ""
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/get_block", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		network, errString := get("network", r)

		if errString != "" {
			network = "mainnet"
		}

		blockNumber, err := api.GetBlock(network)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		res.Data = strconv.Itoa(int(blockNumber.NumberU64()))
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/get_balance", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		network, errString := get("network", r)

		if errString != "" {
			network = "mainnet"
		}

		address, errString := get("wallet", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		balance, err := api.GetBalance(address, network)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(balance)
	})

	http.HandleFunc("/create_wallet", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result := api.CreateWallet()

		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/get_gas_price", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		network, errString := get("network", r)

		if errString != "" {
			network = "mainnet"
		}

		gasPrice, err := api.GetGasPrice(network)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(gasPrice)
	})

	http.HandleFunc("/get_token_balance", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		network, errString := get("network", r)
		if errString != "" {
			network = "mainnet"
		}

		wallet, errString := get("wallet", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		contract, errString := get("contract", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		result, err := api.GetTokenBalance(wallet, contract, network)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(result.String())
	})

	http.HandleFunc("/get_token_supply", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		network, errString := get("network", r)
		if errString != "" {
			network = "mainnet"
		}

		contract, errString := get("contract", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		result, err := api.GetTokenSupply(contract, network)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(result.String())
	})

	http.HandleFunc("/get_owned_ids", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		wallet, errString := get("wallet", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		contract, errString := get("contract", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		network, errString := get("network", r)
		if errString != "" {
			network = "mainnet"
		}

		result, err := api.GetOwnedIds(wallet, contract, network)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/is_token_holder", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		wallet, errString := get("wallet", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		contract, errString := get("contract", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		network, errString := get("network", r)
		if errString != "" {
			network = "mainnet"
		}

		result, err := api.IsTokenHolder(wallet, contract, network)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/get_coin_price", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		from, errString := get("from", r)
		to, errString := get("to", r)

		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		result, err := api.GetCoinPrice(from, to)

		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/write_coin_price", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		from, errString := get("from", r)
		to, errString := get("to", r)

		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		result, err := api.WriteCoinPrice(from, to)

		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/get_token_metadata", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		contract, errString := get("contract", r)
		if errString != "" {
			json.NewEncoder(w).Encode(errString)
			return
		}

		network, errString := get("network", r)
		if network == "" {
			network = "mainnet"
		}

		result, err := api.GetTokenMetadata(contract, network)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}

		json.NewEncoder(w).Encode(result)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		res.Success = false
		res.Data = err.Error()

		panic(res)
	}

}
