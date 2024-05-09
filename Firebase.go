package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	credsFile string
	client    *firestore.Client
	ctx       context.Context
}

// Remember to not use the value type here as reciever, use pointer type.
// else, the changes will not persist in your struct
func (fc *FirebaseClient) Client() {
	if fc.client == nil {
		// initialize firebase client as singleton
		opt := option.WithCredentialsFile(fc.credsFile)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		fc.ctx = context.Background()
		if err != nil {
			log.Fatalf("Firebase not connected: %v", err)
		}
		fc.client, err = app.Firestore(fc.ctx)
		if err != nil {
			log.Fatalf("Firestore connection failed: %v", err)
		}
	}
}

func (fc *FirebaseClient) AddToEmail(email string, data map[string]Bookdata) {

	for key, val := range data {
		_, err := fc.client.Collection("users").Doc(email).Collection("books").Doc(key).Set(fc.ctx, val)

		if err != nil {
			log.Fatalf("Failed to add data to firebase: %v", err)
		}
	}
}

func (fc *FirebaseClient) Read(email string) map[string]Bookdata {
	result := make(map[string]Bookdata)
	book_docs := fc.client.Collection("users").Doc(email).Collection("books").Documents(fc.ctx)
	for {
		doc, err := book_docs.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			fmt.Printf("Failed to iterate over firestore docs: %v", err)
		}
		single_book_data := result[doc.Ref.ID]
		doc.DataTo(&single_book_data)
		result[doc.Ref.ID] = single_book_data
	}
	fmt.Printf("RETURNED RESULT: %v\n", result)
	return result
}
