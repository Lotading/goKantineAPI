package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

type Item struct {
	name  string
	price int8
}

func main() {
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
}

func getItem() *Item {

}
