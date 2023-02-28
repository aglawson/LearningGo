package api

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type Price struct {
	From  string      `json:"from"`
	To    string      `json:"to"`
	Value interface{} `json:"value"`
	Time  interface{} `json:"time"`
}

func GetCoinPrice(from string, to string) (Price, error) {
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: FBConf.projectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	query := client.Collection("prices").OrderBy("time", firestore.Desc)

	iter := query.Documents(ctx)

	for {
		doc, err := iter.Next()

		if err != nil {
			return Price{}, err
		}

		if doc.Data()["to"] == to {
			price := Price{
				From:  from,
				To:    to,
				Value: doc.Data()["value"],
				Time:  doc.Data()["time"],
			}

			return price, nil
		}
	}

}
