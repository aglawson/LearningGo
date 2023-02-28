package api

import (
	"log"

	"context"

	firebase "firebase.google.com/go"
)

type Price struct {
	From  string
	To    string
	Value interface{}
	Time  interface{}
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

	iter := client.Collection("prices").Where("from", "==", from).Documents(ctx)
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
