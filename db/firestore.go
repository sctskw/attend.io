package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreClient interface {
	DatabaseClient
}

type firestoreClient struct {
	client *firestore.Client
}

func NewFirestoreClient() FirestoreClient {
	return &firestoreClient{}
}

func (f *firestoreClient) Connect() {

	d, _ := os.Getwd()
	creds, err := filepath.Abs(filepath.Join(d, "../creds.json"))

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	opts := option.WithCredentialsFile(creds)
	client, err := firestore.NewClient(ctx, "attend-io-294107", opts)
	if err != nil {
		panic(fmt.Sprintf("couldn't connect to firestore db: %s", err.Error()))
	}

	f.client = client
}

func (f *firestoreClient) Ping() {

}

func (f *firestoreClient) FetchAllTalks() (results FetchAllResponse) {
	talks := f.client.Collection("talks")
	docs, err := talks.Documents(context.Background()).GetAll()

	if err != nil {
		panic(err)
	}

	for _, doc := range docs {
		//convert to bytes
		b, _ := json.Marshal(doc.Data())
		results = append(results, b)
	}

	return results
}

func (f *firestoreClient) FetchTalkById(id string) FetchOneResponse {

	doc, err := f.client.Collection("talks").Doc(id).Get(context.Background())

	if err != nil {
		return nil
	}

	//convert to bytes
	b, err := json.Marshal(doc.Data())

	if err != nil {
		panic(err)
	}

	return b

}
