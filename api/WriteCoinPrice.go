package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"context"
	"time"

	firebase "firebase.google.com/go"
)

func WriteCoinPrice(from string, to string) (float64, error) {
	// Use the application default credentials
	conf := &firebase.Config{ProjectID: FBConf.projectId}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(context.Background())
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

	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + from + "&vs_currencies=" + to

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

	now := time.Now()

	// Get the Unix timestamp in seconds
	timestamp := now.Unix()

	fmt.Println(timestamp)

	fmt.Println(answer)
	_, _, err = client.Collection("prices").Add(context.Background(), map[string]interface{}{
		"from":  from,
		"to":    to,
		"value": answer,
		"time":  timestamp,
	})
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}

	return answer, nil
}
