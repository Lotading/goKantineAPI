package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
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

	item := getItem()
	result, err := client.Collection("base").Doc("").Set(context.Background(), item)
	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()

	fetchItem := fetchItems()
	router.HandleFunc("/", fetchItem).Methods("GET")

	defer client.Close()
}

func getItem() *Item {

}

func fetchItems() {

}
