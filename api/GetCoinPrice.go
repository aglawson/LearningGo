package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"context"

	firebase "firebase.google.com/go"
)

func GetCoinPrice(from string, to string) (float64, error) {
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: Conf.FirebaseProjectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	if from == "polygon" || from == "matic" {
		from = "wmatic"
	}

	if from == "eth" {
		from = "ethereum"
	}

	if from == "btc" {
		from = "bitcoin"
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
	_, _, err = client.Collection("prices").Add(ctx, map[string]interface{}{
		"from":  from,
		"to":    to,
		"value": answer,
	})
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}

	return answer, nil
}
