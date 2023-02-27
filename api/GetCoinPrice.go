package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoinGeckoResponse struct {
	MarketData struct {
		CurrentPrice map[string]float64 `json:"current_price"`
	} `json:"market_data"`
}

func GetCoinPrice(from string, to string) (float64, error) {

	if from == "polygon" {
		from = "wmatic"
	}

	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + from + "&vs_currencies=" + to //"https://api.coingecko.com/api/v3/simple/price?ids=" + name + "&vs_currencies=usd"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return -1, err
	}

	jsonString := string(body)

	data := make(map[string]map[string]float64)

	err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return -1, err
	}

	answer := data[from][to]

	fmt.Println(answer)

	return answer, nil
}
